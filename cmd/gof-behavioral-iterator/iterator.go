package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Collection interface {
	CreateIterator() Iterator
}

type ConcreteCollection struct {
	items []string
}

func (c *ConcreteCollection) CreateIterator() Iterator {
	return &ConcreteIterator{collection: c, index: 0}
}

type ConcreteIterator struct {
	collection *ConcreteCollection
	index      int
}

func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.collection.items)
}

func (i *ConcreteIterator) Next() interface{} {
	if i.HasNext() {
		item := i.collection.items[i.index]
		i.index++
		return item
	}
	return nil
}

func main() {
	collection := &ConcreteCollection{items: []string{"One", "Two", "Three"}}
	iterator := collection.CreateIterator()

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
