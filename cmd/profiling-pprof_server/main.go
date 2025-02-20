// 1. Start the application by running `go run main.go`.
// 2. Access the pprof profiling data by navigating to `http://localhost:6060/debug/pprof/` in your web browser.
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/gorilla/mux"
)

func doWork() {
	var result int
	for i := 0; i < 1e7; i++ {
		result += i
	}
	fmt.Println("Work done:", result)
}

func main() {
	// Start pprof server
	go func() {
		log.Println("Starting pprof server on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Create a new router
	r := mux.NewRouter()

	// Define a simple handler
	r.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		go doWork()
		_, _ = w.Write([]byte("Work started"))
	}).Methods("GET")

	// Start the HTTP server
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server on :8080")
	log.Fatal(srv.ListenAndServe())
}
