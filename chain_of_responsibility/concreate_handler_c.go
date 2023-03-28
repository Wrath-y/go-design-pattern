package chain_of_responsibility

import "fmt"

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
