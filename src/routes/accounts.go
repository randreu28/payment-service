package routes

import (
	"encoding/json"
	"net/http"
	db "payment_service/utils/database"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type AccountRequest struct {
	AccountOwner string `json:"owner"`
}

func CreateNewAccount(res http.ResponseWriter, req *http.Request) {
	var payload AccountRequest
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if len(payload.AccountOwner) > 1000 {
		http.Error(res, "Account owner exceeds 1000 characters", http.StatusBadRequest)
		return
	}

	db := db.Open()
	defer db.Close()

	var id int
	err = db.QueryRow("INSERT INTO accounts (account_owner) VALUES ($1) RETURNING id", payload.AccountOwner).Scan(&id)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]string{
		"status":  "OK",
		"message": "Account successfully created",
		"id":      strconv.Itoa(id),
	})

}
func GetAccountDetails(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	unparsedAccountId := vars["id"]

	accountId, err := strconv.Atoi(unparsedAccountId)

	if err != nil {
		http.Error(res, "Account ID must be a number", http.StatusBadRequest)
		return
	}

	db := db.Open()
	defer db.Close()

	var id int
	var accountOwner string
	var balance string
	var createdAt time.Time
	var updatedAt time.Time

	err = db.QueryRow("SELECT * FROM accounts WHERE id = $1", accountId).Scan(&id, &accountOwner, &balance, &createdAt, &updatedAt)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]string{
		"account_id":    strconv.Itoa(id),
		"account_owner": accountOwner,
		"balance":       balance,
		"created_at":    createdAt.Format(time.RFC3339),
		"updated_at":    updatedAt.Format(time.RFC3339),
	})
}
