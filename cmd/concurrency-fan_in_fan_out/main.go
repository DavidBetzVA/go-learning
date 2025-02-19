// This program demonstrates the use of concurrency in Go using goroutines and channels.
// It includes a producer function that generates integers and sends them to a channel,
// and a fanIn function that merges multiple channels into a single channel.

package main

import (
	"fmt"
	"time"
)

// producer generates integers and sends them to the provided channel.
// The id parameter is used to differentiate between different producers.
func producer(id int, out chan<- int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 300)
		out <- i + (id * 10)
	}
}

// fanIn merges multiple input channels into a single output channel.
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	for _, ch := range channels {
		go func(c <-chan int) {
			for val := range c {
				out <- val
			}
		}(ch)
	}
	return out
}

func main() {
	// Create two channels for the producers.
	ch1, ch2 := make(chan int), make(chan int)

	// Start two producer goroutines.
	go producer(1, ch1)
	go producer(2, ch2)

	// Merge the output of the two channels.
	merged := fanIn(ch1, ch2)

	// Receive and print values from the merged channel.
	for i := 0; i < 10; i++ {
		fmt.Println("Received:", <-merged)
	}
}
