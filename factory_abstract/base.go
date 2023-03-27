package factory_abstract

type AbstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

type ProductA interface {
	OperationA() string
}

type ProductB interface {
	OperationB() string
}
