package builder

import "fmt"

func Run() {
	builder := &ConcreteBuilder{product: &Product{}}
	director := &Director{builder: builder}
	director.Construct()
	product := builder.GetResult()
	fmt.Printf("Product: partA=%s, partB=%s, partC=%s\n", product.partA, product.partB, product.partC)
}
