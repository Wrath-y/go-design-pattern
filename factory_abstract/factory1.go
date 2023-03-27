package factory_abstract

type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProductA() ProductA {
	return &ConcreteProductA1{}
}

func (f *ConcreteFactory1) CreateProductB() ProductB {
	return &ConcreteProductB1{}
}

type ConcreteProductA1 struct{}

func (p *ConcreteProductA1) OperationA() string {
	return "ConcreteProductA1"
}

type ConcreteProductB1 struct{}

func (p *ConcreteProductB1) OperationB() string {
	return "ConcreteProductB1"
}
