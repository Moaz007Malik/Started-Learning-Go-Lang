package main

import "fmt"

func main() {
	type Book struct {
		title  string
		author string
		year   int
	}

	lastBook := Book{title: "Anna Karenina"}
	_ = lastBook
	fmt.Println(lastBook.title)

	// This is an error
	// page := lastBook.pages

	fmt.Printf("%#v\n", lastBook)

	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878
	fmt.Printf("%+v\n", lastBook)

	// aBook := Book{title: "Anna Karenina", author: "", year: 0}
	aBook := Book{title: "Anna Karenina", author: "Leo Tolstoy", year: 1878}

	fmt.Println(aBook == lastBook)

	myBook := aBook
	myBook.year = 2020
	fmt.Println(myBook, aBook)
}
