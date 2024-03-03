package routes

import (
	"encoding/json"
	"net/http"
	db "payment_service/utils/database"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type TransactionDetails struct {
	ID          int       `json:"id"`
	AccountFrom int       `json:"account_from"`
	AccountTo   int       `json:"account_to"`
	Amount      string    `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
}

type TransferRequest struct {
	AccountFrom int     `json:"account_from"`
	AccountTo   int     `json:"account_to"`
	Amount      float64 `json:"amount"`
}

func GetTransactionDetails(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	unparsedTransactionId := vars["id"]

	transactionId, err := strconv.Atoi(unparsedTransactionId)
	if err != nil {
		http.Error(res, "Transaction ID must be a number", http.StatusBadRequest)
		return
	}

	db := db.Open()
	defer db.Close()

	var transaction TransactionDetails

	err = db.QueryRow("SELECT * FROM transactions WHERE id = $1", transactionId).Scan(
		&transaction.ID,
		&transaction.AccountFrom,
		&transaction.AccountTo,
		&transaction.Amount,
		&transaction.CreatedAt,
		&transaction.Description)

	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(transaction)
}

func GetAccountTransactions(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	unparsedAccountId := vars["id"]

	accountId, err := strconv.Atoi(unparsedAccountId)
	if err != nil {
		http.Error(res, "Account ID must be a number", http.StatusBadRequest)
		return
	}

	db := db.Open()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM transactions WHERE account_from = $1 OR account_to = $1", accountId)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var transactions []TransactionDetails

	for rows.Next() {
		var transaction TransactionDetails
		err = rows.Scan(
			&transaction.ID,
			&transaction.AccountFrom,
			&transaction.AccountTo,
			&transaction.Amount,
			&transaction.CreatedAt,
			&transaction.Description)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(transactions)
}

func TransferMoney(res http.ResponseWriter, req *http.Request) {
	var transferDetails TransferRequest
	err := json.NewDecoder(req.Body).Decode(&transferDetails)
	if err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	if transferDetails.AccountFrom == transferDetails.AccountTo {
		http.Error(res, "Sender and receiver accounts must be different", http.StatusBadRequest)
		return
	}

	if transferDetails.Amount <= 0 {
		http.Error(res, "Transfer amount must be positive", http.StatusBadRequest)
		return
	}

	db := db.Open()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		http.Error(res, "Could not start database transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	var balance string
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = $1", transferDetails.AccountFrom).Scan(&balance)

	if err != nil {
		http.Error(res, "Sender account not found", http.StatusBadRequest)
		return
	}
	balanceFloat, err := strconv.ParseFloat(strings.TrimPrefix(balance, "$"), 64)
	if err != nil {
		http.Error(res, "Could not convert balance to float", http.StatusInternalServerError)
		return
	}

	if balanceFloat < transferDetails.Amount {
		http.Error(res, "Insufficient funds", http.StatusBadRequest)
		return
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", transferDetails.Amount, transferDetails.AccountFrom)
	if err != nil {
		http.Error(res, "Could not update sender account balance", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", transferDetails.Amount, transferDetails.AccountTo)
	if err != nil {
		http.Error(res, "Could not update receiver account balance", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("INSERT INTO transactions (account_from, account_to, amount, description, created_at) VALUES ($1, $2, $3, $4, NOW())",
		transferDetails.AccountFrom, transferDetails.AccountTo, transferDetails.Amount, "Transfer")
	if err != nil {
		http.Error(res, "Could not insert transaction record", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(res, "Could not commit transaction", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]string{"status": "success"})
}
