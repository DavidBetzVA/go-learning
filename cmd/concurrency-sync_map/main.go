// This program demonstrates the use of Go's concurrency features, specifically the use of sync.Map for concurrent access to a map.
// It shows how to store, load, iterate, and delete key-value pairs in a sync.Map.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var sharedMap sync.Map

	// Store values in sync.Map
	sharedMap.Store("name", "Alice")
	sharedMap.Store("age", 30)

	// Load values
	if value, ok := sharedMap.Load("name"); ok {
		fmt.Println("Name:", value)
	}

	// Iterating over sync.Map
	sharedMap.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // Continue iteration
	})

	// Delete a key
	sharedMap.Delete("age")

	// Check if key exists
	_, exists := sharedMap.Load("age")
	fmt.Println("Key 'age' exists after deletion:", exists)
}
