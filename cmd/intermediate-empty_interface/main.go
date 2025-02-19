package main

import "fmt"

func printAnything(value interface{}) {
	fmt.Println("Received:", value)
}

func main() {
	printAnything(42)
	printAnything("Hello")
}
