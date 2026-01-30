package main

import "fmt"

// ---------------------------------
// 1. Simple function (no parameters)
// ---------------------------------
func sayHello() {
	fmt.Println("Hello, Go!")
}

// ---------------------------------
// 2. Function with parameters
// ---------------------------------
func add(a int, b int) {
	sum := a + b
	fmt.Println("Sum:", sum)
}

// ---------------------------------
// 3. Function that returns a value
// ---------------------------------
func multiply(a int, b int) int {
	return a * b
}

// ---------------------------------
// 4. Function with multiple return values
// ---------------------------------
func divide(a int, b int) (int, bool) {
	if b == 0 {
		return 0, false // division failed
	}
	return a / b, true
}

// ---------------------------------
// 5. Function using slice parameter
// ---------------------------------
func sumSlice(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {

	// Call simple function
	sayHello()

	// Call function with parameters
	add(5, 3)

	// Call function with return value
	result := multiply(4, 2)
	fmt.Println("Multiply:", result)

	// Call function with multiple returns
	quotient, ok := divide(10, 2)
	if ok {
		fmt.Println("Division result:", quotient)
	} else {
		fmt.Println("Cannot divide by zero")
	}

	// Call function with slice
	numbers := []int{1, 2, 3, 4}
	total := sumSlice(numbers)
	fmt.Println("Sum of slice:", total)
}
