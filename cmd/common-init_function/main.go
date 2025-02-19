package main

import (
	"fmt"
	"sync"
)

var once sync.Once

// Global variable initialization
var globalVar string = initializeGlobal()

func initializeGlobal() string {
	fmt.Println("Global variable: Initialized at package load time")
	return "Initialized in global scope"
}

// init function: Runs automatically before main()
func init() {
	fmt.Println("init() function: Executed before main()")
	globalVar = "Modified by init()"
}

func main() {
	fmt.Println("main() function: Execution begins")

	// Demonstrating sync.Once
	for i := 0; i < 3; i++ {
		once.Do(func() {
			fmt.Println("sync.Once: This should run only once")
		})
		fmt.Println("sync.Once: Attempt", i+1)
	}

	// Final output
	fmt.Println("Final value of globalVar:", globalVar)
}
