package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Import for side effects to enable pprof
	"time"
)

// Simulated handler with request tracing
func handler(w http.ResponseWriter, r *http.Request) {
	// Record the start time of the request
	start := time.Now()
	// Log the received request with its URL path
	log.Println("Received request:", r.URL.Path)

	// Simulate some work by sleeping for 500 milliseconds
	time.Sleep(time.Millisecond * 500)

	// Write a response to the client
	_, _ = fmt.Fprintln(w, "Hello, World!")
	// Log the completion of the request with the time taken
	log.Printf("Request %s completed in %v\n", r.URL.Path, time.Since(start))
}

func main() {
	// Register the handler function for the root URL path
	http.HandleFunc("/", handler)

	// Start a separate goroutine to run the pprof server for profiling
	go func() {
		// Log that the pprof server is starting
		log.Println("Starting pprof server on :6060")
		// Start the pprof server on port 6060 and log any errors
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Log that the main HTTP server is starting
	log.Println("Starting HTTP server on :8080")
	// Start the main HTTP server on port 8080 and log any fatal errors
	log.Fatal(http.ListenAndServe(":8080", nil))
}
