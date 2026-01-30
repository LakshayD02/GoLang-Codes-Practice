package main

import "fmt"

func main() {
	// --- CONSTANTS ---
	// Rule: Use 'const' instead of 'var'.
	// Rule: You MUST give it a value immediately.
	// Rule: Its value can NEVER change (it is "immutable").

	const Pi = 3.14
	const LanguageName = "Go"
	const MaxScore = 100

	// Printing the constants
	fmt.Println("Language:", LanguageName)
	fmt.Println("Math Constant:", Pi)
	fmt.Println("Maximum possible score:", MaxScore)

	// If you try to do: MaxScore = 110
	// The code will throw an error: "cannot assign to MaxScore"
}
