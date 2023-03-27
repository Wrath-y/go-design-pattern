package factory_abstract

import "fmt"

func Run() {
	factory1 := &ConcreteFactory1{}
	productA1 := factory1.CreateProductA()
	productB1 := factory1.CreateProductB()
	fmt.Println(productA1.OperationA())
	fmt.Println(productB1.OperationB())

	factory2 := &ConcreteFactory2{}
	productA2 := factory2.CreateProductA()
	productB2 := factory2.CreateProductB()
	fmt.Println(productA2.OperationA())
	fmt.Println(productB2.OperationB())
}
