package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:        "John",
		salary:      1000,
		contactInfo: Contact{email: "john@example.com", address: "123 Main St", phone: 123456789},
	}

	jane := Employee{
		name:        "Jane",
		salary:      2000,
		contactInfo: Contact{email: "jane@example.com", address: "456 Elm St", phone: 987654321},
	}

	fmt.Printf("%+v\n", john)
	fmt.Printf("%+v\n", jane)

	fmt.Printf("Employee's Email: %s\n", john.contactInfo.email)

	john.contactInfo.email = "john.doe@example.com"

	fmt.Printf("Employee's new Email: %s\n", john.contactInfo.email)

	myContact := Contact{email: "andrei@email.com", phone: 21123213, address: "Islamabad, Pakistan"}

	fmt.Println(myContact)
}
