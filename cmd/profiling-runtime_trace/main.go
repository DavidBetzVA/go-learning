package main

import (
	"context"
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

// Simulated work function
func doWork(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	trace.Log(ctx, "doWork", fmt.Sprintf("Worker %d started", id))
	time.Sleep(time.Millisecond * 500) // Simulate work
	trace.Log(ctx, "doWork", fmt.Sprintf("Worker %d finished", id))
}

func main() {
	// Create a trace file
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println("Error creating trace file:", err)
		return
	}
	defer func() {
		_ = f.Close()
	}()

	// Start tracing
	_ = trace.Start(f)
	defer trace.Stop()

	var wg sync.WaitGroup
	ctx := context.Background()

	// Launch multiple workers
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go doWork(ctx, i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers completed. Run 'go tool trace trace.out' to analyze the trace.")
}
