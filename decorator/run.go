package decorator

import "fmt"

func Run() {
	component := &ConcreteComponent{}
	decoratorA := &ConcreteDecoratorA{}
	decoratorB := &ConcreteDecoratorB{}

	decoratorA.SetComponent(component)
	decoratorB.SetComponent(decoratorA)

	fmt.Println(decoratorB.Operation())
}
