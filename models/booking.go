package models

import "time"

type Booking struct {
	ID       string    `json:"id"`
	UserID   string    `json:"user_id"`
	RoomID   string    `json:"room_id"`
	CheckIn  time.Time `json:"check_in"`
	CheckOut time.Time `json:"check_out"`
	Status   string    `json:"status"` // Booked, Cancelled
	FoodReq  []string  `json:"food_requests"`
	CleanReq bool      `json:"clean_request"`
	Feedback string    `json:"feedback"`
}
