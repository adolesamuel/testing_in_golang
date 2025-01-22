package wallet

import "fmt"

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (w *Wallet) Deposit(balance Bitcoin) {
	w.balance += balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
