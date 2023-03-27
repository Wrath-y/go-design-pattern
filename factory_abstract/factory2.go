package factory_abstract

type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProductA() ProductA {
	return &ConcreteProductA2{}
}

func (f *ConcreteFactory2) CreateProductB() ProductB {
	return &ConcreteProductB2{}
}

type ConcreteProductA2 struct{}

func (p *ConcreteProductA2) OperationA() string {
	return "ConcreteProductA2"
}

type ConcreteProductB2 struct{}

func (p *ConcreteProductB2) OperationB() string {
	return "ConcreteProductB2"
}
