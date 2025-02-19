// This program demonstrates the use of Go's concurrency features, specifically the use of RWMutex for read-write locking.
// It implements a simple example with multiple readers and writers accessing a shared counter.
// Readers acquire a read lock to read the counter value, while writers acquire a write lock to increment the counter.

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	rwMutex sync.RWMutex
)

func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.RLock() // Lock for reading
	fmt.Printf("Reader %d sees counter: %d\n", id, counter)
	time.Sleep(time.Millisecond * 500)
	rwMutex.RUnlock()
}

func writer(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.Lock() // Lock for writing
	counter++
	fmt.Printf("Writer %d incremented counter to: %d\n", id, counter)
	time.Sleep(time.Millisecond * 500)
	rwMutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	// Start multiple readers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go reader(i, &wg)
	}

	// Start multiple writers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go writer(i, &wg)
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}
