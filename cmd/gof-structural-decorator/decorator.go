package main

import "fmt"

type Component interface {
	Operation() string
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent: Operation"
}

type Decorator struct {
	component Component
}

func (d *Decorator) Operation() string {
	return d.component.Operation() + " + Decorator"
}

func main() {
	component := &ConcreteComponent{}
	decorator := &Decorator{component}
	fmt.Println(decorator.Operation())
}
