package main

import "fmt"

type LoggingDecorator struct {
	PS IPaymentStrategy
}

type FraudDetectionDecorator struct {
	PS IPaymentStrategy
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

func (f *FraudDetectionDecorator) Pay(amount float64) error {
	fmt.Println("Checking for fraud...")
	if amount > 1000 {
		return fmt.Errorf("fraud detected for amount %.2f", amount)
	}
	fmt.Println("No fraud detected")
	return f.PS.Pay(amount)
}
