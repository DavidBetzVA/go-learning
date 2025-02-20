package main

import "fmt"

type Product struct {
	PartA string
	PartB string
	PartC string
}

type Builder interface {
	SetPartA()
	SetPartB()
	SetPartC()
	GetResult() Product
}

type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) SetPartA() {
	b.product.PartA = "Part A set"
}

func (b *ConcreteBuilder) SetPartB() {
	b.product.PartB = "Part B set"
}

func (b *ConcreteBuilder) SetPartC() {
	b.product.PartC = "Part C set"
}

func (b *ConcreteBuilder) GetResult() Product {
	return b.product
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() Product {
	d.builder.SetPartA()
	d.builder.SetPartB()
	d.builder.SetPartC()
	return d.builder.GetResult()
}

func main() {
	builder := &ConcreteBuilder{}
	director := &Director{builder}
	product := director.Construct()
	fmt.Println(product)
}
