package main

import "fmt"

type Command interface {
	Execute() string
}

type Receiver struct{}

func (r *Receiver) Action() string {
	return "Receiver: Performing action."
}

type ConcreteCommand struct {
	receiver *Receiver
}

func (c *ConcreteCommand) Execute() string {
	return c.receiver.Action()
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) Invoke() string {
	return i.command.Execute()
}

func main() {
	receiver := &Receiver{}
	command := &ConcreteCommand{receiver}
	invoker := &Invoker{}
	invoker.SetCommand(command)

	fmt.Println(invoker.Invoke())
}
