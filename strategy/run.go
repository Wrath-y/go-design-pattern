package strategy

import "fmt"

// Run
// 在示例中，我们定义了一个 PaymentMethod 接口和两个结构体 CreditCard 和 PayPal 分别实现了该接口。
// 然后我们定义了一个 PaymentProcessor 结构体，并在其中定义了一个 ProcessPayment() 方法，该方法接受一个金额和支付方式作为参数，并返回支付结果。
// 最后，在 Run 函数中我们创建了两个不同的支付方式并调用 ProcessPayment() 方法进行支付。
func Run() {
	paymentProcessor := &PaymentProcessor{}

	creditCard := &CreditCard{}
	result := paymentProcessor.ProcessPayment(100, creditCard)
	fmt.Println(result)

	payPal := &PayPal{}
	result = paymentProcessor.ProcessPayment(200, payPal)
	fmt.Println(result)
}
