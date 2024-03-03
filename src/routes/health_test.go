package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestHealthEndpoint(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/health", Health).Methods("GET")

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":"OK","timestamp":"`
	if !strings.Contains(res.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", res.Body.String(), expected)
	}
}
