package decorators

import (
	"fmt"

	"github.com/mmd-moradi/ecom-design/internal/payment"
)

type FraudDetectionDecorator struct {
	PS payment.PaymentStrategy
}

func (f *FraudDetectionDecorator) Pay(amount float64) error {
	fmt.Println("Checking for fraud...")
	if amount > 1000 {
		return fmt.Errorf("fraud detected for amount %.2f", amount)
	}
	fmt.Println("No fraud detected")
	return f.PS.Pay(amount)
}
