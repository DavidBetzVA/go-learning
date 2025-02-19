package main

import "fmt"

func main() {
	// Array
	arr := [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Loop through array
	for i, v := range arr {
		fmt.Printf("Array index %d: %d\n", i, v)
	}

	// Slice
	slice := []string{"Go", "is", "fun"}
	slice = append(slice, "!")
	fmt.Println("Slice:", slice)

	// Loop through slice
	for i, v := range slice {
		fmt.Printf("Slice index %d: %s\n", i, v)
	}

	// Map
	m := map[string]int{"Alice": 30, "Bob": 25}
	m["Charlie"] = 35
	fmt.Println("Map:", m)

	// Loop through map
	for k, v := range m {
		fmt.Printf("Map key %s: %d\n", k, v)
	}

	// Check if key exists in map
	if age, ok := m["Alice"]; ok {
		fmt.Printf("Alice is %d years old\n", age)
	} else {
		fmt.Println("Alice not found in map")
	}

	if age, ok := m["Eve"]; ok {
		fmt.Printf("Eve is %d years old\n", age)
	} else {
		fmt.Println("Eve not found in map")
	}
}
