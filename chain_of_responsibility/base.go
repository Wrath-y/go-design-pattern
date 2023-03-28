package chain_of_responsibility

import "fmt"

type Handler interface {
	SetNextHandler(handler Handler)
	HandleRequest(request int)
}

type ConcreteHandlerA struct {
	nextHandler Handler
}

func (h *ConcreteHandlerA) SetNextHandler(handler Handler) {
	h.nextHandler = handler
}

func (h *ConcreteHandlerA) HandleRequest(request int) {
	if request >= 0 && request < 10 {
		fmt.Printf("%s handled request %d\n", "ConcreteHandlerA", request)
	} else if h.nextHandler != nil {
		h.nextHandler.HandleRequest(request)
	}
}

type ConcreteHandlerB struct {
	nextHandler Handler
}

func (h *ConcreteHandlerB) SetNextHandler(handler Handler) {
	h.nextHandler = handler
}

func (h *ConcreteHandlerB) HandleRequest(request int) {
	if request >= 10 && request < 20 {
		fmt.Printf("%s handled request %d\n", "ConcreteHandlerB", request)
	} else if h.nextHandler != nil {
		h.nextHandler.HandleRequest(request)
	}
}

type ConcreteHandlerC struct {
	nextHandler Handler
}

func (h *ConcreteHandlerC) SetNextHandler(handler Handler) {
	h.nextHandler = handler
}

func (h *ConcreteHandlerC) HandleRequest(request int) {
	if request >= 20 && request < 30 {
		fmt.Printf("%s handled request %d\n", "ConcreteHandlerC", request)
	} else if h.nextHandler != nil {
		h.nextHandler.HandleRequest(request)
	}
}
