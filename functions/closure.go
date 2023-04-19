package functions

import "fmt"

func Call_closures() {
	number := closure()
	x := 1
	fmt.Println(x)
	fmt.Println(number())
}

func closure() func() int {
	x := 10
	var function = func() int {
		x++
		return x
	}
	return function
}
