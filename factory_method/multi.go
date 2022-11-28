package factory_method

type MultiOperatorFactory struct{}

type MultiOperator struct {
	*BaseOperator
}

func (*MultiOperatorFactory) Create() MathOperator {
	return &MultiOperator{
		BaseOperator: &BaseOperator{},
	}
}

func (ao *MultiOperator) GetResult() int {
	return ao.a * ao.b
}
