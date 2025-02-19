// This example demonstrates the use of panic and recover in Go.
// Panic should be used sparingly, typically for unrecoverable errors.
// Recover can be used to regain control of a panicking goroutine, but it should not be used to ignore errors.

package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Before panic")
	panic("Something went wrong!")
	fmt.Println("After panic (this won't execute)")
}
