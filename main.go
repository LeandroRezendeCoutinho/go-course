package main

import "fmt"

func main() {
	fmt.Print("Course")

	// function
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

	//structs
	var u user
	u.name = "Usuario"
	u.age = 10
	fmt.Print(u)

	address := address{"street", 99}
	usuario2 := user{"Name", 10, address}
	fmt.Print(usuario2)

	usuario2 = user{name: "New Name"}
	fmt.Print(usuario2)
}
