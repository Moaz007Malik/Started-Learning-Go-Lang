package main

import "fmt"

func main() {
	var g = 4
	var t = 5.2

	// Error: cannot use b (variable of type float64) as int value in assignment
	// a = b

	// Now, no error
	g = int(t)

	fmt.Println(g, t)

	var x int
	// Error: cannot use "5" (untyped string constant) as int value in assignment
	// x = "5"

	// Now, no error
	x = 5

	fmt.Println(x)

	// OR

	// you must provide a type for each variable you declare or Go should be able to infer it
    var a int = 10
    var b = 15.5      // type inference (deduction)
    c := "Gopher"     // short declaration, type inference
    _, _, _ = a, b, c // Blank Identifier (_) to get rid of unused variable error

    // Go is a Statically and Strong Typed Programming Language
    // a = 3.14 -> error. A variable cannot change it's type
    // a = b    -> error. It's not allowed to assign a type to another type


	
	// Zero Values

	// - Numerice Types: 0
	// - Bool Type: false
	// - string type: "" (empty string)
	// - pointer type: nil

	var value int
	var price float64
	var name string
	var done bool

	// Only Declared and Initialized. But not assigned. So, by default, these are on zero Values

	fmt.Println(value, price, name, done)

	// Output: 0 0  false
	// Means number has 0, string has empty string, and boolean has false. Pointer is not declared yet otherwise it would be nil
}