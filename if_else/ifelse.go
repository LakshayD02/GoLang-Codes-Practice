package main

import "fmt"

func main() {
	// Call the if-else example
	ifElseExample()
}

// This function demonstrates a simple 'if-else' statement in Go.
// 'if' checks a condition. If the condition is true, the code inside the 'if' block runs.
// 'else' runs when the 'if' condition is false.
func ifElseExample() {
	number := 7

	fmt.Println("Checking if the number is even or odd:")

	// Check if the number is divisible by 2
	if number%2 == 0 {
		// This block runs if the number is even
		fmt.Println(number, "is even")
	} else {
		// This block runs if the number is NOT even (i.e., odd)
		fmt.Println(number, "is odd")
	}

	// You can also chain multiple conditions using 'else if'
	fmt.Println("Checking the size of the number:")

	if number < 5 {
		fmt.Println(number, "is less than 5")
	} else if number == 5 {
		fmt.Println(number, "is exactly 5")
	} else {
		fmt.Println(number, "is greater than 5")
	}

	fmt.Println("Declaring variable inside the if-else")

	if num := 2; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

}
