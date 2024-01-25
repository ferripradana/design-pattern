package complex_subsystem

import "fmt"

type Wallet struct {
	Balance float32
}

func NewWallet() *Wallet {
	return &Wallet{Balance: 0}
}

func (w *Wallet) CreditBalance(amount float32) {
	w.Balance += amount
	fmt.Println("Wallet Balance Added Successfully")
}

func (w *Wallet) DebitBalance(amount float32) error {
	if w.Balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet Balance is Sufficient")
	w.Balance -= amount
	return nil
}
