package main

import (
	"fmt"
)

// Generic constraint with tilde (~)
type Number interface {
	~int | ~float64 // Allows underlying types like custom types based on int/float64
}

// Function using generics
func Add[T Number](a, b T) T {
	return a + b
}

// Custom type
type MyInt int

func main() {
	fmt.Println("Sum of ints:", Add(3, 5))
	fmt.Println("Sum of floats:", Add(2.5, 3.5))

	var x MyInt = 10
	var y MyInt = 20
	fmt.Println("Sum of MyInt:", Add(x, y)) // Works because MyInt is based on int
}
