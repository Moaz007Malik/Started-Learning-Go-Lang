package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// converting string to int:
	i, err := strconv.Atoi("45")

	// error handling
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)

	}

	// simple (short) statement ->  the same effect as the above code
	// i and err are variables scoped to the if statement only
	if i, err := strconv.Atoi("34"); err == nil {
		fmt.Println("No error. i is ", i)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required!")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer! Error: ", err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)/1.609)
	}

}
