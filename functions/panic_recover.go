package functions

func Call_panic_recover() {
	should_not_be_zero(1, 0)
}

func should_not_be_zero(n1, n2 int) {
	defer func() {
		if err := recover(); err != nil {
			println("recover:", err)
		}
	}()

	if n2 == 0 {
		panic("n2 should not be zero")
	}
	if n1/n2 == 0 {
		panic("n1/n2 should not be zero")
	}
}
