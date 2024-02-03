package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Define a new type for the JSON object
type Response struct {
	Value int    `json:"value"`
	Test  string `json:"test"`
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	numStr := req.URL.Query().Get("number")

	if numStr == "" {
		http.Error(res, "Missing 'number' query parameter", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(numStr)

	if err != nil {
		http.Error(res, "'number' query parameter must be an integer", http.StatusBadRequest)
		return
	}

	response := Response{
		Value: num,
		Test:  "Yes, it works!",
	}

	json.NewEncoder(res).Encode(response)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
