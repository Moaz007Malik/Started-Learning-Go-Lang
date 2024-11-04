package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width - r.height)
}

// func printCircle(c circle) {
// 	fmt.Println("Shape: ", c)
// 	fmt.Println("Area: ", c.area())
// 	fmt.Println("Perimeter: ", c.perimeter())
// }

// func printRectangle(r rectangle) {
// 	fmt.Println("Shape: ", r)
// 	fmt.Println("Area: ", r.area())
// 	fmt.Println("Perimeter: ", r.perimeter())
// }

func print(s shape) {
	fmt.Printf("shape: %#v\n", s)
	fmt.Printf("area: %v\n", s.area())
	fmt.Printf("perimeter: %v\n", s.perimeter())
}

func main() {
	var s shape
	fmt.Printf("%T\n", s)

	ball := circle{radius: 5}
	s = ball
	s = rectangle{width: 3, height: 2.1}

	print(s)
	fmt.Printf("Type of s: %T\n", s)

	room := rectangle{width: 3, height: 2.1}
	s = room

	print(s)
	fmt.Printf("Type of s: %T\n", s)
}
