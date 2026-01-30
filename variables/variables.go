package main

import "fmt"

func main() {
	// --- LONG FORM ---
	// Rule: Use 'var' + name + type. Good for being 100% explicit.
	var language string = "golang"
	var version int = 1

	// --- SHORTHAND (The "Walrus" Operator) ---
	// Rule: Use ':=' to declare and assign at once. Go "guesses" the type.
	// Rule: This style ONLY works inside functions.
	rating := 4.8  // Go sees the decimal and makes it a float64
	isEasy := true // Go sees true/false and makes it a bool

	// --- THE "LOCK-IN" RULE ---
	// Rule: Even with shorthand, types are STATIC.
	// You cannot change 'language' to a number later.

	fmt.Println("Topic:", language)
	fmt.Println("Major Version:", version)
	fmt.Println("User Rating:", rating)
	fmt.Println("Is it beginner friendly?", isEasy)
}
