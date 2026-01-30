package main

import "fmt"

func main() {
	// Call the arrays example
	arraysExample()
}

// This function demonstrates how to use arrays in Go.
//
// An array is a collection of elements of the same type.
// Arrays have a fixed size, and each element can be accessed using an index.
func arraysExample() {

	// Checking the length of an array
	fmt.Println("=== Arrays Example ===")
	var num [5]int
	fmt.Println("Length of num array:", len(num))

	// Declare an array of 5 integers
	var numbers [5]int

	// Assign values to each element
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	// Print the whole array
	fmt.Println("Array:", numbers)

	// Access and print individual elements using their index
	fmt.Println("First element:", numbers[0])
	fmt.Println("Third element:", numbers[2])

	// You can also declare and initialize an array in one line
	colors := [3]string{"Red", "Green", "Blue"}
	fmt.Println("Colors array:", colors)

	// Loop through an array using a standard for loop
	fmt.Println("\nLooping through numbers array using for loop:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index:", i, "Value:", numbers[i])
	}

	// Loop through an array using range
	fmt.Println("\nLooping through colors array using range:")
	for index, color := range colors {
		fmt.Println("Index:", index, "Color:", color)
	}
}

// Zeroed arrays in Go have default values:
// - Numeric types default to 0 (e.g., int, float64)
// - Boolean types default to false
// - String types default to "" (empty string)
// - Other types have their respective zero values
