package functions

import "fmt"

func Call_pointer() {
	number := 10

	invertedNumber := invertSignal(number)
	fmt.Println(number)
	fmt.Println(invertedNumber)

	invertSignalPointer(&number)
	fmt.Println(number)
}

func invertSignalPointer(number *int) {
	*number = *number * -1
}

func invertSignal(number int) int {
	return number * -1
}
