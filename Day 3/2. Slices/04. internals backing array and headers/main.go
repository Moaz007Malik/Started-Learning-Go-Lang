package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3]
	fmt.Println(s1, s3, s4)
	s3[1] = 600
	fmt.Println(s1, s3, s4)

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	fmt.Println(arr1, slice1, slice2)
	arr1[1] = 2
	fmt.Println(arr1, slice1, slice2)

	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)
	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	ss1 := []int{10, 20, 30, 40, 50}
	newSlice := ss1[0:3]
	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = ss1[2:5]
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Printf("%p\n", &ss1)
	fmt.Printf("%p %p\n", &ss1, &newSlice)

	newSlice[0] = 1000
	fmt.Println("ss1: ", ss1)

	a := [5]int{1, 2, 3, 4, 5} // a is an array
	s := []int{1, 2, 3, 4, 5}  // s is a slice
	fmt.Printf("array's size in bytes: %d\n", unsafe.Sizeof(a))
	fmt.Printf("slice's size in bytes: %d\n", unsafe.Sizeof(s))
}
