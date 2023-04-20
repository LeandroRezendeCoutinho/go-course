package main

import "fmt"

func Call_generics() {
	genericWithInterface(10)
	genericWithInterface("string")
	genericWithInterface(true)
	genericWithInterface(10.5)
	genericWithInterface([]string{"a", "b", "c"})
	genericWithInterface(map[string]string{"a": "b", "c": "d"})
}

// when interface is empty anyone implements it
func genericWithInterface(i interface{}) {
	fmt.Println(i)
}
