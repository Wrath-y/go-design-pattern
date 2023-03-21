package strategy

import "fmt"

type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paying %.2f using credit card", amount)
}
