package main

import "fmt"

func main() {
	var cities []string
	fmt.Println(cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities))

	// Error: We cannot add values in the slices like this
	// cities[0] = "Paris"

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("numbers: %v\n", numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	type names []string
	friends := names{"John", "Maria"}
	fmt.Printf("friends: %v\n", friends)
	myfriends := friends[0]
	fmt.Println(myfriends)

	friends[0] = "Paul"
	fmt.Printf("friends: %v\n", friends)

	for index, value := range friends {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	var n []int
	n = numbers
	fmt.Printf("n: %v\n", n)
}
