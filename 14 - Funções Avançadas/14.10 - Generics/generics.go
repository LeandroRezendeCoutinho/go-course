package main

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func genericSum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	println(genericSum(m))

	m2 := map[string]float32{
		"a": 1.1,
		"b": 2.2,
		"c": 3.3,
	}
	println(genericSum(m2))
}
