package main

import "fmt"

// Function to be benchmarked
func Sum(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += i
	}
	return total
}

// Optimized function to be benchmarked
func SumOptimized(n int) int {
	return n * (n - 1) / 2
}

func main() {
	fmt.Println("Sum of first 100 numbers:", Sum(100))
	fmt.Println("SumOptimized of first 100 numbers:", SumOptimized(100))
}
