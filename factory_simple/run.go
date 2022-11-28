package factory_simple

func Run() {
	println(NewPrinter("cn").Print("你好"))
	println(NewPrinter("cn").Print("Hello"))
}
