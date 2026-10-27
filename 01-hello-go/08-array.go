package main
import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp", a)

	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("get", a[4])

	b :=[5]int{1,2,3,4,5}
	fmt.Println("dcl", b)

	var two [2][3]int
	for i := 0; i < len(two); i++ {
		fmt.Println("two", two, len(two))
		for j :=0; j<len(two[0]); j++ {
			fmt.Println("two[0][0]", len(two[0]))
			two[i][j] = i+j

		}
			fmt.Println("two[i][j]", two)
	}
}
