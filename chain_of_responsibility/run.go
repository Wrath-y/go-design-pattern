package chain_of_responsibility

func Run() {
	h1 := &ConcreteHandlerA{}
	h2 := &ConcreteHandlerB{}
	h3 := &ConcreteHandlerC{}

	h1.SetNextHandler(h2)
	h2.SetNextHandler(h3)

	requests := []int{2, 5, 14, 22, 18, 3, 27, 20}

	for _, request := range requests {
		h1.HandleRequest(request)
	}
}
