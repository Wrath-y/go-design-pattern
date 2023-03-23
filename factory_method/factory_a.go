package factory_method

import "fmt"

type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) Operation() {
	fmt.Println("ConcreteProductA operation")
}
