package main

import "fmt"

type Target interface {
	Request() string
}

type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee: Specific request."
}

type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee}
	fmt.Println(adapter.Request())
}
