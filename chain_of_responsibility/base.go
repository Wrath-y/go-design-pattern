package chain_of_responsibility

type Handler interface {
	SetNextHandler(handler Handler)
	HandleRequest(request int)
}
