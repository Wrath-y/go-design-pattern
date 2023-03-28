package builder

type ConcreteBuilder struct {
	product *Product
}

func (b *ConcreteBuilder) BuildPartA() {
	b.product.SetPartA("PartA")
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.SetPartB("PartB")
}

func (b *ConcreteBuilder) BuildPartC() {
	b.product.SetPartC("PartC")
}

func (b *ConcreteBuilder) GetResult() *Product {
	return b.product
}
