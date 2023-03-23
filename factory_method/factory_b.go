package factory_method

import "fmt"

type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Operation() {
	fmt.Println("ConcreteProductB operation")
}
