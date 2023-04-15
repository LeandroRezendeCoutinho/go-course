package functions

func Call_recursive() {
	// recursive
	result := factorial(5)
	println("Factorial of 5")
	println(result)
	result = fibonacci(5)
	println("Fibonacci of 5")
	println(result)
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func fibonacci(x int) int {
	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	return fibonacci(x-1) + fibonacci(x-2)
}
