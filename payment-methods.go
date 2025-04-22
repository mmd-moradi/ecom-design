package main

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

type PayPal struct {
	Email    string
	Password string
}

func (pm *PayPal) Pay(amount float64) error {
	fmt.Println("Paying with PayPal")
	return nil
}

type Crypto struct {
	WalletAddress string
	Currency      string
}

func (pm *Crypto) Pay(amount float64) error {
	fmt.Println("Paying with crypto")
	return nil
}
