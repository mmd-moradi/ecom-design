package strategies

import "fmt"

type Crypto struct {
	WalletAddress string
	Currency      string
}

func (pm *Crypto) Pay(amount float64) error {
	fmt.Println("Paying with crypto")
	return nil
}
