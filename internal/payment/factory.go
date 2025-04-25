package payment

import (
	"fmt"

	"github.com/mmd-moradi/ecom-design/internal/strategies"
)

type PaymentFactory interface {
	NewPayment(method string) PaymentStrategy
}

type PaymentProccessorFactory struct {
	CardInfo   *strategies.CreditCard
	PayPalInfo *strategies.PayPal
	CryptoInfo *strategies.Crypto
}

func (p *PaymentProccessorFactory) NewPayment(method string) (PaymentStrategy, error) {
	switch method {
	case "creditcard":
		if p.CardInfo == nil {
			return nil, fmt.Errorf("credit card info is not provided")
		}
		return &strategies.CreditCard{
			CardNumber:  p.CardInfo.CardNumber,
			ExpirayDate: p.CardInfo.ExpirayDate,
			CVV:         p.CardInfo.CVV,
			CardHolder:  p.CardInfo.CardHolder,
		}, nil
	case "paypal":
		if p.PayPalInfo == nil {
			return nil, fmt.Errorf("PayPal info is not provided")
		}
		return &strategies.PayPal{
			Email:    p.PayPalInfo.Email,
			Password: p.PayPalInfo.Password,
		}, nil
	case "crypto":
		if p.CryptoInfo == nil {
			fmt.Println("Crypto info is not provided")
			return nil, fmt.Errorf("crypto info is not provided")
		}
		return &strategies.Crypto{
			WalletAddress: p.CryptoInfo.WalletAddress,
			Currency:      p.CryptoInfo.Currency,
		}, nil
	}
	return nil, fmt.Errorf("unsupported payment method: %s", method)
}
