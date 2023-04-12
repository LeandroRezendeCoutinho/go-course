package main

import "fmt"

func call_variables() {
	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println(variable1, variable2)

	variable2++
	fmt.Println(variable1, variable2)

	// using reference
	var variable3 int
	var pointer1 *int

	variable3 = 100
	pointer1 = &variable3             // user memory reference, not value
	fmt.Println(variable3, pointer1)  // get pointer memory address
	fmt.Println(variable3, *pointer1) //dereferencing(get pointer value)

	variable3 = 150
	fmt.Println(variable3, pointer1)
	fmt.Println(variable3, *pointer1)

	*pointer1 = 200
	fmt.Println(variable3, pointer1)
	fmt.Println(variable3, *pointer1)
}
