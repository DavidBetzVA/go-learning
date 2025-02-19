// This program demonstrates the use of Go's concurrency features, specifically the use of WaitGroup and Mutex.
// It shows how to synchronize multiple goroutines using a WaitGroup and protect a shared counter using a Mutex.

package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
