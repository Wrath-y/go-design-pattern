package factory_method

func Run() {
	var factory OperatorFactory
	var mathOp MathOperator

	factory = new(AddOperatorFactory)
	mathOp = factory.Create()
	mathOp.SetA(1)
	mathOp.SetB(2)
	println(mathOp.GetResult())

	factory = new(MultiOperatorFactory)
	mathOp = factory.Create()
	mathOp.SetA(1)
	mathOp.SetB(2)
	println(mathOp.GetResult())
}
