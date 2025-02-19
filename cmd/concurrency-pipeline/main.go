// This program demonstrates the use of Go's concurrency features, specifically goroutines and channels.
// It implements a simple pipeline with three stages:
// 1. Generate numbers and send them into a channel.
// 2. Square the numbers received from the input channel and send the results to an output channel.
// 3. Print the squared numbers received from the output channel.

package main

import (
	"fmt"
	"time"
)

// Stage 1: Generate numbers and send them into a channel
func generateNumbers(out chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Generating:", i)
		out <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(out)
}

// Stage 2: Square numbers received from input channel and send to output channel
func squareNumbers(in <-chan int, out chan<- int) {
	for num := range in {
		squared := num * num
		fmt.Println("Squaring:", num, "->", squared)
		out <- squared
	}
	close(out)
}

// Stage 3: Print the squared numbers
func printResults(in <-chan int) {
	for result := range in {
		fmt.Println("Result:", result)
	}
}

func main() {
	// Create channels for each stage
	numChannel := make(chan int)
	squareChannel := make(chan int)

	// Start pipeline stages in separate goroutines
	go generateNumbers(numChannel)
	go squareNumbers(numChannel, squareChannel)

	// Print final results (blocking call to ensure main waits)
	printResults(squareChannel)
}
