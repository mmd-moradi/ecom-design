package decorators

import (
	"fmt"

	"github.com/mmd-moradi/ecom-design/internal/payment"
)

type LoggingDecorator struct {
	PS payment.PaymentStrategy
}

func (l *LoggingDecorator) Pay(amount float64) error {
	fmt.Println("Start payment process of amount:", amount)
	err := l.PS.Pay(amount)
	if err != nil {
		fmt.Println("Payment failed:", err)
		return err
	}
	fmt.Println("Payment Successful")
	return nil
}
