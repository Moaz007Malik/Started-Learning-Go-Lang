package main

import (
	"fmt"
	"strings"
)

func main() {
	// Append
	fmt.Println(strings.Repeat("#", 30))
	fmt.Println("Append")
	fmt.Println(strings.Repeat("#", 30))

	numbers := []int{2, 3}
	fmt.Printf("numbers: %v\n", numbers)

	numbers = append(numbers, 10)
	fmt.Printf("numbers: %v\n", numbers)

	numbers = append(numbers, 20, 30, 40)
	fmt.Printf("numbers: %v\n", numbers)

	n := []int{100, 200}
	numbers = append(numbers, n...)
	fmt.Printf("numbers: %v\n", numbers)

	src := []int{10, 20, 30}
	// dst := make([]int, len(src)) // Coping whole length
	// dst := make([]int, 2) // Coping only 2 elements
	dst := make([]int, 0) // if I replace it with 0, then nothing will be copied
	nn := copy(dst, src)
	fmt.Println(src, dst, nn)

	// Expressions
	fmt.Println(strings.Repeat("#", 30))
	fmt.Println("Expressions")
	fmt.Println(strings.Repeat("#", 30))
	// First lets see with Array
	a := [5]int{1, 2, 3, 4, 5}
	// a[start:stop]
	b := a[1:3]
	fmt.Printf("%v, %T\n", b, b)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Printf("%v, %T\n", s2, s2)
	// It will give all items from index 2 to the end
	fmt.Println(s1[2:]) // same as s1[2:len(s1)]
	// It will give all items from index 3 to the start
	fmt.Println(s1[:3]) // same as s1[0:3]
	// It will give all items from start to end
	fmt.Println(s1[:]) // same as s1[0:len(s1)]
	// This will give runtime error: slicee bounds out of range
	// fmt.Println(s1[1:100])

	s1 = append(s1[:4], 100) // It iwll add 100 after index 3
	fmt.Println(s1)

	s1 = append(s1[:4], 200) // It iwll add 200 after index 3 and remove 100. Means it will overwrite the last element
	fmt.Println(s1)
}
