package chain_of_responsibility

import "fmt"

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
