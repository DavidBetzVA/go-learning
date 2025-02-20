package main

import "fmt"

type Component interface {
	Operation() string
}

type Leaf struct {
	name string
}

func (l *Leaf) Operation() string {
	return "Leaf: " + l.name
}

type Composite struct {
	children []Component
}

func (c *Composite) Operation() string {
	result := "Composite containing:\n"
	for _, child := range c.children {
		result += "- " + child.Operation() + "\n"
	}
	return result
}

func (c *Composite) Add(child Component) {
	c.children = append(c.children, child)
}

func main() {
	leaf1 := &Leaf{name: "Leaf 1"}
	leaf2 := &Leaf{name: "Leaf 2"}

	composite := &Composite{}
	composite.Add(leaf1)
	composite.Add(leaf2)

	fmt.Println(composite.Operation())
}
