package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	s := square{
		sideLength: 10,
	}
	t := triangle{
		height: 10,
		base:   10,
	}

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println("Area: ", s.getArea())
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2.0
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
