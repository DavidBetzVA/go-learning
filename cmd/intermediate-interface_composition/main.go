package main

import "fmt"

// Two interfaces
type Animal interface {
	Speak() string
}

type Walker interface {
	Walk() string
}

// Composed interface
type Dog interface {
	Animal
	Walker
}

type MyDog struct{}

func (MyDog) Speak() string { return "Woof!" }
func (MyDog) Walk() string  { return "Walking..." }

func main() {
	var d Dog = MyDog{}
	fmt.Println(d.Speak(), d.Walk())
}
