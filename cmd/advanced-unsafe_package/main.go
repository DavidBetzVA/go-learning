// Package main demonstrates the use of the unsafe package in Go.
// This example shows how to reinterpret the memory of an integer as a float64.
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Declare an integer variable.
	var i int = 42

	// Use unsafe.Pointer to get a pointer to the integer.
	ptr := unsafe.Pointer(&i)

	// Reinterpret the pointer as a pointer to a float64.
	floatPtr := (*float64)(ptr)

	// Print the original integer value.
	fmt.Println("Original int:", i)

	// Print the reinterpreted value as a float64.
	fmt.Println("Reinterpreted as float64:", *floatPtr)
}
