package main

import "fmt"

func main() {
	title1, author1, year1 := "The Hobbit", "JRR Tolkien", 1937
	title2, author2, year2 := "The Lord of the Rings", "JRR Tolkien", 1954

	fmt.Println("Book1:", title1, author1, year1)
	fmt.Println("Book2:", title2, author2, year2)

	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		title, author string
		year          int
	}

	myBook := book{"The Hobbit", "JRR Tolkien", 1937}
	fmt.Println(myBook)

	bestBook := book{title: "Harry Potter", year: 2000, author: "JK Rowling"}
	_ = bestBook

	aBook := book{title: "Just RandomBook"}
	fmt.Println(aBook)
	fmt.Printf("%#v\n", aBook)
}
