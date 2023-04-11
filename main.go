package main

import "fmt"

func main() {
	fmt.Print("Course")

	sum := sum(1, 2)
	fmt.Print(sum)

	result := f("Texto da função")
	fmt.Print(result)

	sumResult, subResult := calculations(10, 20)
	fmt.Print(sumResult, subResult)

	// Ignore second result
	sumResult, _ = calculations(10, 20)
	fmt.Print(sumResult)

	_, subResult = calculations(10, 20)
	fmt.Print(subResult)
}
