// This file demonstrates how to use Go assembly code.

// Package main is the entry point for the Go program.
package main

// Add is a function implemented in Go assembly that adds two integers.
func Add(a, b int) int

// main is the entry point for the Go program.
func main() {
	// Call the Add function and print the result.
	result := Add(3, 4)
	println("3 + 4 =", result)
}
