// This program demonstrates the use of goroutines in Go for concurrent execution.
// It includes a function that prints a message multiple times with a delay,
// and shows how to run this function concurrently using a goroutine.

package main

import (
	"fmt"
	"time"
)

// printMessage prints the provided message three times with a delay.
func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// Start the printMessage function as a goroutine.
	go printMessage("Hello from Goroutine!")

	// Continue execution in the main function.
	fmt.Println("Main function execution continues...")

	// Wait for the goroutine to finish.
	time.Sleep(time.Second * 2)
}
