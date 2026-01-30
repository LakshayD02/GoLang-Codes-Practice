package main

import "fmt"

func main() {

	// ---------------------------------
	// 1. range with a slice
	// ---------------------------------
	fmt.Println("Range with slice:")

	numbers := []int{10, 20, 30}

	// range gives index and value
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}

	// ---------------------------------
	// 2. range ignoring index
	// ---------------------------------
	fmt.Println("\nIgnore index:")

	for _, value := range numbers {
		fmt.Println("Value:", value)
	}

	// ---------------------------------
	// 3. range with a map
	// ---------------------------------
	fmt.Println("\nRange with map:")

	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Tom":   20,
	}

	// range gives key and value
	for name, age := range ages {
		fmt.Println(name, "is", age, "years old")
	}

	// ---------------------------------
	// 4. range with a string
	// ---------------------------------
	fmt.Println("\nRange with string:")

	word := "GoLang"

	// range gives index and character (rune)
	for index, char := range word {
		fmt.Println("Index:", index, "Character:", string(char))
	}

	// ---------------------------------
	// 5. range using only index
	// ---------------------------------
	fmt.Println("\nOnly index:")

	for index := range numbers {
		fmt.Println("Index:", index)
	}
}
