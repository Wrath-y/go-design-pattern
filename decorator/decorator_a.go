package decorator

type ConcreteDecoratorA struct {
	Decorator
}

func (c *ConcreteDecoratorA) Operation() string {
	return "ConcreteDecoratorA(" + c.Decorator.Operation() + ")"
}
