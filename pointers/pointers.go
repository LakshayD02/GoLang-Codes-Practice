package main

import "fmt"

// updateValue takes a pointer to an integer
// This allows the function to change the original value
func updateValue(num *int) {

	// *num means "value stored at the address"
	// This changes the original variable
	*num = 100
}

// Person is a simple struct
type Person struct {
	name string
	age  int
}

func main() {

	// ------------------------------------------------
	// BASIC POINTER EXAMPLE
	// ------------------------------------------------

	// Create a normal integer variable
	x := 10

	// &x gives the memory address of x
	// p now stores the address of x
	p := &x

	fmt.Println("BASIC POINTER EXAMPLE")
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of p (address of x):", p)

	// *p means "go to the address stored in p and get the value"
	fmt.Println("Value at address p:", *p)

	// Change the value of x using the pointer
	*p = 20
	fmt.Println("Value of x after change using pointer:", x)

	fmt.Println()

	// ------------------------------------------------
	// POINTER WITH FUNCTION
	// ------------------------------------------------

	// Create another variable
	y := 5
	fmt.Println("POINTER WITH FUNCTION EXAMPLE")
	fmt.Println("Value of y before function call:", y)

	// Pass the address of y to the function
	updateValue(&y)

	// y is changed because the function used a pointer
	fmt.Println("Value of y after function call:", y)

	fmt.Println()

	// ------------------------------------------------
	// POINTER TO STRUCT
	// ------------------------------------------------

	// Create a struct variable
	person := Person{
		name: "Alice",
		age:  25,
	}

	fmt.Println("POINTER TO STRUCT EXAMPLE")
	fmt.Println("Before update:", person)

	// Create a pointer to the struct
	ptr := &person

	// Update struct field using pointer
	// Go automatically handles (*ptr).age for us
	ptr.age = 30

	// Original struct is updated
	fmt.Println("After update:", person)

	fmt.Println()

	// ------------------------------------------------
	// NIL POINTER EXAMPLE
	// ------------------------------------------------

	// Declare a pointer without assigning an address
	var n *int

	fmt.Println("NIL POINTER EXAMPLE")

	// Always check if pointer is nil before using it
	if n == nil {
		fmt.Println("Pointer n is nil, cannot dereference it")
	}

	// Using *n here would cause a runtime panic
}
