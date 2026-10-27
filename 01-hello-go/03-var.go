package main

import "fmt"
var test string = "test"

func main() {
   fmt.Println(test)
		// var是声明变量用的 声明一个string变量
    var a = "initial"
    fmt.Println(a)
		//  var 可以生成多个变量
    var b, c int = 1, 2
    fmt.Println(b, c)

		// var 可以自动推到变量类型
    var d = true
    fmt.Println(d)
		//  如果没有给默认值，会自动复制一个默认值0
    var e int
    fmt.Println(e)
		//  := 是一个语法糖 声明变量并赋值  var f string = "short"
    f := "short"
    fmt.Println(f)
}