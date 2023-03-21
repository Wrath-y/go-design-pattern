package strategy

type PaymentMethod interface {
	Pay(amount float64) string
}

type PaymentProcessor struct{}

func (pp *PaymentProcessor) ProcessPayment(amount float64, paymentMethod PaymentMethod) string {
	return paymentMethod.Pay(amount)
}
