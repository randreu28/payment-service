package routes

import (
	"encoding/json"
	"net/http"
	db "payment_service/utils/database"
	"strconv"
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
