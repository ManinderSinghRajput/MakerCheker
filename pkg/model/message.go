package model

// Message represents a message entity.
type Message struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Recipient string `json:"recipient"`
	Status    string `json:"status"` // "pending", "approved", "rejected"
}
