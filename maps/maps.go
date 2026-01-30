package main

import (
	"fmt"
	"maps"
)

func main() {

	// -----------------------------
	// 1. Creating a map
	// -----------------------------
	// A map stores data in key-value pairs
	// map[keyType]valueType
	ages := make(map[string]int)

	// -----------------------------
	// 2. Adding values to the map
	// -----------------------------
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 22

	// -----------------------------
	// 3. Reading values from the map
	// -----------------------------
	// Use the key to get its value
	fmt.Println("Alice's age:", ages["Alice"])
	fmt.Println("Bob's age:", ages["Bob"])

	// -----------------------------
	// 4. Updating a value
	// -----------------------------
	// Just assign a new value to the same key
	ages["Alice"] = 26
	fmt.Println("Alice's updated age:", ages["Alice"])

	// -----------------------------
	// 5. Checking if a key exists
	// -----------------------------
	// The second value tells us if the key exists
	//If the key does not exist, the value will be the zero value of the map's value type
	age, exists := ages["David"]
	if exists {
		fmt.Println("David's age:", age)
	} else {
		fmt.Println("David is not in the map")
	}

	// -----------------------------
	// 6. Looping through a map
	// -----------------------------
	// range lets us go through all key-value pairs
	for name, age := range ages {
		fmt.Println(name, "is", age, "years old")
	}

	// -----------------------------
	// 7. Deleting a key from the map
	// -----------------------------
	delete(ages, "Bob")
	fmt.Println("After deleting Bob:", ages)

	// -----------------------------
	// 8. Creating and initializing a map in one line
	// -----------------------------
	countries := map[string]string{
		"US": "United States",
		"IN": "India",
		"FR": "France",
	}

	fmt.Println("Countries map:", countries)

	// -----------------------------
	// 8. Creating and initializing a map in one line
	// -----------------------------
	countries2 := map[string]string{
		"US": "United States",
		"IN": "India",
		"FR": "France",
	}

	fmt.Println("Countries map:", countries2)

	fmt.Println(maps.Equal(countries, countries2))

}
