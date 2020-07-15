package pointerserrors

// Wallet definition
type Wallet struct {
	balance int
}

// Deposit from wallet
func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance of a wallet
func (w Wallet) Balance() int {
	return w.balance
}
