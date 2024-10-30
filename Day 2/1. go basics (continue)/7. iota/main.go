package main

import "fmt"

func main(){
	// iota is a constant that is incremented by 1 each time it is used
	// It is useful for initializing variables with consecutive values
	
	const (
	c1 = iota
	c2 = iota
	c3 = iota
	)

	fmt.Println(c1, c2, c3)

	const (
		zero = iota
		one
		two
		three
	)
	
	fmt.Println(zero, one, two, three)

	//const (
	//	zero = iota
	//	one
	//	two
	//	three
	//)

	//Error: iota is not a variable


	const (
		east = iota * 2
		west
		north
		south
	)

	fmt.Println(east, west, north, south)
	
	const (
		a = (iota * 2)
		b
		c
		d
	)
	
    fmt.Println(a, b, c, d)

	// x = -2, y =-4 and z = -5

	const (
		x = -(iota + 2) // 2
		_ // 3
		y // 4
		z // 5
	)
	fmt.Println(x, y, z)

	
}