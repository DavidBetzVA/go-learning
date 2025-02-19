package main

import "fmt"

func main() {
	x := 42
	ptr := &x // Pointer to x

	fmt.Println("Value of x:", x)
	fmt.Println("Pointer to x:", ptr)
	fmt.Println("Value at pointer:", *ptr)
}
