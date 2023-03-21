package factory_simple

func Run() {
	factory := &Factory{}

	productA := factory.CreateProduct("A")
	productA.Operation()

	productB := factory.CreateProduct("B")
	productB.Operation()
}
