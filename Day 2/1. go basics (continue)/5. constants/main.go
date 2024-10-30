package main

import "fmt"

func main(){
	const days int = 7
	var i int
	fmt.Println(i)

	// It is mandatory to assign value white declaration
	// const pi float64
	const pi = 3.14159
	const secondsInHour = 3600

	duration := 234
	fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

	x, y := 5, 0

	// error while exeution
	// fmt.Println(x / y)

	_, _ = x, y

	const a = 5
	const b = 0

	// runtime Error 
	// fmt.Println(a / b)

	const n, m int = 4, 5
	const n1, m2 = 6, 7

	const (
		min1 = 12
		min2 = 13
		min3 = 14
	)
	fmt.Println(min1, min2, min3)

	const (
		max1 = 12
		max2
		max3
	)
	fmt.Println(max1, max2, max3)



	// Contant Rules
	// 1. You cannot change a constant

	const temp int = 100
	// Error
	// temp = 50

	// 2. You cannot initiate a constant at runtime

	// const power math.Pow(2, 3)

	// 3. You cannot use a variable to initialize a constant
	t := 5
	//Error
	// const tc = t
	_ = t

	// 4. You can use a function like len. len returns the number of elements in a string or in another sequence
	// You can use this function to initialize a constant if it has as argument a constant string literal, 
	// not a variable.
	// This is because a string literal is in fact an unnamed constant

	const l1 = len("Hello Go")
	fmt.Println(l1)

	//len is built-in but math.Pow is not

	
}