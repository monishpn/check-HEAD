package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func cost(s Shape) float64 {
	switch s.(type) {
	case *Rectangle:
		return s.Area() * 2
	case *Circle:
		return s.Area() * 1.5
	case Rectangle:
		return s.Area() * 2
	case Circle:
		return s.Area() * 1.5
	default:
		return 0

	}

}

func main() {
	var s Shape = &Rectangle{width: 10, height: 10}

	fmt.Println(cost(s))

	s = &Circle{radius: 10}
	fmt.Println(cost(s))
}
