package main

import (
	"fmt"

	"github.com/mmd-moradi/ecom-design/internal/decorators"
	"github.com/mmd-moradi/ecom-design/internal/payment"
	"github.com/mmd-moradi/ecom-design/internal/strategies"
)

func main() {
	payment := &payment.PaymentProccessorFactory{
		CardInfo: &strategies.CreditCard{
			CardNumber:  "1234-5678-9012-3456",
			ExpirayDate: "12/25",
			CVV:         "123",
			CardHolder:  "John Doe",
		},
	}
	paymentMethod, err := payment.NewPayment("creditcard")
	if err != nil {
		panic(err)
	}
	loggingDec := &decorators.LoggingDecorator{
		PS: paymentMethod,
	}
	loggingDec.Pay(100.0)

	fraudDec := &decorators.FraudDetectionDecorator{
		PS: paymentMethod,
	}

	err = fraudDec.Pay(2000.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}
