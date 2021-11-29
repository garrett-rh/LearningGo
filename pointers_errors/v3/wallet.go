package v3_main

import "fmt"

type Bitcoin int //This creates a new type which we can declare methods on

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

//Creating the stringer interface for bitcoin object so that it prints NUM BTC instead of just NUM
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdrawal(amount Bitcoin) {
	w.balance -= amount
}
