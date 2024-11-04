package main

import "fmt"

func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	add := func(x, y int) int { return x + y }
	result := applyOperation(5, 3, add)
	fmt.Println(result)
}
