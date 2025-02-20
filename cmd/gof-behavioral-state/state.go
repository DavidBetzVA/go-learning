package main

import "fmt"

type State interface {
	Handle() string
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() string {
	return "State A: Handling request"
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() string {
	return "State B: Handling request"
}

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() string {
	return c.state.Handle()
}

func main() {
	context := &Context{}

	stateA := &ConcreteStateA{}
	stateB := &ConcreteStateB{}

	context.SetState(stateA)
	fmt.Println(context.Request())

	context.SetState(stateB)
	fmt.Println(context.Request())
}
