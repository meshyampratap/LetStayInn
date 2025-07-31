package models

type Room struct {
	ID          string  `json:"id"`
	Number      int     `json:"number"`
	Type        string  `json:"type"` // e.g. Single, Double, Suite
	Price       float64 `json:"price"`
	IsAvailable bool    `json:"is_available"`
	Description string  `json:"description"`
}
