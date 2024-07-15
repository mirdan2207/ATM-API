package models

import (
	"fmt"
	"sync"
)

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	ID      int
	balance float64
	mu      sync.Mutex
}

func NewAccount(id int) *Account {
	return &Account{ID: id, balance: 0}
}

func (a *Account) Deposit(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount <= 0 {
		return fmt.Errorf("invalid amount")
	}
	a.balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount <= 0 {
		return fmt.Errorf("invalid amount")
	}
	if a.balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	a.balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}
