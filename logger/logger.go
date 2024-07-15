package logger

import (
	"log"
	"time"
)

func LogOperation(operation string, accountID int) {
	log.Printf("[%s] Operation: %s, Account ID: %d\n", time.Now().Format(time.RFC3339), operation, accountID)
}

func LogDeposit(accountID int, amount float64) {
	log.Printf("[%s] Operation: Deposit, Account ID: %d, Amount: %.2f\n", time.Now().Format(time.RFC3339), accountID, amount)
}

func LogWithdraw(accountID int, amount float64) {
	log.Printf("[%s] Operation: Withdraw, Account ID: %d, Amount: %.2f\n", time.Now().Format(time.RFC3339), accountID, amount)
}

func LogBalance(accountID int, balance float64) {
	log.Printf("[%s] Operation: Balance checked, Account ID: %d, Balance: %.2f\n", time.Now().Format(time.RFC3339), accountID, balance)
}
