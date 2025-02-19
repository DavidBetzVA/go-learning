package main

import "fmt"

// Counter returns a function that increments a private variable
func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := Counter()   // Create a new counter
	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
	fmt.Println(counter()) // Output: 3
}
