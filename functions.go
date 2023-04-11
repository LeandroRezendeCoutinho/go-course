package main

import "fmt"

var f = func(txt string) string {
	fmt.Print(txt)
	return txt
}

func sum(n1 int, n2 int) int8 {
	return int8(n1) + int8(n2)
}

func calculations(n1, n2 int) (int8, int8) {
	sum := int8(n1) + int8(n2)
	sub := int8(n1) - int8(n2)
	return sum, sub
}
