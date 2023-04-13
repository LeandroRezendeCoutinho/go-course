package main

import "fmt"

func call_maps() {
	fmt.Println("Maps")

	user := map[string]string{
		"user":     "user",
		"password": "password",
	}
	fmt.Println(user)

	table := map[int]string{
		1: "user1",
		2: "user2",
	}
	fmt.Println(table)

	fmt.Println("user_info")
	user_info := map[string]map[string]string{
		"user": {
			"first_name":  "John",
			"second_name": "Doe",
		},
		"course": {
			"name": "engineering",
		},
	}
	fmt.Println(user_info)

	fmt.Println("delete course")
	delete(user_info, "course")
	fmt.Println(user_info)

	//assing new value
	fmt.Println("user_info period")
	user_info["period"] = map[string]string{
		"name": "first",
	}
	fmt.Println(user_info)

}
