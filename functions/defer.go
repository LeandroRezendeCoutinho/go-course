package functions

func Call_defer() {
	// defer
	defer fun1()
	fun2()
	makecalculation(100, 60)
}

func fun1() {
	println("function 1 defer")
}

func fun2() {
	println("function 2 called after function 1")
}

func makecalculation(num1, num2 float32) bool {
	defer println("Median calculated") //defer is executed before the function return
	println("Making calcs")
	median := (num1 + num2) / 2
	println("Median: ", median)
	return median > 6
}

// Think defer like a finally in Javascript
