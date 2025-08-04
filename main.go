package main

import (
	"fmt"
	"hotel-management-cli/services"
	"hotel-management-cli/utils"
	"os"
)

func main() {
	utils.EnsureDataFilesExist()
	for {
		fmt.Printf("\n%s\n", utils.ColorTitle("--- HOTEL MANAGEMENT CLI ---"))
		fmt.Println(utils.ColorPrompt("1. Login"))
		fmt.Println(utils.ColorPrompt("2. Signup"))
		fmt.Println(utils.ColorPrompt("3. Exit"))
		fmt.Print(utils.ColorInfo("Select option: "))
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			services.Login()
		case 2:
			services.Signup()
		case 3:
			fmt.Println(utils.ColorInfo("Exiting..."))
			os.Exit(0)
		default:
			fmt.Println(utils.ColorError("Invalid option. Try again."))
		}
	}
}
