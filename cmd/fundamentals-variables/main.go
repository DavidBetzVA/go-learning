package main

import "fmt"

func main() {
	// Variable declaration
	var message string = "Hello, Go!"
	count := 3 // Short declaration

	// Constants
	const version = 1.0

	fmt.Println(message)
	fmt.Printf("Count: %d, Version: %.1f\n", count, version)
}
