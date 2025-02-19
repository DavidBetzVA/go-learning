// Package main demonstrates the use of worker pools for concurrent processing in Go.
// This example shows how to create a pool of workers to process jobs concurrently.
package main

import (
	"fmt"
	"sync"
	"time"
)

// worker is a function that processes jobs from the jobs channel and sends results to the results channel.
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		// Print the worker ID and the job being processed.
		fmt.Printf("Worker %d processing job %d\n", id, job)
		// Simulate work by sleeping for a second.
		time.Sleep(time.Second)
		// Send the result of the job to the results channel.
		results <- job * 2
	}
}

func main() {
	// Define the number of workers.
	const numWorkers = 3

	// Create channels for jobs and results.
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Use a WaitGroup to wait for all workers to finish.
	var wg sync.WaitGroup

	// Start the workers.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results)
		}(i)
	}

	// Send jobs to the jobs channel.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	// Close the jobs channel to indicate no more jobs will be sent.
	close(jobs)

	// Wait for all workers to finish.
	wg.Wait()
	// Close the results channel to indicate no more results will be sent.
	close(results)

	// Print the results.
	for res := range results {
		fmt.Println("Result:", res)
	}
}
