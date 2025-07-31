package main

import (
	"fmt"
	"os"
	"hotel-management-cli/services"
	"hotel-management-cli/utils"
)

func main() {
	utils.EnsureDataFilesExist()
	for {
		fmt.Println("\n--- HOTEL MANAGEMENT CLI ---")
		fmt.Println("1. Login")
		fmt.Println("2. Signup")
		fmt.Println("3. Exit")
		fmt.Print("Select option: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			services.Login()
		case 2:
			services.Signup()
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
