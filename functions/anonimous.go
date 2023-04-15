package functions

func Call_anonimous() {
	// anonimous function
	sum := func(a int, b int) int {
		return a + b
	}(1, 2)

	println("Anonimous function of sum 1 + 2")
	println(sum)
}
