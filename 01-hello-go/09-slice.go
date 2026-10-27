package main

import "fmt"
func main() {
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("apd", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)

	l := s[2:5]
	 fmt.Println("sl1:", l)

	l = s[2:]
	 fmt.Println("sl2:", l)
	l = s[:5]
	fmt.Println("sl3:", l)
	fmt.Println(l)

	f := []string{"g", "h", "i"}
	fmt.Println("dcl:" ,f)

	two := make([][]int, 3)

	for i :=0; i<3; i++ {
		inner := i+1
		two[i] = make([]int, inner)
		for j := 0; j < len(two[i]); j++ {
			two[i][j] = i + j
		}
	}
	fmt.Println("2d: ", two)
}