package internal

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/facade/internal/complexsubsystem"
)

type WalletFacade struct {
	account      *complexsubsystem.Account
	wallet       *complexsubsystem.Wallet
	securityCode *complexsubsystem.SecurityCode
	notification *complexsubsystem.Notification
	ledger       *complexsubsystem.Ledger
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &WalletFacade{
		account:      complexsubsystem.NewAccount(accountID),
		securityCode: complexsubsystem.NewSecurityCode(code),
		wallet:       complexsubsystem.NewWallet(),
		notification: &complexsubsystem.Notification{},
		ledger:       &complexsubsystem.Ledger{},
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.CreditBalance(amount)
	w.notification.SendWalletCreditNotification()
	w.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.SendWalletDebitNotification()
	w.ledger.MakeEntry(accountID, "debit", amount)
	return nil
}
