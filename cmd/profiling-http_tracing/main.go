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
		// Start the pprof server on port 6060 and log any errors
		server := &http.Server{
			Addr:         "localhost:6060",
			Handler:      nil,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		}
		log.Println(server.ListenAndServe())
	}()

	// Log that the main HTTP server is starting
	log.Println("Starting HTTP server on :8080")
	// Start the main HTTP server on port 8080 with timeouts and log any fatal errors
	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
