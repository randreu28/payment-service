package main

import (
	"fmt"
	"net/http"
	"os"
	"payment_service/routes"
	"payment_service/utils/env"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	env.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := mux.NewRouter()

	r.HandleFunc("/health", routes.Health)
	// TODO: Add more routes here

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	fmt.Println("Server listening to port", port)
	http.ListenAndServe(":"+port, loggedRouter)

}
