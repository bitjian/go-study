package main

import "fmt"

func fun1() (int, int) {
	return 11, 2
}

func main() {
	v1, v2 := fun1()
	fmt.Println("v1", v1, "v2", v2)

	_, c := fun1()
	fmt.Println("c",  c)
}