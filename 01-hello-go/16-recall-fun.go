package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))
		// 不能写成一行 fib := func(n int) int { ... return fib(n-1) ... }。
		// 因为 Go 规定：:= 声明的变量，作用域从这条语句结束后才开始
		// ，所以右侧匿名函数里的 fib 还没"出生"，
		// 编译器报 undefined: fib。
		// 先用 var 把 fib 声明出来（此时它的作用域覆盖后面所有语句），
		// 再赋值，内部的递归才能引用到它。
    var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n-1) + fib(n-2)

    }

    fmt.Println(fib(7))
}