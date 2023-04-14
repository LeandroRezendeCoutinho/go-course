package main

import "fmt"

func call_control_structures() {

	// if else
	number := 10
	if number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Lesser than 15")

	}

	// number2 only works in this scope
	if number2 := 10; number2 > 0 {
		fmt.Println("Greater than 10")
	} else {
		fmt.Println("Lesser than 10")
	}

	// switch
	fmt.Println(weekday(1))
	fmt.Println(weekday(3))
	fmt.Println(weekday(7))
	fmt.Println(yearMonth(1))
	fmt.Println(yearMonth(3))
	fmt.Println(yearMonth(7))
	fmt.Println(yearMonth(12))
}

func weekday(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid weekday number"
	}
}

func yearMonth(number int) string {
	switch {
	case number == 1:
		return "January"
	case number == 2:
		return "February"
	case number == 3:
		return "March"
	case number == 4:
		return "April"
	case number == 5:
		return "May"
	case number == 6:
		return "June"
	case number == 7:
		return "July"
	case number == 8:
		return "August"
	case number == 9:
		return "September"
	case number == 10:
		return "October"
	case number == 11:
		return "November"
	case number == 12:
		return "December"
	default:
		return "Invalid month number"
	}
}
