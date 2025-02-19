package main

import (
	"fmt"
)

// Numeric is a generic constraint that allows only numeric types (int, float32, etc.).
type Numeric interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Add adds two values of a type that supports addition (string or numeric types).
func Add[T string | Numeric](a, b T) T {
	return a + b
}

// Min finds the minimum of two values of a type that satisfies the Numeric constraint.
func Min[T Numeric](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example with integers
	intA, intB := 10, 20
	fmt.Printf("Add(%d, %d) = %d\n", intA, intB, Add(intA, intB))
	fmt.Printf("Min(%d, %d) = %d\n", intA, intB, Min(intA, intB))

	// Example with floats
	floatA, floatB := 5.5, 3.2
	fmt.Printf("Add(%.1f, %.1f) = %.1f\n", floatA, floatB, Add(floatA, floatB))
	fmt.Printf("Min(%.1f, %.1f) = %.1f\n", floatA, floatB, Min(floatA, floatB))

	// Example with strings (Only works for Add, since Min requires Numeric)
	strA, strB := "Hello, ", "World!"
	fmt.Printf("Add(%q, %q) = %q\n", strA, strB, Add(strA, strB))

	// Uncommenting the below line will cause a compile-time error because strings are not Numeric
	// fmt.Println(Min("apple", "banana"))
}
