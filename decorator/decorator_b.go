package decorator

type ConcreteDecoratorB struct {
	Decorator
}

func (c *ConcreteDecoratorB) Operation() string {
	return "ConcreteDecoratorB(" + c.Decorator.Operation() + ")"
}
