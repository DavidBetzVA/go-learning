package main

import "fmt"

type Memento struct {
	state string
}

func (m *Memento) GetState() string {
	return m.state
}

type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) SaveState() *Memento {
	return &Memento{state: o.state}
}

func (o *Originator) RestoreState(m *Memento) {
	o.state = m.GetState()
}

func main() {
	originator := &Originator{}
	originator.SetState("State1")
	memento := originator.SaveState()

	originator.SetState("State2")
	fmt.Println("Current State:", originator.state)

	originator.RestoreState(memento)
	fmt.Println("Restored State:", originator.state)
}
