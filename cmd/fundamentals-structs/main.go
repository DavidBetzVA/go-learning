package main

import "fmt"

// Person struct with Name and Age fields
type Person struct {
	Name string
	Age  int
}

// Greet method for Person
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	p.Greet()
}
