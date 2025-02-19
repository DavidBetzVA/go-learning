package main

import "fmt"

func main() {
	// Creating a slice
	slice := []int{1, 2, 3}
	fmt.Println("Initial slice:", slice, "len:", len(slice), "cap:", cap(slice))

	// Appending elements
	slice = append(slice, 4, 5)
	fmt.Println("After append:", slice, "len:", len(slice), "cap:", cap(slice))

	// Copying a slice
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	fmt.Println("Copied slice:", newSlice)

	// Slicing a slice
	subSlice := slice[1:4] // Includes index 1 to 3
	fmt.Println("Sub-slice (1:4):", subSlice)

	// Checking slice capacity growth
	slice = append(slice, 6, 7, 8, 9)
	fmt.Println("Extended slice:", slice, "len:", len(slice), "cap:", cap(slice))
}
