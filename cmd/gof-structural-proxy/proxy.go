package main

import "fmt"

type Subject interface {
	Request() string
}

type RealSubject struct{}

func (r *RealSubject) Request() string {
	return "RealSubject: Handling request"
}

type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Request() string {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	return "Proxy: Delegating request -> " + p.realSubject.Request()
}

func main() {
	proxy := &Proxy{}
	fmt.Println(proxy.Request())
}
