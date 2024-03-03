package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	db "payment_service/utils/database"
	env "payment_service/utils/env"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetTransactionDetails(t *testing.T) {
	env.Load()

	database := db.Open()
	db.CleanupDB(database)
	db.PopulateDB(database, 100)
	defer database.Close()

	r := mux.NewRouter()
	r.HandleFunc("/transactions/{id}", GetTransactionDetails).Methods("GET")

	req, err := http.NewRequest("GET", "/transactions/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestGetAccountTransactions(t *testing.T) {
	env.Load()

	database := db.Open()
	db.CleanupDB(database)
	db.PopulateDB(database, 100)
	defer database.Close()

	r := mux.NewRouter()
	r.HandleFunc("/accounts/{id}/transactions", GetAccountTransactions).Methods("GET")

	req, err := http.NewRequest("GET", "/accounts/1/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var transactions []TransactionDetails
	err = json.NewDecoder(res.Body).Decode(&transactions)
	if err != nil {
		t.Fatal("Failed to decode response body")
	}

	if len(transactions) == 0 {
		t.Error("Expected at least one transaction, got 0")
	}
}

func TestTransferMoney(t *testing.T) {
	env.Load()

	database := db.Open()
	db.CleanupDB(database)
	db.PopulateDB(database, 100)
	defer database.Close()

	r := mux.NewRouter()
	r.HandleFunc("/transfer", TransferMoney).Methods("POST")

	transferDetails := TransferRequest{
		AccountFrom: 1,
		AccountTo:   2,
		Amount:      10,
	}

	transferBody, err := json.Marshal(transferDetails)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/transfer", bytes.NewBuffer(transferBody))
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
