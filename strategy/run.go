package strategy

import "fmt"

func Run() {
	paymentProcessor := &PaymentProcessor{}

	creditCard := &CreditCard{}
	result := paymentProcessor.Pay(100, creditCard)
	fmt.Println(result)

	payPal := &PayPal{}
	result = paymentProcessor.Pay(200, payPal)
	fmt.Println(result)
}
