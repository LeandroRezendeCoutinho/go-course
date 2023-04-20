package main

import "math"

type rectangle struct {
	width  int
	height int
}

type circle struct {
	radius int
}

type shape interface {
	area() float64
}

func writeArea(s shape) {
	println(s.area())
}

func (r rectangle) area() float64 {
	return float64(r.width * r.height)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}

func Call_interfaces() {
	var r rectangle
	r.width = 10
	r.height = 5

	writeArea(r)

	var c circle
	c.radius = 10

	writeArea(c)
}
