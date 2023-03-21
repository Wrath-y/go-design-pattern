package strategy

import "fmt"

type PayPal struct{}

func (p *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paying %.2f using PayPal", amount)
}
