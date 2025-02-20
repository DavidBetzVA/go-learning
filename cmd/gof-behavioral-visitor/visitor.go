package main

import "fmt"

type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA)
	VisitConcreteElementB(*ConcreteElementB)
}

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Println("Visiting ConcreteElementA")
}

func (v *ConcreteVisitor) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Println("Visiting ConcreteElementB")
}

func main() {
	elements := []Element{&ConcreteElementA{}, &ConcreteElementB{}}
	visitor := &ConcreteVisitor{}

	for _, elem := range elements {
		elem.Accept(visitor)
	}
}
