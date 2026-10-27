package main

import "fmt"
func main() {
	// 初始化 slice
	nums := []int{2,3,4}
	sum := 0
	// 遍历 slice  index, num
	for _, num := range nums {
		sum += num
	}

	fmt.Println("range sum", sum)

	for index, num := range nums {
			fmt.Printf("value %v  index %v\n", num, index)
	}

	// 初始化 map
	kvs := map[string]int{"foo":2, "boo": 3}
	// 遍历map
	for k, v := range kvs {
		fmt.Printf("key %v, value %v \n", k, v)
	}
	// 遍历字符串  index , unicode 码点
	for index, chat := range "hello" {
		fmt.Printf("value %v  index %v\n", chat, index)
	}
}