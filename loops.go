package main

import "fmt"

func call_loops() {
	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// goto
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("goto label")
		}
		fmt.Println(i)
	}

	//range
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Println(index, value)
	}

	// range with string
	for index, value := range "Hello World" {
		fmt.Println(index, string(value))
	}

	// range without index
	for _, value := range slice {
		fmt.Println(value)
	}

	// range with map
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range m {
		fmt.Println(key, value)
	}

	//range with struct does not work
}
