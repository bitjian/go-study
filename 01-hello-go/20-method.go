package main

import "fmt"

// ============================================================
// 方法（Method）与接收者（Receiver）
// ------------------------------------------------------------
// func (r rect) perim() int 里的 (r rect) 叫「接收者」：
//   - 作用：把函数绑定到 rect 类型上，变成方法，可以用 r.perim() 调用
//   - r 只是接收者的变量名（相当于 JS class 里的 this），可以任意取名
//   - rect 是接收者类型；写成 *rect 就是指针接收者
//
// 没有接收者的写法是普通函数：func perim(r rect) int，调用时写 perim(r)
// ============================================================

type rect struct {
	width, height int
}

// 指针接收者 (r *rect)：拿到的是原对象的指针
//   - 方法内修改 r.width 会改到原对象
//   - 只传一个指针（8 字节），结构体大时避免拷贝
func (r *rect) area() int {
	return r.width * r.height
}

// 值接收者 (r rect)：拿到的是 rect 的一份副本
//   - 方法内修改 r.width 改不到原对象（改的是副本）
//   - 结构体越大，拷贝开销越大
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// 值接收者 vs 指针接收者，最能体现差异的场景是「修改字段」：
//   func (r rect)  setW1(w int) { r.width = w }  // 无效：改的是副本
//   func (r *rect) setW2(w int) { r.width = w }  // 有效：改到原对象
//
//   r := rect{10, 5}
//   r.setW1(99)   // r.width 仍是 10
//   r.setW2(99)   // r.width 变成 99
//
// 对比 JS：JS 里 this 永远是对象的引用，方法内赋值必然改到原对象；
// Go 需要你通过接收者写法显式选择是「拷贝」还是「引用」。

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go 会自动做地址 / 解引用转换（语法糖）：
	//   r.area()    → (&r).area()     值变量自动取地址
	//   rp.perim()  → (*rp).perim()   指针自动解引用
	// 所以「值」和「指针」都能调用两种接收者的方法。
	//
	// 但这个糖只对「可寻址」的值生效，以下写法编译不过：
	//   rect{10, 5}.area()   // cannot call pointer method on rect
	//   m["a"].area()        // map 元素不可寻址
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

// ============================================================
// 接收者选型准则
//   1. 方法要修改字段                     → 必须用指针
//   2. 结构体较大                         → 用指针，避免拷贝
//   3. 含 sync.Mutex 等禁止拷贝的字段      → 必须用指针
//   4. 小结构、只读、值语义（time.Time / 坐标 / 金额）→ 用值
//   5. 同一个类型的方法尽量统一用同一种接收者，混用容易踩坑
//
// 接口实现的影响（面试常考点）
//   值接收者方法   → T 和 *T 都算实现了接口
//   指针接收者方法 → 只有 *T 算实现，T 不算
//
//     type Shape interface{ area() int }
//     var s Shape = r    // ❌ area 是指针接收者时，rect 不实现 Shape
//     var s Shape = &r   // ✅
//
// 其他注意点
//   - 接收者名习惯取类型首字母，且同类型所有方法统一用一个名字
//     （Go 官方建议：不要用 this / self）
//   - 指针接收者可能是 nil，方法内使用前要判空，否则 panic
//     （类似 JS 里 this 可能是 undefined）
//
// 回到本文件：area / perim 都只是读字段不修改，用哪种接收者都能跑；
// area 写成指针接收者只是为了演示这种写法。
// ============================================================
