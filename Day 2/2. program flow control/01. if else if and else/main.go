package main

import "fmt"

func main() {
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Price is high")
	} else if price > 50 {
		fmt.Println("Price is medium")
	} else {
		fmt.Println("Price is low")
	}

	if inStock {
		fmt.Println("In Stock")
	} else {
		fmt.Println("Out of Stock")
	}

	if price <= 100 && inStock == true {
		fmt.Println("Buy it")
	}

	if price < 100 {
		fmt.Println("Its cheap")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("Its expensive")
	}

	age := 7
	if age >= 0 && age < 18 {
		fmt.Printf("You cannot vote! Please return in %d years", 18-age)
	} else if age == 18 {
		fmt.Println("This is your first vote")
	} else if age > 18 && age <= 100 {
		fmt.Println("Please  vote, it's imporant")
	} else {
		fmt.Println("Invalid age")
	}
}
