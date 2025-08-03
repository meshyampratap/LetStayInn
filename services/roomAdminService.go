package services

import (
	"bufio"
	"fmt"
	"hotel-management-cli/models"
	"hotel-management-cli/utils"
	"os"
	"strings"
)

func ListRooms() {
	var rooms []models.Room
	utils.ReadJSON("data/rooms.json", &rooms)
	fmt.Println("\n--- All Rooms ---")
	for _, r := range rooms {
		fmt.Printf("Room %d (%s): Rs.%.2f - %s | Available: %v\n", r.Number, r.Type, r.Price, r.Description, r.IsAvailable)
	}
}

func UpdateRoom() {
	var rooms []models.Room
	utils.ReadJSON("data/rooms.json", &rooms)
	fmt.Print("Enter room number to update: ")
	var number int
	fmt.Scanln(&number)
	var found *models.Room
	for i, r := range rooms {
		if r.Number == number {
			found = &rooms[i]
			break
		}
	}
	if found == nil {
		fmt.Println("Room not found.")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Current type: %s. Enter new type (or press Enter to keep): ", found.Type)

	typeStr, _ := reader.ReadString('\n')
	typeStr = strings.TrimSpace(typeStr)
	if typeStr != "" {
		found.Type = typeStr
	}
	fmt.Printf("Current price: %.2f. Enter new price (or press Enter to keep): ", found.Price)
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	if priceStr != "" {
		fmt.Sscanf(priceStr, "%f", &found.Price)
	}
	fmt.Printf("Current description: %s. Enter new description (or press Enter to keep): ", found.Description)
	desc, _ := reader.ReadString('\n')
	desc = strings.TrimSpace(desc)
	if desc != "" {
		found.Description = desc
	}
	utils.WriteJSON("data/rooms.json", rooms)
	fmt.Println("Room updated!")
}

func DeleteRoom() {
	var rooms []models.Room
	utils.ReadJSON("data/rooms.json", &rooms)
	fmt.Print("Enter room number to delete: ")
	var number int
	fmt.Scanln(&number)
	idx := -1
	for i, r := range rooms {
		if r.Number == number {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("Room not found.")
		return
	}
	rooms = append(rooms[:idx], rooms[idx+1:]...)
	utils.WriteJSON("data/rooms.json", rooms)
	fmt.Println("Room deleted!")
}
