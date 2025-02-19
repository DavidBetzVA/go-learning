// This program demonstrates the use of condition variables in Go for synchronizing
// access to a shared resource (a queue) between multiple producer and consumer goroutines.
// Producers add items to the queue, and consumers process items from the queue.
// The condition variable is used to signal consumers when new items are added to the queue.

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	queue      []int
	queueMutex sync.Mutex
	queueCond  = sync.NewCond(&queueMutex) // Condition variable
)

// Producer function adds items to the queue
func producer(id int) {
	for i := 1; i <= 5; i++ {
		queueMutex.Lock() // Lock the queue before modifying it
		item := i + (id * 10)
		queue = append(queue, item)
		fmt.Printf("Producer %d added item %d\n", id, item)
		queueCond.Signal()                 // Signal waiting consumers that a new item is available
		queueMutex.Unlock()                // Unlock the queue after modification
		time.Sleep(time.Millisecond * 500) // Simulate work
	}
}

// Consumer function processes items from the queue
func consumer(id int) {
	for {
		queueMutex.Lock() // Lock the queue before accessing it
		for len(queue) == 0 {
			queueCond.Wait() // Wait for items to be added to the queue
		}
		item := queue[0]
		queue = queue[1:]
		fmt.Printf("Consumer %d processed item %d\n", id, item)
		queueMutex.Unlock()                // Unlock the queue after accessing it
		time.Sleep(time.Millisecond * 300) // Simulate processing time
	}
}

func main() {
	// Start consumers
	for i := 1; i <= 2; i++ {
		go consumer(i)
	}

	// Start producers
	for i := 1; i <= 2; i++ {
		go producer(i)
	}

	// Let them run for a while before exiting
	time.Sleep(time.Second * 5)
	fmt.Println("All processing complete")
}
