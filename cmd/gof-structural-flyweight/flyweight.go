package main

import "fmt"

type Flyweight interface {
	Operation(extrinsicState string) string
}

type ConcreteFlyweight struct {
	intrinsicState string
}

func (f *ConcreteFlyweight) Operation(extrinsicState string) string {
	return "ConcreteFlyweight: Intrinsic[" + f.intrinsicState + "] Extrinsic[" + extrinsicState + "]"
}

type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{flyweights: make(map[string]Flyweight)}
}

func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if _, exists := f.flyweights[key]; !exists {
		f.flyweights[key] = &ConcreteFlyweight{intrinsicState: key}
	}
	return f.flyweights[key]
}

func main() {
	factory := NewFlyweightFactory()

	fw1 := factory.GetFlyweight("A")
	fmt.Println(fw1.Operation("X"))

	fw2 := factory.GetFlyweight("A")
	fmt.Println(fw2.Operation("Y"))
}
