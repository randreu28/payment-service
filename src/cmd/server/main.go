package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"payment_service/routes"
	db "payment_service/utils/database"
	"payment_service/utils/env"
	"syscall"
	"time"

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
	router.HandleFunc("/accounts", routes.CreateNewAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", routes.GetAccountDetails).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: loggedRouter,
	}

	log.Println("Server listening on port", port)
	go func() {
		err := server.ListenAndServe()

		if err == http.ErrServerClosed {
			log.Println("Shutting down server...")
			return
		}

		if err != nil {
			log.Fatal("Error starting server")
			panic(err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan // Blocks main() until signal is received

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Printf("Server Shutdown Failed:%+v", err)
	}

	db.Close()

	log.Println("Server and database connection gracefully closed")
}
