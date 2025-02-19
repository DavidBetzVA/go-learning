package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Method 1: Direct struct initialization
	a := Person{Name: "Alice", Age: 30}
	
	// Method 2: Using new (returns a pointer)
	b := new(Person)
	b.Name = "Bob"
	b.Age = 25

	fmt.Println("Struct initialized directly:", a)
	fmt.Println("Struct initialized with new():", *b)
}
