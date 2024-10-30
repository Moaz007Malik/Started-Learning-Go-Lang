package main

import (
	"fmt"
	"time"
)

func main() {
	language := "Go"

	switch language {
	case "Python":
		fmt.Println("You're learning Python! You don't use curly braces but indentation!")

	case "Go", "golang":
		fmt.Println("You're learning Go!")
		fmt.Println("Go is a statically typed language.")
		fmt.Println("It is a compiled language.")
		fmt.Println("It is a garbage collected language.")
		fmt.Println("It has no pointers.")
		fmt.Println("It has no classes.")
		fmt.Println("It has no inheritance.")
		fmt.Println("It has no interfaces.")
		fmt.Println("It has no generics.")

	default:
		fmt.Println("Any other programming language is a good start.")
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even number!")
	case n%2 != 0:
		fmt.Println("Odd number!")
	default:
		fmt.Println("Never here!")
	}

	hour := time.Now().Hour()
	// fmt.Println(hour)
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	case hour < 22:
		fmt.Println("Good night!")
	default:
		fmt.Println("Good evening!")
	}

	//** Switch simple statement **//

	// Syntax: statement (n:=10), semicolon and a switch condition
	//(true in this case, we are comparing boolean expressions that return true)
	// we can remove the word "true" because it's the default
	switch n := 10; true {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}

}
