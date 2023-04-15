package functions

import "fmt"

func Call_variatic() {
	// variadic
	total := sumNumbers(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
}

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
