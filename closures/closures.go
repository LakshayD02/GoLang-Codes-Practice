package main

import "fmt"

// counter returns a function (this is a closure)
func counter() func() int {
	count := 0 // This variable is captured by the closure

	// The returned function remembers `count`
	return func() int {
		count++
		return count
	}
}

func main() {

	// Create a closure
	c := counter()

	// Each call updates the same `count`
	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3

	fmt.Println()

	// Another closure with its own memory
	c2 := counter()
	fmt.Println(c2()) // 1 (separate count)
}
