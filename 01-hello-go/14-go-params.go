package main

import "fmt"

// 变参函数， 参数数量不固定的 int
func sum(nums ...int) int {
		fmt.Print(nums," ")
		res := 0
		for _, val := range nums{
			res += val
		}
		return res
}
func main() {
	 res1 := sum(1,2,3)
	fmt.Println("res1", res1)

	nums := []int{21,32,43}

	// 如果传slice  需要 nums...
	res2 := sum(nums...)
	fmt.Println("res2", res2)
}