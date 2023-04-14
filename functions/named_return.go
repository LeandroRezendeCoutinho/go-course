package functions

import "fmt"

func Call_named_return() {
	// named return
	fmt.Println(sum(1, 2))
}

func sum(i1, i2 int) (sum int, sub int) {
	sum = i1 + i2
	sub = i1 - i2
	return
}
