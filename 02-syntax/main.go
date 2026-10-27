package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// ---------- 1. struct + method：替代 JS 的 class ----------
type User struct {
	ID   int    // 大写 = 导出字段（公开）
	Name string
	age  int // 小写 = 包内私有
}

// 值接收者：拷贝一份（不改原值）
func (u User) Greet() string {
	return "hi, " + u.Name
}

// 指针接收者：可修改原值（JS class 里的 this 行为）
func (u *User) SetAge(a int) {
	u.age = a
}

// ---------- 2. interface：隐式实现，不写 implements ----------
type Repo interface {
	FindByID(id int) (*User, error)
}

type MemRepo struct {
	data map[int]*User
}

func (r MemRepo) FindByID(id int) (*User, error) {
	if u, ok := r.data[id]; ok {
		return u, nil
	}
	return nil, ErrNotFound // 哨兵 error，可用 errors.Is 判断
}

// ---------- 3. error：显式返回，没有 try/catch ----------
var ErrNotFound = errors.New("user not found")

func LoadUser(r Repo, id int) (*User, error) {
	u, err := r.FindByID(id)
	if err != nil {
		// %w 包装：保留原始 error 链（类似 JS 的 error.cause）
		return nil, fmt.Errorf("load user %d: %w", id, err)
	}
	return u, nil
}

// ---------- 4. goroutine + channel：并发模型 ----------
// 对照 Node：
//
//	Go:   go f() 启动协程，chan 通信，CSP 模型
//	Node: Promise + await / worker_threads
func fetchAll(ids []int) []string {
	var wg sync.WaitGroup
	ch := make(chan string, len(ids)) // 带缓冲 channel，避免 goroutine 泄漏

	for _, id := range ids {
		wg.Add(1)
		go func(id int) { // 注意：参数传进来，Go 1.22+ 循环变量才是每轮独立
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			ch <- fmt.Sprintf("user-%d", id)
		}(id)
	}

	// 必须等所有 goroutine 结束后再关 channel，否则 panic
	go func() {
		wg.Wait()
		close(ch)
	}()

	var out []string
	for s := range ch { // channel 关闭后自动退出
		out = append(out, s)
	}
	return out
}

// ---------- 5. context：控制 goroutine 生命周期，防泄漏 ----------
func slowJob(ctx context.Context) (string, error) {
	select {
	case <-time.After(2 * time.Second):
		return "done", nil
	case <-ctx.Done(): // 上游取消或超时
		return "", ctx.Err()
	}
}

func main() {
	// 1-2. struct / interface
	u := &User{ID: 1, Name: "bitjian"}
	u.SetAge(18)
	fmt.Println(u.Greet(), u.age)

	repo := MemRepo{data: map[int]*User{2: {ID: 2, Name: "go"}}}
	got, err := LoadUser(repo, 2)
	fmt.Println(got, err)

	// 3. error 链判断
	_, err = LoadUser(repo, 999)
	fmt.Println("Is ErrNotFound:", errors.Is(err, ErrNotFound), "|", err)

	// 4. 并发
	start := time.Now()
	fmt.Println(fetchAll([]int{1, 2, 3, 4, 5}), "耗时", time.Since(start))

	// 5. context 超时
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel() // 关键：不 cancel 会泄漏
	res, err := slowJob(ctx)
	fmt.Println(res, err)
}
