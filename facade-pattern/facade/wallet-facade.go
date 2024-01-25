package facade

import (
	complex_subsystem "facade-pattern/complex-subsystem"
	"fmt"
)

type WalletFacade struct {
	Account      *complex_subsystem.Account
	Wallet       *complex_subsystem.Wallet
	SecurityCode *complex_subsystem.SecurityCode
	Ledger       *complex_subsystem.Ledger
	Notification *complex_subsystem.Notification
}

func NewWalletFacade(accountId string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacade := &WalletFacade{
		Account:      complex_subsystem.NewAccount(accountId),
		Wallet:       complex_subsystem.NewWallet(),
		SecurityCode: complex_subsystem.NewSecurityCode(code),
		Ledger:       &complex_subsystem.Ledger{},
		Notification: &complex_subsystem.Notification{},
	}
	fmt.Println("Account created")
	return walletFacade
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount float32) error {
	fmt.Println("Starting add money to wallet")
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	w.Wallet.CreditBalance(amount)
	w.Notification.SendWalletCreditNotification()
	w.Ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount float32) error {
	fmt.Println("Starting debit money from wallet")
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = w.Wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.Notification.SendWalletDebitNotification()
	w.Ledger.MakeEntry(accountID, "debit", amount)
	return nil
}
