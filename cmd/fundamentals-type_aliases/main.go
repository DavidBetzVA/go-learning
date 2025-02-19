package main

import "fmt"

// Define a new type alias 'MyInt' for the built-in type 'int'
type MyInt int

// Define a new type alias 'StringAlias' for the built-in type 'string'
type StringAlias string

func main() {
	// Use the type alias 'MyInt'
	var number MyInt = 10
	fmt.Printf("Type: %T, Value: %v\n", number, number)

	// Use the type alias 'StringAlias'
	var text StringAlias = "Hello, Go!"
	fmt.Printf("Type: %T, Value: %v\n", text, text)
}
