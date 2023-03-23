package factory_method

func Run() {
	factoryA := &ConcreteFactoryA{}
	factoryA.CreateProduct().Operation()

	factoryB := &ConcreteFactoryB{}
	factoryB.CreateProduct().Operation()
}
