package main

import "fmt"

// Variadic functions are functions that can take a variable number of arguments.
// sum is a variadic function
// It can take any number of integers as arguments
func sum(numbers ...int) int {
	total := 0

	// Loop through all the numbers passed to the function
	for _, num := range numbers {
		total += num
	}

	return total
}

func main() {
	// Calling the variadic function with different numbers of arguments
	fmt.Println(sum(1, 2))       // Output: 3
	fmt.Println(sum(1, 2, 3, 4)) // Output: 10
	fmt.Println(sum(10, 20, 30)) // Output: 60
}
