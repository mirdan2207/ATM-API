package main

import (
	"ATM-API/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/accounts", handlers.CreateAccountHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/deposit", handlers.DepositHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/withdraw", handlers.WithdrawHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}/balance", handlers.BalanceHandler).Methods("GET")

	address := ":8080"
	log.Printf("Server starting on http://localhost%s\n", address)
	log.Fatal(http.ListenAndServe(address, r))
}
