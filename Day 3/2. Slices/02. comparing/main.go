package main

import "fmt"

func main() {
	var n []int
	fmt.Println(n == nil)
	// Output: true; Because n is not initialized

	m := []int{}
	fmt.Println(m == nil)
	// Output: false; Because m is initialized

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	// fmt.Println(a == b)
	// Error: Slice can only be compared to nil

	var eq bool = true

	a = nil
	for i, valueA := range a{
		if valueA != b[i]{
			eq = false
			break
		}
	}
	if len(a) != len(b){
		eq = false
	}
	if eq{
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}
}
