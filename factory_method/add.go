package factory_method

type AddOperatorFactory struct{}

type AddOperator struct {
	*BaseOperator
}

func (*AddOperatorFactory) Create() MathOperator {
	return &AddOperator{
		BaseOperator: &BaseOperator{},
	}
}

func (ao *AddOperator) GetResult() int {
	return ao.a + ao.b
}
