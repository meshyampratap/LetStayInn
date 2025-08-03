package services

import (
	"fmt"
	"hotel-management-cli/models"
	"hotel-management-cli/utils"
	"time"
)

func BookRoom(user *models.User) {
	var rooms []models.Room
	utils.ReadJSON("data/rooms.json", &rooms)
	fmt.Println("\n--- Book Room ---")
	for _, r := range rooms {
		if r.IsAvailable {
			fmt.Printf("Room %d (%s): Rs.%.2f - %s\n", r.Number, r.Type, r.Price, r.Description)
		}
	}
	var roomNum int
	fmt.Print("Enter room number to book: ")
	fmt.Scanln(&roomNum)
	var selected *models.Room
	for i, r := range rooms {
		if r.Number == roomNum && r.IsAvailable {
			selected = &rooms[i]
			break
		}
	}
	if selected == nil {
		fmt.Println("Room not available.")
		return
	}
	var checkIn, checkOut string
	fmt.Print("Enter check-in date (YYYY-MM-DD): ")
	fmt.Scanln(&checkIn)
	fmt.Print("Enter check-out date (YYYY-MM-DD): ")
	fmt.Scanln(&checkOut)
	ci, _ := time.Parse("2006-01-02", checkIn)
	co, _ := time.Parse("2006-01-02", checkOut)
	var bookings []models.Booking
	utils.ReadJSON("data/bookings.json", &bookings)
	id := fmt.Sprintf("b%d", len(bookings)+1)
	booking := models.Booking{
		ID:       id,
		UserID:   user.ID,
		RoomID:   selected.ID,
		CheckIn:  ci,
		CheckOut: co,
		Status:   "Booked",
	}
	bookings = append(bookings, booking)
	utils.WriteJSON("data/bookings.json", bookings)
	selected.IsAvailable = false
	utils.WriteJSON("data/rooms.json", rooms)
	fmt.Println("Room booked!")
}

func CancelBooking(user *models.User) {
	var bookings []models.Booking
	utils.ReadJSON("data/bookings.json", &bookings)
	fmt.Println("\n--- Cancel Booking ---")
	userBookings := []models.Booking{}
	for _, b := range bookings {
		if b.UserID == user.ID && b.Status == "Booked" {
			userBookings = append(userBookings, b)
		}
	}
	if len(userBookings) == 0 {
		fmt.Println("No active bookings.")
		return
	}
	for _, b := range userBookings {
		fmt.Printf("Booking ID: %s, Room: %s, Check-in: %s\n", b.ID, b.RoomID, b.CheckIn.Format("2006-01-02"))
	}
	var bid string
	fmt.Print("Enter Booking ID to cancel: ")
	fmt.Scanln(&bid)
	cancelled := false
	for i, b := range bookings {
		if b.ID == bid && b.UserID == user.ID && b.Status == "Booked" {
			bookings[i].Status = "Cancelled"
			cancelled = true
			break
		}
	}
	if cancelled {
		utils.WriteJSON("data/bookings.json", bookings)
		fmt.Println("Booking cancelled.")
	} else {
		fmt.Println("Booking not found.")
	}
}

func ViewMyBookings(user *models.User) {
	var bookings []models.Booking
	utils.ReadJSON("data/bookings.json", &bookings)
	fmt.Println("\n--- My Bookings ---")
	for _, b := range bookings {
		if b.UserID == user.ID {
			fmt.Printf("Booking ID: %s, Room: %s, Status: %s, Check-in: %s\n", b.ID, b.RoomID, b.Status, b.CheckIn.Format("2006-01-02"))
		}
	}
}
