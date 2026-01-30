package main

import "fmt"

func main() {
	// Call the switch statement example
	switchExample()
}

// This function demonstrates a simple 'switch' statement in Go.
// A 'switch' allows you to compare a value against multiple cases.
// When a match is found, the corresponding block runs.
// 'default' runs if no cases match.
func switchExample() {
	day := 3

	fmt.Println("Checking the day of the week:")

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// You can also use multiple values in a single case
	grade := "B"

	fmt.Println("\nChecking the grade:")

	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B", "C":
		// This block runs for either B or C
		fmt.Println("Well done")
	case "D":
		fmt.Println("You passed")
	case "F":
		fmt.Println("Better try again")
	default:
		fmt.Println("Invalid grade")
	}
}

// No need to write break statements in Go's switch cases.
