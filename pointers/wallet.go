package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int 

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// 복사본이 아니라 주소를 전달하여 원본 데이터를 변경
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	
	w.balance -= amount
	return nil
}