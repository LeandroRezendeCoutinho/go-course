package functions

func Call_anonimous() {
	// anonimous function
	sum := func(a int, b int) int {
		return a + b
	}(1, 2)

	println(sum)
}
