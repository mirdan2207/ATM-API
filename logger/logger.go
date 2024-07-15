package logger

import (
	"log"
	"time"
)

// LogOperation логирует операцию с указанием времени, типа операции и идентификатора аккаунта.
func LogOperation(operation string, accountID int) {
	log.Printf("[%s] Operation: %s, Account ID: %d\n", time.Now().Format(time.RFC3339), operation, accountID)
}

// LogDeposit логирует операцию пополнения баланса.
func LogDeposit(accountID int, amount float64) {
	log.Printf("[%s] Operation: Deposit, Account ID: %d, Amount: %.2f\n", time.Now().Format(time.RFC3339), accountID, amount)
}

// LogWithdraw логирует операцию снятия средств.
func LogWithdraw(accountID int, amount float64) {
	log.Printf("[%s] Operation: Withdraw, Account ID: %d, Amount: %.2f\n", time.Now().Format(time.RFC3339), accountID, amount)
}

// LogBalance логирует операцию проверки баланса.
func LogBalance(accountID int, balance float64) {
	log.Printf("[%s] Operation: Balance checked, Account ID: %d, Balance: %.2f\n", time.Now().Format(time.RFC3339), accountID, balance)
}
