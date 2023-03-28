package chain_of_responsibility

import "fmt"

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
