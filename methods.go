package main

import "fmt"

func Call_methods() {
	var user1 user
	user1.name = "UserName 1"
	user1.age = 10
	user1.address = address{"street", 99}

	user1.show() //{UserName 1 10 {street 99}}

	user2 := user{"UserName 2", 10, address{"street", 99}}
	user2.makeBirhtday()
	user2.show() //{UserName 2 11 {street 99}}
}

func (u user) show() {
	fmt.Println(u)
}

func (u *user) makeBirhtday() {
	u.age++
}
