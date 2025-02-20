package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type ConcretePrototype struct {
	field string
}

func (p *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{field: p.field}
}

func main() {
	original := &ConcretePrototype{field: "Prototype Example"}
	clone := original.Clone().(*ConcretePrototype)

	fmt.Println("Original:", original.field)
	fmt.Println("Clone:", clone.field)
}
