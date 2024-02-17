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

	router := mux.NewRouter()

	router.HandleFunc("/health", routes.Health)
	router.HandleFunc("/accounts", routes.CreateNewAccount)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	fmt.Println("Server listening to port", port)
	http.ListenAndServe(":"+port, loggedRouter)

}
