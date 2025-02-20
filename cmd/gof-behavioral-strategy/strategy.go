package main

import "fmt"

type Strategy interface {
	Execute(a, b int) int
}

type AddStrategy struct{}

func (s *AddStrategy) Execute(a, b int) int {
	return a + b
}

type SubtractStrategy struct{}

func (s *SubtractStrategy) Execute(a, b int) int {
	return a - b
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func main() {
	context := &Context{}

	addStrategy := &AddStrategy{}
	context.SetStrategy(addStrategy)
	fmt.Println("Addition:", context.ExecuteStrategy(5, 3))

	subtractStrategy := &SubtractStrategy{}
	context.SetStrategy(subtractStrategy)
	fmt.Println("Subtraction:", context.ExecuteStrategy(5, 3))
}
