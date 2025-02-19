// This program demonstrates the use of Go's concurrency features, specifically atomic operations.
// It shows how to use atomic functions to safely increment and read a shared counter across multiple goroutines.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64 // Atomic counter

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&counter, 1) // Atomic increment
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter Value:", atomic.LoadInt64(&counter)) // Atomic read
}
