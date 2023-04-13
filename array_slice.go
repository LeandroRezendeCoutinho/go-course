package main

import (
	"fmt"
	"reflect"
)

func call_array() {
	// fixed array
	var array1 [5]int
	array1[0] = 10
	array1[1] = 20
	fmt.Println("array1")
	fmt.Println(array1)

	array2 := [2]string{"value1", "value2"}
	fmt.Println("array2")
	fmt.Println(array2)

	// inferred fixed array
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("array3")
	fmt.Println(array3)
	fmt.Println(reflect.TypeOf(array3))
}

func call_slice() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice")
	fmt.Println(slice)

	slice = append(slice, 7)
	fmt.Println("slice")
	fmt.Println(slice)

	array := [...]int{1, 2, 3, 4, 5}
	slice2 := array[2:4]
	fmt.Println("slice2")
	fmt.Println(slice2)

	array[2] = 10
	fmt.Println("slice2 array altered")
	fmt.Println(slice2)

	fmt.Println("reflect.TypeOf(slice)")
	fmt.Println(reflect.TypeOf(slice))

	// internal array slice
	// go creates a internal array to initialize the slice
	slice3 := make([]float32, 10, 12)
	fmt.Println("slice3")
	fmt.Println(slice3)
	fmt.Println("len(slice3)")
	fmt.Println(len(slice3))
	fmt.Println("cap(slice3)")
	fmt.Println(cap(slice3))

	// when the value overflows, it creates internally a new slice with double capacity
	slice3 = append(slice3, 1)
	slice3 = append(slice3, 2)
	slice3 = append(slice3, 3)
	fmt.Println("slice3")
	fmt.Println(slice3)
	fmt.Println("len(slice3)")
	fmt.Println(len(slice3))
	fmt.Println("cap(slice3)")
	fmt.Println(cap(slice3))

	slice4 := make([]int, 5)
	fmt.Println("slice4")
	fmt.Println(slice4)
	fmt.Println("len(slice4)")
	fmt.Println(len(slice4))
	fmt.Println("cap(slice4)")
	fmt.Println(cap(slice4))
}
