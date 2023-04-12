package main

type user struct {
	name    string
	age     int8
	address address
}

type address struct {
	street string
	number int16
}
