package main

import "fmt"

func main() {
	// Defining and executing an anonymous function immediately
	func() {
		fmt.Println("Hello from an anonymous function!")
	}()

	// Assigning an anonymous function to a variable
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}
	greet("Alice")
}
