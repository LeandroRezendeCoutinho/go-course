package main

import "fmt"

func main() {
	fmt.Print("Course")

	// function
	sum := sum(1, 2)
	fmt.Println(sum)

	result := f("Texto da função")
	fmt.Println(result)

	sumResult, subResult := calculations(10, 20)
	fmt.Println(sumResult, subResult)

	// Ignore second result
	sumResult, _ = calculations(10, 20)
	fmt.Println(sumResult)

	_, subResult = calculations(10, 20)
	fmt.Println(subResult)

	//struct composition
	var u user
	u.name = "UserName 1"
	u.age = 10
	fmt.Println(u)

	address := address{"street", 99}
	usuario2 := user{"UserName 2", 10, address}
	fmt.Println(usuario2)

	usuario2 = user{name: "User 2"}
	fmt.Println(usuario2)

	//inheritance
	student1 := student{usuario2, "Course", "College"}
	fmt.Println((student1))
}
