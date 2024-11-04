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

	// Anonymous structs

	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}
	fmt.Printf("%#v\n", diana)
	fmt.Printf("Diana's age is %d\n", diana.age)

	type book struct {
		string
		float64
		bool
	}

	b1 := book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)

	type Employee struct {
		name string
		salary int
		bool
	}

	// e := Employee{name: "John", salary: 1000, bool: true}
	e := Employee{"John", 1000, false}
	fmt.Printf("%#v\n", e)

	e.bool = true

	fmt.Printf("%#v\n", e)
}
