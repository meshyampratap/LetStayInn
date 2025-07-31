package models

import "time"

// UserRole type for role-based access
const (
	RoleGuest           = "guest"
	RoleKitchenStaff    = "kitchen_staff"
	RoleCleaningStaff   = "cleaning_staff"
	RoleManager         = "manager"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	Available bool      `json:"available"` // for staff
}
