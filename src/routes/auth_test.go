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

func TestAuthorizeAccount(t *testing.T) {
	env.Load()

	database := db.Open()
	db.CleanupDB(database)
	db.PopulateDB(database, 100)
	defer database.Close()

	r := mux.NewRouter()
	r.HandleFunc("/auth", AuthorizeAccount).Methods("POST")

	payload := jwt.Payload{
		Id:    1,
		Owner: "Owner1",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/auth", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var jsonRes map[string]string
	err = json.NewDecoder(res.Body).Decode(&jsonRes)
	if err != nil {
		t.Fatal("Failed to decode response body")
	}

	if jsonRes["token"] == "" {
		t.Error("Expected a token, got none")
	}
}
