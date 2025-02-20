// This example demonstrates how to use pprof to dump CPU and memory profiles from a Go application.
// It creates CPU and memory profile files and writes profiling data to them.
// To view and interact with the profiles using a web-based UI, you can start an HTTP server with pprof
// and then access the profiles through a web browser.
//
// To start the pprof server, run the following command in the directory containing the profile files:
// go tool pprof -http=:8080 cpu.prof
// go tool pprof -http=:8080 mem.prof
//
// Then visit http://localhost:8080 in your browser to view the profiles.

package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"
	"time"
)

func main() {
	// Create CPU profile file
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatalf("could not create CPU profile: %v", err)
	}
	defer func() {
		_ = cpuFile.Close()
	}()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatalf("could not start CPU profile: %v", err)
	}
	defer pprof.StopCPUProfile()

	// Simulate some work
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
	}

	// Create memory profile file
	memFile, err := os.Create("mem.prof")
	if err != nil {
		log.Fatalf("could not create memory profile: %v", err)
	}
	defer func() {
		_ = memFile.Close()
	}()

	// Write memory profile
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		log.Fatalf("could not write memory profile: %v", err)
	}

	// Handle graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("Application exiting")
}
