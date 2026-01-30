package main

import "fmt"

func main() {
	// Call the slices example
	slicesExample()
}

// This function demonstrates how to use slices in Go.
// A slice is similar to an array but more flexible:
// 1. Slices can change size (grow or shrink).
// 2. Slices are references to arrays, so modifying a slice changes the underlying array.
func slicesExample() {
	// Declare and initialize a slice of integers
	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println("Numbers slice:", numbers)

	// Access individual elements using index
	fmt.Println("First element:", numbers[0])
	fmt.Println("Third element:", numbers[2])

	// Add a new element using append
	numbers = append(numbers, 60)
	fmt.Println("After append:", numbers)

	// Slice a slice (get a portion)
	subSlice := numbers[1:4] // elements at index 1,2,3
	fmt.Println("Sub-slice (index 1 to 3):", subSlice)

	// Loop through a slice using standard for loop
	fmt.Println("\nLooping through numbers slice using for loop:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index:", i, "Value:", numbers[i])
	}

	// Loop through a slice using range
	fmt.Println("\nLooping through numbers slice using range:")
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}

	// 1. 'make' creates a slice with a specified length and optional capacity.
	// 2. 'len' returns the current number of elements in the slice.
	// 3. 'cap' returns the total capacity (how many elements it can hold before growing).

	// Create a slice of integers with length 3 and capacity 5
	num := make([]int, 3, 5)

	fmt.Println("Initial slice:", num)
	fmt.Println("Length:", len(num))   // Current number of elements
	fmt.Println("Capacity:", cap(num)) // Maximum elements before resizing

	// Assign values to the slice
	num[0] = 10
	num[1] = 20
	num[2] = 30
	fmt.Println("After assigning values:", num)

	// Add elements using append
	num = append(num, 40)
	num = append(num, 50)
	fmt.Println("After appending 40 and 50:", num)
	fmt.Println("Length:", len(num))   // Updated length
	fmt.Println("Capacity:", cap(num)) // Capacity remains 5 (no resize yet)

	// Append one more element to exceed capacity
	num = append(num, 60)
	fmt.Println("After appending 60 (exceeds initial capacity):", num)
	fmt.Println("Length:", len(num))
	fmt.Println("Capacity:", cap(num)) // Capacity automatically increases
}
