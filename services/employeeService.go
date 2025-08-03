package services

import (
	"fmt"
	"hotel-management-cli/models"
	"hotel-management-cli/utils"
)

func ListBookingsAndGuests() {
	var bookings []models.Booking
	var users []models.User
	var rooms []models.Room
	utils.ReadJSON("data/bookings.json", &bookings)
	utils.ReadJSON("data/users.json", &users)
	utils.ReadJSON("data/rooms.json", &rooms)
	fmt.Println("\n--- All Bookings and Guests ---")
	for _, b := range bookings {
		var guestName, roomNum string
		for _, u := range users {
			if u.ID == b.UserID {
				guestName = u.Name
				break
			}
		}
		for _, r := range rooms {
			if r.ID == b.RoomID {
				roomNum = fmt.Sprintf("%d", r.Number)
				break
			}
		}
		fmt.Printf("Booking ID: %s | Guest: %s | Room: %s | Status: %s | Check-in: %s | Check-out: %s\n", b.ID, guestName, roomNum, b.Status, b.CheckIn.Format("2006-01-02"), b.CheckOut.Format("2006-01-02"))
	}
}

func ListEmployees() {
	var users []models.User
	utils.ReadJSON("data/users.json", &users)
	fmt.Println("\n--- Employees ---")
	for _, u := range users {
		if u.Role != models.RoleGuest {
			fmt.Printf("ID: %s | Name: %s | Role: %s | Available: %v\n", u.ID, u.Name, u.Role, u.Available)
		}
	}
}

func AddEmployee() {
	fmt.Println("Use Signup from main menu to add employees.")
}

func UpdateEmployeeAvailability() {
	var users []models.User
	utils.ReadJSON("data/users.json", &users)
	fmt.Print("Enter employee ID to update: ")
	var eid string
	fmt.Scanln(&eid)
	found := false
	for i, u := range users {
		if u.ID == eid && u.Role != models.RoleGuest {
			users[i].Available = !users[i].Available
			found = true
			break
		}
	}
	if found {
		utils.WriteJSON("data/users.json", users)
		fmt.Println("Employee availability updated!")
	} else {
		fmt.Println("Employee not found.")
	}
}

func DeleteEmployee() {
	var users []models.User
	utils.ReadJSON("data/users.json", &users)
	fmt.Print("Enter employee ID to delete: ")
	var eid string
	fmt.Scanln(&eid)
	idx := -1
	for i, u := range users {
		if u.ID == eid && u.Role != models.RoleGuest {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("Employee not found.")
		return
	}
	users = append(users[:idx], users[idx+1:]...)
	utils.WriteJSON("data/users.json", users)
	fmt.Println("Employee deleted!")
}
