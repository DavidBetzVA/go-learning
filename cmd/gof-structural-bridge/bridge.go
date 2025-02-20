package main

import "fmt"

type Implementor interface {
	OperationImp() string
}

type ConcreteImplementorA struct{}

func (c *ConcreteImplementorA) OperationImp() string {
	return "ConcreteImplementorA: Operation Implementation."
}

type ConcreteImplementorB struct{}

func (c *ConcreteImplementorB) OperationImp() string {
	return "ConcreteImplementorB: Operation Implementation."
}

type Abstraction struct {
	impl Implementor
}

func (a *Abstraction) Operation() string {
	return a.impl.OperationImp()
}

func main() {
	implA := &ConcreteImplementorA{}
	abstractionA := &Abstraction{implA}
	fmt.Println(abstractionA.Operation())

	implB := &ConcreteImplementorB{}
	abstractionB := &Abstraction{implB}
	fmt.Println(abstractionB.Operation())
}
