package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Outer x:", x)

	{
		x := 20 // Shadowing x
		fmt.Println("Inner x:", x)
	}

	fmt.Println("Outer x after block:", x)
}
