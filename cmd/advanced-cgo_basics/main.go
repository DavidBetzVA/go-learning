// This file demonstrates how to use C code in Go using CGO.
// CGO allows Go packages to call C code.

// Package main is the entry point for the Go program.
package main

/*
#include <stdio.h>
#include <time.h>

// getCurrentTime is a C function that returns the current time as a string.
const char* getCurrentTime() {
    time_t t;
    time(&t);
    return ctime(&t);
}
*/
import "C"
import "fmt"

// main is the entry point for the Go program.
func main() {
	currentTime := C.getCurrentTime()
	fmt.Printf("Current server time: %s", C.GoString(currentTime))
}
