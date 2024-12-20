package main

import "fmt"

func main() {
	name := "Moaz"
	fmt.Println(&name)

	var x int = 2
	ptr := &x
	fmt.Printf("pts is of type %T with a value of %v and address of %p\n", ptr, ptr, &ptr)
	fmt.Printf("address of x is %p\n", &x)

	var ptr1 *float64
	_ = ptr1

	p := new(int)

	x = 100
	p = &x

	fmt.Printf("ps is of type %T with a value of %v and address of %p\n", p, p, &p)

	*p = 90
	fmt.Println(x, *p)

	fmt.Println("*p -- x :", *p == x)

	*p = 10
	*p = *p / 2
	fmt.Println(x)

	// &value => pointer
	// *pointer => value
}
