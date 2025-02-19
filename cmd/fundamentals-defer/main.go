package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Example 1: Closing a file
func exampleCloseFile() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Ensure the file is closed when the function exits
	defer file.Close()

	// ...existing code to work with the file...
	fmt.Println("File opened successfully")
}

// Example 2: Unlocking a mutex
func exampleUnlockMutex() {
	var mu sync.Mutex
	mu.Lock()
	// Ensure the mutex is unlocked when the function exits
	defer mu.Unlock()

	// ...existing code that requires the mutex to be locked...
	fmt.Println("Mutex locked")
}

// Example 3: Releasing a resource
func exampleReleaseResource() {
	resource := acquireResource()
	// Ensure the resource is released when the function exits
	defer releaseResource(resource)

	// ...existing code that uses the resource...
	fmt.Println("Resource acquired")
}

// Example 4: Executing cleanup code
func exampleCleanup() {
	// Ensure cleanup code is executed when the function exits
	defer func() {
		fmt.Println("Cleanup code executed")
	}()

	// ...existing code that may need cleanup...
	fmt.Println("Function body executed")
}

// Example 5: Measuring execution time
func exampleMeasureTime() {
	start := time.Now()
	// Ensure the execution time is measured when the function exits
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("Function took %s\n", elapsed)
	}()

	// ...existing code whose execution time is to be measured...
	time.Sleep(2 * time.Second) // Simulate a long-running task
	fmt.Println("Function body executed")
}

func main() {
	exampleCloseFile()
	exampleUnlockMutex()
	exampleReleaseResource()
	exampleCleanup()
	exampleMeasureTime()
}

// Dummy functions for the sake of example
func acquireResource() string {
	return "resource"
}

func releaseResource(resource string) {
	fmt.Println("Resource released:", resource)
}
