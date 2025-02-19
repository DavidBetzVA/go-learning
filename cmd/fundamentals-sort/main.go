package main

import (
	"fmt"
	"sort"
)

// Person defines a custom type with Name and Age fields
type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface for []Person based on the Age field
type ByAge []Person

// Len returns the number of elements in the collection
func (a ByAge) Len() int { return len(a) }

// Swap swaps the elements with indexes i and j
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less reports whether the element with index i should sort before the element with index j
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	// Sorting a slice of integers
	numbers := []int{5, 3, 8, 1}
	sort.Ints(numbers)
	fmt.Println(numbers) // Output: [1 3 5 8]

	// Sorting a slice of Person by Age
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Diana", 28},
	}
	sort.Sort(ByAge(people))
	fmt.Println(people) // Output: [{Bob 25} {Diana 28} {Alice 30} {Charlie 35}]
}
