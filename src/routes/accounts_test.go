package routes

import (
	"net/http"
	"net/http/httptest"
	db "payment_service/utils/database"
	env "payment_service/utils/env"
	"payment_service/utils/jwt"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateNewAccount(t *testing.T) {
	env.Load()

	database := db.Open()
	db.CleanupDB(database)
	db.PopulateDB(database, 100)
	defer database.Close()

	r := mux.NewRouter()
	r.HandleFunc("/accounts", CreateNewAccount).Methods("POST")

	body := strings.NewReader(`{"owner":"ruben"}`)
	req, err := http.NewRequest("POST", "/accounts", body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestGetAccountDetails(t *testing.T) {
	env.Load()
	database := db.Open()
	db.CleanupDB(database)
	db.PopulateDB(database, 100)
	defer database.Close()

	r := mux.NewRouter()
	r.HandleFunc("/account", GetAccountDetails).Methods("GET")
	r.Use(jwt.AuthMiddleware)

	req, err := http.NewRequest("GET", "/account", nil)
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

	t.Log(res.Body.String())

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteAccount(t *testing.T) {
	env.Load()

	database := db.Open()
	db.CleanupDB(database)
	db.PopulateDB(database, 100)
	defer database.Close()

	r := mux.NewRouter()
	r.HandleFunc("/account", DeleteAccount).Methods("DELETE")
	r.Use(jwt.AuthMiddleware)

	req, err := http.NewRequest("DELETE", "/account", nil)
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
