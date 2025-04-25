package strategies

import "fmt"

type CreditCard struct {
	CardNumber  string
	ExpirayDate string
	CVV         string
	CardHolder  string
}

func (pm *CreditCard) Pay(amount float64) error {
	fmt.Printf("Paying with credit card %.2f\n", amount)
	return nil
}
