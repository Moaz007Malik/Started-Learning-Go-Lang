package main

import "fmt"

// package scoped variables and constants
var x int = 100

const y = 0

// a function declaration. main is a special function name
// it is the entry point for the executable program
// Go uses brace brackets to delimit code blocks

func main(){
	// Local Scoped Variables and Constants Declarations, calling functions etc
	var a int = 7
	var b float64 = 3.5
	var c string = "Hello, World!"

	// Println() function prints out a line to stdout
    // It belongs to package fmt
	fmt.Println("Hello This is Go Lang")
	fmt.Println("Values of a,b,c= \n", a, b, c)

	//printing package scope variables
	fmt.Println("Value of x", x)
	fmt.Println("Value of y", y)
}