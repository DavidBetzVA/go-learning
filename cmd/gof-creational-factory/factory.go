package main

import "fmt"

type Product interface {
	Use() string
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string { return "Using Product A" }

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string { return "Using Product B" }

func CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}

func main() {
	p1 := CreateProduct("A")
	fmt.Println(p1.Use())

	p2 := CreateProduct("B")
	fmt.Println(p2.Use())
}
