// The Facade pattern provides a simplified interface to a complex subsystem.
// This example demonstrates the Facade pattern by creating a Facade struct
// that simplifies the interaction with SubsystemA and SubsystemB.

package main

import "fmt"

// SubsystemA is a complex subsystem that has its own operations.
type SubsystemA struct{}

func (s *SubsystemA) OperationA() string {
	return "SubsystemA: Operation A"
}

// SubsystemB is another complex subsystem that has its own operations.
type SubsystemB struct{}

func (s *SubsystemB) OperationB() string {
	return "SubsystemB: Operation B"
}

// Facade provides a simplified interface to the complex subsystems.
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

// NewFacade creates a new Facade with initialized subsystems.
func NewFacade() *Facade {
	return &Facade{&SubsystemA{}, &SubsystemB{}}
}

// Operation provides a simplified interface to the operations of the subsystems.
func (f *Facade) Operation() string {
	return f.subsystemA.OperationA() + " | " + f.subsystemB.OperationB()
}

func main() {
	facade := NewFacade()
	fmt.Println(facade.Operation())
}
