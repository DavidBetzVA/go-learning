package main

import (
	"fmt"
	"time"
)

// sendMessage sends a message to the provided channel
func sendMessage(ch chan string) {
	ch <- "Hello from channel!"
}

// sendNumbers sends a sequence of numbers to the provided channel
func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func main() {
	// Unbuffered channel example
	ch := make(chan string)
	go sendMessage(ch)
	fmt.Println(<-ch)

	// Buffered channel example
	buffered := make(chan int, 2)
	buffered <- 1
	buffered <- 2
	fmt.Println(<-buffered, <-buffered)

	// Range over channel example
	numCh := make(chan int)
	go sendNumbers(numCh)
	for num := range numCh {
		fmt.Println(num)
	}

	// Select statement example
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
