package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/ManinderSinghRajput/MakerCheker/pkg/model"
)

// Storage for messages (mock database)
var (
	messages = make(map[string]*model.Message)
	mu       sync.Mutex
)

// SubmitMessageHandler handles message submissions.
func SubmitMessageHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Message   string `json:"message"`
		Recipient string `json:"recipient"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Message == "" || req.Recipient == "" {
		http.Error(w, "Message and recipient are required", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	id := generateID()
	message := &model.Message{
		ID:        id,
		Message:   req.Message,
		Recipient: req.Recipient,
		Status:    "pending",
	}
	messages[id] = message

	//w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(message)
	respondWithJSON(w, http.StatusCreated, message)
}

// GetMessageHandler retrieves the status of a message.
func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	mu.Lock()
	defer mu.Unlock()

	message, exists := messages[id]
	if !exists {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	//json.NewEncoder(w).Encode(message)
	respondWithJSON(w, http.StatusOK, message)
}

// UpdateMessageStatusHandler handles message approval/rejection.
func UpdateMessageStatusHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var req struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Status != "approved" && req.Status != "rejected" {
		http.Error(w, "Invalid status, must be 'approved' or 'rejected'", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	message, exists := messages[id]
	if !exists {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	if message.Status != "pending" {
		http.Error(w, "Cannot update status of a non-pending message", http.StatusBadRequest)
		return
	}

	message.Status = req.Status
	respondWithJSON(w, http.StatusOK, message)
	//json.NewEncoder(w).Encode(message)
}

// generateID generates a UUID for messages.
func generateID() string {
	return uuid.New().String()
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Format JSON with indentation
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
