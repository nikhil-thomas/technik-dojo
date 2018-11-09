package bank

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds means wallet doesnot have enough funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin represents the number of bit coins
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stroers the number of bitcoins someone owns
type Wallet struct {
	balance Bitcoin
}

// Deposit will add some bitcoin to current balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw will withdraw some bitcoin to current balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance returns current balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
