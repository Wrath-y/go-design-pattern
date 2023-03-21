package factory_simple

import "fmt"

type Product interface {
	Operation()
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) Operation() {
	fmt.Println("ConcreteProductA operation")
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Operation() {
	fmt.Println("ConcreteProductB operation")
}

type Factory struct{}

func (f *Factory) CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}
