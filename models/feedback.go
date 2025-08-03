package models

type Feedback struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type ServiceRequest struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Type      string `json:"type"` // food or cleaning
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
