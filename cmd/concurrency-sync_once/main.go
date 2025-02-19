package main

import (
	"fmt"
	"sync"
)

// Global sync.Once object to ensure the function runs only once
var once sync.Once

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	once.Do(func() {
		fmt.Println("Initializing... This should only print once.")
	}) // This ensures the code is only run once
	fmt.Println("Worker executing")
}

func main() {
	var wg sync.WaitGroup

	// Launch multiple goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	fmt.Println("All workers finished")
}
