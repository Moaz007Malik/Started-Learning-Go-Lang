package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{10, 5}
	fmt.Println("Area:", rect.Area()) // Output: Area: 50
}
