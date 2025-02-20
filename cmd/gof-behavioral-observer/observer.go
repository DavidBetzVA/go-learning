package main

import "fmt"

type Observer interface {
	Update(string)
}

type Subject struct {
	observers []Observer
	state     string
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) SetState(state string) {
	s.state = state
	s.Notify()
}

func (s *Subject) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

type ConcreteObserver struct {
	id int
}

func (o *ConcreteObserver) Update(state string) {
	fmt.Printf("Observer %d received new state: %s\n", o.id, state)
}

func main() {
	subject := &Subject{}
	observer1 := &ConcreteObserver{id: 1}
	observer2 := &ConcreteObserver{id: 2}

	subject.Attach(observer1)
	subject.Attach(observer2)

	subject.SetState("New State")
}
