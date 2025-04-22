package main

import "fmt"

func main() {
	payment := &PaymentProccessorFactory{
		CardInfo: &CreditCard{
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
	loggingDec := &LoggingDecorator{
		PS: paymentMethod,
	}
	loggingDec.Pay(100.0)

	fraudDec := &FraudDetectionDecorator{
		PS: paymentMethod,
	}

	err = fraudDec.Pay(2000.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}
