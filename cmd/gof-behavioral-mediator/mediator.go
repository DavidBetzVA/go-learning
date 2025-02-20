package main

import "fmt"

type Mediator interface {
	Notify(sender Component, event string)
}

type Component interface {
	SetMediator(mediator Mediator)
}

type ConcreteMediator struct {
	componentA *ComponentA
	componentB *ComponentB
}

func (m *ConcreteMediator) Notify(sender Component, event string) {
	if event == "A" {
		fmt.Println("Mediator reacts on A and triggers B.")
		m.componentB.DoB()
	}
}

type ComponentA struct {
	mediator Mediator
}

func (c *ComponentA) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ComponentA) DoA() {
	fmt.Println("ComponentA does A.")
	c.mediator.Notify(c, "A")
}

type ComponentB struct {
	mediator Mediator
}

func (c *ComponentB) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ComponentB) DoB() {
	fmt.Println("ComponentB does B.")
}

func main() {
	mediator := &ConcreteMediator{}
	componentA := &ComponentA{}
	componentB := &ComponentB{}

	componentA.SetMediator(mediator)
	componentB.SetMediator(mediator)

	mediator.componentA = componentA
	mediator.componentB = componentB

	componentA.DoA()
}
