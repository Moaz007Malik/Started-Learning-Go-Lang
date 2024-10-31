package main

import "fmt"

func main() {
	var employees map[string]string
	fmt.Printf("%#v\n", employees)

	fmt.Printf("No of pairs: %d\n", len(employees))

	fmt.Printf("The value for key %q is %q\n", "Dan", employees["Dan"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	var m1 map[[2]int]string
	_ = m1

	// employees["Dan"] = "Programmer"

	people := map[string]float64{}
	people["Dan"] = 10
	people["Maria"] = 20
	people["Paul"] = 30

	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 8

	fmt.Println(map1)

	balances := map[string]float64{
		"USD": 123.12,
		"EUR": 234.21,
		// 234:234.12 // Error: all keys must be the same type
		"GBP": 234.123,
	}
	fmt.Println(balances)

	m := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m

	balances["USD"] = 500.2
	balances["GBP"] = 100
	fmt.Println(balances)

	fmt.Println(balances["RON"])

	balances["RON"] = 0

	v, ok := balances["RON"]

	if ok {
		fmt.Println("The RON balance is:", v)
	} else {
		fmt.Println("No RON balance")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	delete(balances, "USD")
	fmt.Println(balances)

	// Comparing Maps

	a := map[string]string{"A": "X"}
	// b := map[string]string{"B": "Y"} // Now not equal
	b := map[string]string{"A": "X"} // Now Equal
	// invalid operation: a == b (map can only be compared to nil)
	// fmt.Println(a == b)

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	fmt.Println(s1, s2)

	if s1 == s2 {
		fmt.Println("a and b are equal")
	} else {
		fmt.Println("a and b are not equal")
	}

	// Map Header Cloning

	friends := map[string]int{"Dan":40, "Maria":30}
	neighbors := friends
	friends["Dan"] = 50

	fmt.Println(neighbors)
	peoples := make(map[string]int)

	for k, v := range friends {
		peoples[k] = v
	}

	peoples["Anne"] = 18
	fmt.Printf("%#v -------------- %#v\n", peoples, friends)

	delete(friends, "Dan")
	fmt.Println(friends)

	delete(people, "Andrei")
	fmt.Println(people)
}
