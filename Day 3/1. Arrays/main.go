package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Numbers: %#v\n", numbers)

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4

	fmt.Printf("Numbers: %v\n", numbers)

	var a1 = [4]float64{1, 2, 3, 4}
	fmt.Printf("%v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%v\n", a2)

	a3 := [4]string{"lars", "paul", "jonh", "diana"}
	fmt.Printf("%v\n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%v\n", a4)
	fmt.Printf("%#v\n", a4)

	fmt.Println(strings.Repeat("#", 20))
	// ellipsis operator
	a5 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%v\n", a5)
	fmt.Printf("%#v\n", a5)
	fmt.Printf("the length of a5 is %d\n", len(a5))

	fmt.Println(strings.Repeat("#", 20))
	a6 := [...]int{1,
		2,
		3,
		4,
		5, // comma is mandatory
	}
	fmt.Printf("%v\n", a6)

	fmt.Println(strings.Repeat("#", 20))
	// Arrays Operators

	numbers2 := [4]int{1, 2, 3, 4}
	fmt.Printf("numbers2: %v\n", numbers2)

	numbers2[0] = 10
	fmt.Printf("numbers2 after modification: %v\n", numbers2)

	//Error: There is no index 5 in this array
	// numbers2[5] = 100

	for i, v := range numbers2 {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	// Practice

	myArray := [5]string{"apple", "banana", "cherry", "date", "elderberry"}
	for i, v := range myArray {
		fmt.Println("index: ", i, "items:", v)
	}

	fmt.Println(strings.Repeat("#", 20))

	for i := 0; i < len(numbers2); i++ {
		fmt.Println("index: ", i, "items:", numbers2[i])
	}

	balances := [2][3]int{
		{5, 6, 7},
		[3]int{1, 2, 3},

		// OR

		//{5, 6, 7},
		//{1, 2, 3},
	}
	fmt.Println(balances)

	m := [3]int{1, 2, 3}
	n := m

	fmt.Println(m, n)
	fmt.Println("n is equal to m:", n == m)
	m[0] = 10
	fmt.Println("n is equal to m:", n == m)

	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades)

	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	names := [...]string{
		5: "Dan",
	}
	fmt.Println(names)
	fmt.Println(len(names))

	cities := [...]string{
		5: "Paris",
		"Londnn", // index: 6
		1: "New York",
	}
	fmt.Println(cities)
	fmt.Printf("%#v\n", cities)

	weekends := [...]bool{5: true, 6: true}
	fmt.Println(weekends)
}
