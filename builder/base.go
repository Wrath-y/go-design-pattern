package builder

type Product struct {
	partA string
	partB string
	partC string
}

func (p *Product) SetPartA(partA string) {
	p.partA = partA
}

func (p *Product) SetPartB(partB string) {
	p.partB = partB
}

func (p *Product) SetPartC(partC string) {
	p.partC = partC
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetResult() *Product
}
