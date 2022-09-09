package pointer

import (
	"errors"
	"fmt"
)

var ErrInSufficientBalance = errors.New("can not withdraw, insufficient balance")

type Bitcoin int

// When you print this values with %s format, this method will be called automatically
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// Deposit In Go, when you call a function or a method the arguments are copied.
// When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from

//func (w Wallet) Deposit(amount int) {
//	fmt.Printf("Address of w balance in Deposit: %v\n", &w.balance)
//
//	w.balance += amount
//}
//
//func (w Wallet) Balance() int {
//	fmt.Printf("Address of w balance in Balance: %v\n", &w.balance)
//
//	return w.balance
//}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInSufficientBalance
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
