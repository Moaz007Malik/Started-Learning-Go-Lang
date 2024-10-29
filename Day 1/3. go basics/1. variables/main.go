package main

import "fmt"

func main(){
	var w = 2
	var x int = 4
	var y int
	y = 6

	//no need to write var with it
	z := 9
	// fmt.Println(w, x, y, z)
	// if i don't want to print z, i will get error, to remove that error i will use _ instead.
	//just like this
	_ = z
	fmt.Println(w, x, y)

	//Multiple Declaration

	car,cost := "Audi", 50000
	fmt.Println(car, cost)

	// Error on next line
	// car, cost := "BMW", 60000
	
	// No error on next line because now its a new one. It's not same as the previous one.
	car, year := "BMW", 2024
	// but we have not printed year so it will give error so to fix this we will use _ instead.
	_ = year

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary float64
		firstName string
		gender bool
	)

	fmt.Println(salary, firstName, gender)


	var a, b, c int
	fmt.Println (a, b, c)

	var d, e, f int = 10, 20, 30
	fmt.Println(d, e, f)

	var i, j int
	i, j = 5, 8

	j, i = i, j //Swapping Variables

	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)
}




// package main

// import "fmt"

// func main() {

//     //** DECLARING VARIABLES **///

//     // Syntax: var var_name type
//     var s1 string
//     s1 = "Learning Go!"
//     fmt.Println(s1) // printing string s1

//     //** TYPE INFERENCE **//

//     // Go deduces automatically the type of the variable by looking at the initial value (bool, int, string etc)

//     var k int = 6 // not necessary to say the type (int). It is inferred from the literal on the right side of =
//     var i = 5     // type int
//     var j = 5.6   // type float64

//     // printing i, j and k
//     fmt.Println("i:", i, "j:", j, "k:", k)

//     // ii == jj  // -> error: cannot assign float to int (Go is a strong typed language)

//     // declaring and initializing a new variable of type string (type inference)
//     var s2 = "Go!"
//     _ = s2 //in Go each variable must be used or there is a compile-time error
//     // _ is the Blank Identifier and mutes the error of unused variables
//     // _ can be only on the left hand side of the = operator

//     // multiple assignments
//     var ii, jj int
//     ii, jj = 5, 8 // -> tuple assignment. It allows several variables to be assigned at once

//     // swapping two variables using multiple assignments
//     ii, jj = jj, ii

//     fmt.Println(ii, jj)

//     //** Short Declaration (works only in Block Scope!) **//

//     // := (colon equals syntax) used only when declaring a new variable (or at least a new variable)
//     // := tells go we are going to create a new variable and go figures out what type it will be
//     s3 := "Learning golang!"
//     _ = s3

//     // can't use short declaration at Package Scope (outside main() or other function)
//     // all statements at package scope must start with a Go keyword (package, var, import, func etc)

//     // multiple short declaration
//     car, cost := "Audi", 50000
//     fmt.Println(car, cost)

//     // redeclaration with short declaration syntax
//     // at least one variable must be NEW on the left side of :=
//     var deleted = false
//     deleted, file := true, "a.txt"
//     _, _ = deleted, file

//     // expressions in short declarations are allowed
//     sum := 5 + 2.5
//     fmt.Println(sum)

//     // multiple declaration is good for readability
//     var (
//         age       float64
//         firstName string
//         gender    bool
//     )
//     _, _, _ = age, firstName, gender

//     // a concise way to declare multiple variables that have the same type
//     var a, b int
//     _, _ = a, b

// }