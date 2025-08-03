package services

import (
	"fmt"
	"hotel-management-cli/models"
	"hotel-management-cli/utils"
)

func ListAvailableRooms() {
	var rooms []models.Room
	utils.ReadJSON("data/rooms.json", &rooms)
	fmt.Println("\n--- Available Rooms ---")
	for _, r := range rooms {
		if r.IsAvailable {
			fmt.Printf("Room %d (%s): Rs.%.2f - %s\n", r.Number, r.Type, r.Price, r.Description)
		}
	}
}

func AddRoom() {
	var rooms []models.Room
	utils.ReadJSON("data/rooms.json", &rooms)
	var number int
	var rtype, desc string
	var price float64
	fmt.Print("Enter room number: ")
	fmt.Scanln(&number)
	fmt.Print("Enter room type: ")
	fmt.Scanln(&rtype)
	fmt.Print("Enter price: ")
	fmt.Scanln(&price)
	fmt.Print("Enter description: ")
	fmt.Scanln(&desc)
	id := fmt.Sprintf("r%d", len(rooms)+1)
	room := models.Room{
		ID:          id,
		Number:      number,
		Type:        rtype,
		Price:       price,
		IsAvailable: true,
		Description: desc,
	}
	rooms = append(rooms, room)
	utils.WriteJSON("data/rooms.json", rooms)
	fmt.Println("Room added!")
}
