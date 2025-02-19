package main

import "fmt"

func main() {
	// if-else
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less or equal to 5")
	}

	// for loop
	for i := 0; i < 3; i++ {
		fmt.Println("Loop iteration:", i)
	}

	// switch
	switch x {
	case 10:
		fmt.Println("x is 10")
	default:
		fmt.Println("x is not 10")
	}

	// range with index and item
	items := []string{"apple", "banana", "cherry"}
	for i, item := range items {
		fmt.Printf("Index: %d, Item: %s\n", i, item)
	}

	// range with only item
	for _, item := range items {
		fmt.Printf("Item: %s\n", item)
	}

	// range with only index
	for i := range items {
		fmt.Printf("Index: %d\n", i)
	}

	// range by itself
	for range items {
		fmt.Println("Loop iteration")
	}
}
