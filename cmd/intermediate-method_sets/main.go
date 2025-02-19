package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func main() {
	var s Speaker
	p := Person{Name: "Charlie"}
	s = p
	fmt.Println(s.Speak())
}
