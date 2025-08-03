package services

import (
	"fmt"
	"hotel-management-cli/models"
	"hotel-management-cli/utils"
)

func GenerateReport() {
	var bookings []models.Booking
	var feedbacks []models.Feedback
	var tasks []models.Task
	utils.ReadJSON("data/bookings.json", &bookings)
	utils.ReadJSON("data/feedback.json", &feedbacks)
	utils.ReadJSON("data/tasks.json", &tasks)
	fmt.Println("\n--- Hotel Report ---")
	fmt.Printf("Total Bookings: %d\n", len(bookings))
	fmt.Printf("Total Feedbacks: %d\n", len(feedbacks))
	fmt.Printf("Total Tasks Assigned: %d\n", len(tasks))
	booked, cancelled := 0, 0
	for _, b := range bookings {
		if b.Status == "Booked" {
			booked++
		} else if b.Status == "Cancelled" {
			cancelled++
		}
	}
	fmt.Printf("Active Bookings: %d | Cancelled Bookings: %d\n", booked, cancelled)

	pending, done := 0, 0
	for _, t := range tasks {
		if t.Status == "Done" {
			done++
		} else {
			pending++
		}
	}
	fmt.Printf("Pending Tasks: %d | Completed Tasks: %d\n", pending, done)
}
