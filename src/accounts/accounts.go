package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type account struct {
	owner   string
	balance int
}

var ErrInsufficientBalance = errors.New("insufficient balance")

// NewAccount creates a new account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount to the account
func (x *account) Deposit(amount int) {
	x.balance += amount
}

// Balance returns the balance of the account
func (x *account) Balance() int {
	return x.balance
}

// Withdraw x amount from the account
func (x *account) Withdraw(amount int) error {
	if x.balance < amount {
		return ErrInsufficientBalance
	}
	x.balance -= amount
	return nil
}

// ChangeOwner changes the owner of the account
func (x *account) ChangeOwner(newOwner string) {
	x.owner = newOwner
}

// Owner returns the owner of the account
func (x *account) Owner() string {
	return x.owner
}

func (x *account) String() string {
	return fmt.Sprint(x.owner, "'s account has ", x.balance, " dollars")
}