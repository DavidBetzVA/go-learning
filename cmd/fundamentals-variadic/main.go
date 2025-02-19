package main

import "fmt"

// Sum takes multiple integers and returns their sum
func Sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Sum:", Sum(1, 2, 3, 4, 5))   // Output: 15
	fmt.Println("Sum with no values:", Sum()) // Output: 0
}
