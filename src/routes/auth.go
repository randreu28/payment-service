package routes

import (
	"encoding/json"
	"net/http"
	db "payment_service/utils/database"
	jwt "payment_service/utils/jwt"
)

// AuthorizeAccount Checks if a user is authorized to access an account, and returns a JWT if pertinent
func AuthorizeAccount(res http.ResponseWriter, req *http.Request) {
	var payload jwt.Payload
	err := json.NewDecoder(req.Body).Decode(&payload)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	db := db.Open()
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM accounts WHERE id = $1 AND account_owner = $2", payload.Id, payload.Owner).Scan(&count)
	if err != nil || count == 0 {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	jwt, err := jwt.CreateJwt(payload.Id, payload.Owner)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	jsonRes := map[string]string{
		"token": jwt,
	}

	json.NewEncoder(res).Encode(jsonRes)
}
