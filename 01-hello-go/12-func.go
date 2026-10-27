package main

import "fmt"

// go 需要明确返回， 函数参数 参数 类型
func plus(a int , b int ) int {
	return a+b
}

//  参数如果都是相同的类型，可以在最后声明一个类型
func plusPlus(a, b, c int) int {
	return a+b+c
}

func main() {
	res := plus(1,2)
	fmt.Println("plus res value", res)

	res = plusPlus(1,2,3)
	fmt.Println("plusPlus res", res)
}