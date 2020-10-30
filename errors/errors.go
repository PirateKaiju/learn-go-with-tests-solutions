package errors

import "errors"

type Wallet struct {
	balance int
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) deposit(value int) {
	w.balance += value
}

var ErrInsufficientFunds = errors.New("cannot withdraw, isufficient funds")

func (w *Wallet) withdraw(value int) error {
	if value > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= value

	return nil
}
