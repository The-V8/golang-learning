package pointersanderrors

import (
	"errors"
	"fmt"
)

// // Stringer lorem ipsum
// type Stringer interface {
// 	String() string
// }

// Bitcoin makes people crazy
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// ErrInsufficientFunds cannot withdraw, insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Wallet keeps some bitcoins for me
type Wallet struct {
	balance Bitcoin
}

// Withdraw an amount of Bitcoins
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Deposit adds an amount of bitcoins to the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns current balance of wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
