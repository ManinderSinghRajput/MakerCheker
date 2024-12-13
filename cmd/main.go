package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ManinderSinghRajput/MakerCheker/internal/handler"
)

func main() {
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/messages", handler.SubmitMessageHandler).Methods(http.MethodPost)
	r.HandleFunc("/messages/{id}/status", handler.UpdateMessageStatusHandler).Methods(http.MethodPost)
	r.HandleFunc("/messages/{id}", handler.GetMessageHandler).Methods(http.MethodGet)

	// Start server
	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
