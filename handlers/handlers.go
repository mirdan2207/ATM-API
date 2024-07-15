package handlers

import (
	"ATM-API/logger"
	"ATM-API/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var (
	accounts = make(map[int]*models.Account)
	nextID   = 1
	mu       sync.Mutex
)

// CreateAccountHandler создает новый аккаунт
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	acc := models.NewAccount(nextID)
	accounts[nextID] = acc
	nextID++
	mu.Unlock()

	logger.LogOperation("Account created", acc.ID)
	w.WriteHeader(http.StatusCreated)
}

// DepositHandler пополнение баланса
func DepositHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	acc, exists := accounts[id]
	mu.Unlock()
	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	var request struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resultChan := make(chan error)
	go func() {
		resultChan <- acc.Deposit(request.Amount)
	}()

	if err := <-resultChan; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		logger.LogDeposit(acc.ID, request.Amount)
		w.WriteHeader(http.StatusOK)
	}
}

// WithdrawHandler снятие средств
func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	acc, exists := accounts[id]
	mu.Unlock()
	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	var request struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resultChan := make(chan error)
	go func() {
		resultChan <- acc.Withdraw(request.Amount)
	}()

	if err := <-resultChan; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		logger.LogWithdraw(acc.ID, request.Amount)
		w.WriteHeader(http.StatusOK)
	}
}

// BalanceHandler проверка баланса
func BalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	acc, exists := accounts[id]
	mu.Unlock()
	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	resultChan := make(chan float64)
	go func() {
		resultChan <- acc.GetBalance()
	}()

	balance := <-resultChan
	logger.LogBalance(acc.ID, balance)
	fmt.Fprintf(w, "Balance: %.2f", balance)
}
