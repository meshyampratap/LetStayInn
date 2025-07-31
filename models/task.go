package models

import "time"

type Task struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"` // cleaning, food
	AssignedTo string    `json:"assigned_to"`
	BookingID  string    `json:"booking_id"`
	Status     string    `json:"status"` // Pending, In Progress, Done
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Details    string    `json:"details"`
}
