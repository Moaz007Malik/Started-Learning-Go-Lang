package main

import (
	"fmt"
)

// Basic Pointer Example
func basicPointerExample() {
	fmt.Println("Basic Pointer Example:")

	// Declare an integer variable
	x := 10
	// Get the address of x using the & operator
	p := &x

	fmt.Printf("x: %d, Address of x (p): %p\n", x, p)

	// Dereference p to get the value of x
	fmt.Printf("Value at address p (*p): %d\n", *p)

	// Modify the value at the address pointed to by p
	*p = 20
	fmt.Printf("Updated value of x after *p = 20: %d\n", x)
}

////////////////////////////////////////

// Pointers with Functions
func updateValue(val *int) {
	// Dereference the pointer to change the original value
	*val = 50
}

func pointerWithFunctions() {
	fmt.Println("\nPointers with Functions Example:")

	x := 10
	fmt.Printf("Initial value of x: %d\n", x)

	// Pass the address of x to the function
	updateValue(&x)
	fmt.Printf("Value of x after calling updateValue(&x): %d\n", x)
}

////////////////////////////////////////

// Using Pointers to Avoid Copying Large Structures
type LargeStruct struct {
	values [1000]int // Simulating a large data structure
}

func modifyLargeStruct(ls *LargeStruct) {
	// Modify the first element to show we can change it via pointer
	ls.values[0] = 99
}

func avoidCopyingLargeStruct() {
	fmt.Println("\nAvoiding Copying Large Structs Example:")

	// Create an instance of LargeStruct
	largeStruct := LargeStruct{}

	fmt.Printf("Initial first value of largeStruct.values[0]: %d\n", largeStruct.values[0])

	// Pass a pointer to avoid copying the entire struct
	modifyLargeStruct(&largeStruct)
	fmt.Printf("First value of largeStruct.values[0] after modification: %d\n", largeStruct.values[0])
}

////////////////////////////////////////

// Pointers with Structs and Methods
type Person struct {
	Name string
	Age  int
}

// Method with a pointer receiver to modify the struct
func (p *Person) updateName(newName string) {
	p.Name = newName // Modifies the Name field directly
}

func pointerWithStructs() {
	fmt.Println("\nPointers with Structs Example:")

	// Create an instance of Person
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("Original Person: %+v\n", person)

	// Call method with pointer receiver to update name
	person.updateName("Bob")
	fmt.Printf("Updated Person: %+v\n", person)
}

////////////////////////////////////////

// Using nil Pointers
func nilPointerExample() {
	fmt.Println("\nNil Pointer Example:")

	// Declare a pointer with no initial value, so it points to nil
	var p *int
	if p == nil {
		fmt.Println("p is a nil pointer.")
	}

	// Uncommenting the following line would cause a runtime panic:
	// *p = 10 // Dereferencing nil causes an error
}

////////////////////////////////////////

// Main function to run all examples
func main() {
	basicPointerExample()     // Basic pointer usage
	pointerWithFunctions()    // Passing pointers to functions
	avoidCopyingLargeStruct() // Avoid copying large structs with pointers
	pointerWithStructs()      // Modifying struct fields through pointer methods
	nilPointerExample()       // Handling nil pointers safely
}
