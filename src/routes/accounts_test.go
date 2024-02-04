package routes

import (
	"net/http"
	"net/http/httptest"
	"payment_service/utils/env"
	"strings"
	"testing"
)

func TestCreateNewAccount(t *testing.T) {
	env.Load("../../.env.local")

	body := strings.NewReader(`{"owner":"John Doe"}`)
	req, err := http.NewRequest("POST", "/account", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateNewAccount)

	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
