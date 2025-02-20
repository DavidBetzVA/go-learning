package main

import "fmt"

type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string) string
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) HandleRequest(request string) string {
	if h.next != nil {
		return h.next.HandleRequest(request)
	}
	return "No handler processed the request."
}

type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) HandleRequest(request string) string {
	if request == "A" {
		return "Handler A processed the request."
	}
	return h.BaseHandler.HandleRequest(request)
}

type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) HandleRequest(request string) string {
	if request == "B" {
		return "Handler B processed the request."
	}
	return h.BaseHandler.HandleRequest(request)
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerA.SetNext(handlerB)

	fmt.Println(handlerA.HandleRequest("A"))
	fmt.Println(handlerA.HandleRequest("B"))
	fmt.Println(handlerA.HandleRequest("C"))
}
