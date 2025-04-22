package main

type IPaymentStrategy interface {
	Pay(amount float64) error
}
