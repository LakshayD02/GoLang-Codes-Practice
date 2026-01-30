package main

import "fmt"

func main() {
	// Call the operators example
	operatorsExample()
}

// There are no ternary operators in Go.
// This function demonstrates different types of operators in Go:
// 1. Arithmetic operators: +, -, *, /, %
// 2. Comparison operators: ==, !=, >, <, >=, <=
// 3. Logical operators: &&, ||, !
func operatorsExample() {
	fmt.Println("=== Arithmetic Operators ===")
	a := 10
	b := 3

	fmt.Println("a =", a, ", b =", b)
	fmt.Println("a + b =", a+b) // Addition
	fmt.Println("a - b =", a-b) // Subtraction
	fmt.Println("a * b =", a*b) // Multiplication
	fmt.Println("a / b =", a/b) // Division
	fmt.Println("a % b =", a%b) // Remainder

	fmt.Println("\n=== Comparison Operators ===")
	fmt.Println("a == b:", a == b) // Equal to?
	fmt.Println("a != b:", a != b) // Not equal to?
	fmt.Println("a > b:", a > b)   // Greater than?
	fmt.Println("a < b:", a < b)   // Less than?
	fmt.Println("a >= b:", a >= b) // Greater than or equal to?
	fmt.Println("a <= b:", a <= b) // Less than or equal to?

	fmt.Println("\n=== Logical Operators ===")
	x := true
	y := false

	fmt.Println("x =", x, ", y =", y)
	fmt.Println("x && y:", x && y) // Logical AND
	fmt.Println("x || y:", x || y) // Logical OR
	fmt.Println("!x:", !x)         // Logical NOT
	fmt.Println("!y:", !y)
}
