package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	db "payment_service/utils/database"
	env "payment_service/utils/env"
	jwt "payment_service/utils/jwt"
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
	r.HandleFunc("/account/transactions", GetAccountTransactions).Methods("GET")
	r.Use(jwt.AuthMiddleware)

	req, err := http.NewRequest("GET", "/account/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	token, err := jwt.CreateJwt(1, "Owner1")
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", token)
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
	r.Use(jwt.AuthMiddleware)

	transferDetails := TransferRequest{
		AccountTo: 2,
		Amount:    10,
	}

	transferBody, err := json.Marshal(transferDetails)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/transfer", bytes.NewBuffer(transferBody))
	if err != nil {
		t.Fatal(err)
	}

	token, err := jwt.CreateJwt(1, "Owner1")
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", token)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
