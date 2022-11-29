package factory_method

type OperatorFactory interface {
	Create() MathOperator
}

type MathOperator interface {
	SetA(int)
	SetB(int)
	GetResult() int
}

type BaseOperator struct {
	a, b int
}

func (bo *BaseOperator) SetA(a int) {
	bo.a = a
}

func (bo *BaseOperator) SetB(b int) {
	bo.b = b
}
