package routes

import (
	"encoding/json"
	"net/http"
	db "payment_service/utils/database"
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

	_, err = db.Exec("INSERT INTO accounts (account_owner) values ($1)", payload.AccountOwner)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]string{
		"status":  "OK",
		"message": "Account successfully created",
	})

}
