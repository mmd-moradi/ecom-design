package strategies

import "fmt"

type PayPal struct {
	Email    string
	Password string
}

func (pm *PayPal) Pay(amount float64) error {
	fmt.Println("Paying with PayPal")
	return nil
}
