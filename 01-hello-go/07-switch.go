package main
import (
    "fmt"
    "time"
)

func main() {
	i := 2
	fmt.Println("write", i, "as") 
	
	switch i {
	case 1: 
		fmt.Println("one")
	case 2:
		fmt.Println("TWO")	
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the Weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI :=func(i interface {}) {
		switch t := i.(type){
		case bool:
			fmt.Println("i am bool")
		case int:
			fmt.Println("i am int")
		default:
			// fmt.Println("don't know type %T\n", t)
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(1)
	whatAmI(true)
	whatAmI("hey")
}
