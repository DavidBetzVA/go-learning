package main

import "fmt"

type AbstractClass interface {
	Step1()
	Step2()
	TemplateMethod()
}

type ConcreteClass struct{}

func (c *ConcreteClass) Step1() {
	fmt.Println("Step 1 executed")
}

func (c *ConcreteClass) Step2() {
	fmt.Println("Step 2 executed")
}

func (c *ConcreteClass) TemplateMethod() {
	c.Step1()
	c.Step2()
}

func main() {
	concrete := &ConcreteClass{}
	concrete.TemplateMethod()
}
