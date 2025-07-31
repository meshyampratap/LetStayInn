package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"hotel-management-cli/models"
	"hotel-management-cli/utils"
)

func Signup() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- Signup ---")
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Println("Select role: 1.Guest 2.Kitchen Staff 3.Room Cleaning Staff 4.Manager")
	fmt.Print("Enter choice: ")
	var roleChoice int
	fmt.Scanln(&roleChoice)

	var role string
	switch roleChoice {
	case 1:
		role = models.RoleGuest
	case 2:
		role = models.RoleKitchenStaff
	case 3:
		role = models.RoleCleaningStaff
	case 4:
		role = models.RoleManager
	default:
		fmt.Println("Invalid role.")
		return
	}

	var users []models.User
	utils.ReadJSON("data/users.json", &users)

	if utils.FindUserByEmail(users, email) != nil {
		fmt.Println("Email already exists.")
		return
	}

	id := fmt.Sprintf("u%d", len(users)+1)
	newUser := models.User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  utils.HashPassword(password),
		Role:      role,
		CreatedAt: time.Now(),
		Available: role != models.RoleGuest,
	}

	users = append(users, newUser)
	utils.WriteJSON("data/users.json", users)
	fmt.Println("Signup successful! Please login.")
}

func Login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- Login ---")
	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	var users []models.User
	utils.ReadJSON("data/users.json", &users)
	user := utils.FindUserByEmail(users, email)

	if user == nil || !utils.CheckPassword(user.Password, password) {
		fmt.Println("Invalid credentials.")
		return
	}

	fmt.Printf("Welcome, %s!\n", user.Name)

	switch user.Role {
	case models.RoleGuest:
		GuestDashboard(user)
	case models.RoleKitchenStaff:
		KitchenDashboard(user)
	case models.RoleCleaningStaff:
		CleaningDashboard(user)
	case models.RoleManager:
		ManagerDashboard(user)
	default:
		fmt.Println("Unknown role.")
	}
}

func GuestDashboard(user *models.User) {
	for {
		fmt.Println("\n--- Guest Dashboard ---")
		fmt.Println("1. View Available Rooms")
		fmt.Println("2. Book Room")
		fmt.Println("3. Cancel Booking")
		fmt.Println("4. View My Bookings")
		fmt.Println("5. Request Food")
		fmt.Println("6. Request Room Cleaning")
		fmt.Println("7. Give Feedback")
		fmt.Println("8. Logout")
		fmt.Print("Select option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			ListAvailableRooms()
		case 2:
			BookRoom(user)
		case 3:
			CancelBooking(user)
		case 4:
			ViewMyBookings(user)
		case 5:
			fmt.Println("Food request feature coming soon.")
		case 6:
			fmt.Println("Room cleaning request feature coming soon.")
		case 7:
			fmt.Println("Feedback feature coming soon.")
		case 8:
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func KitchenDashboard(user *models.User) {
	for {
		fmt.Println("\n--- Kitchen Staff Dashboard ---")
		fmt.Println("1. View Assigned Food Requests")
		fmt.Println("2. Update Food Order Status")
		fmt.Println("3. Mark Availability")
		fmt.Println("4. Logout")
		fmt.Print("Select option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("View assigned food requests feature coming soon.")
		case 2:
			fmt.Println("Update food order status feature coming soon.")
		case 3:
			fmt.Println("Mark availability feature coming soon.")
		case 4:
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func CleaningDashboard(user *models.User) {
	for {
		fmt.Println("\n--- Room Cleaning Staff Dashboard ---")
		fmt.Println("1. View Assigned Cleaning Requests")
		fmt.Println("2. Update Task Status")
		fmt.Println("3. Mark Availability")
		fmt.Println("4. Logout")
		fmt.Print("Select option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			ViewAssignedTasks(user)
		case 2:
			UpdateTaskStatus(user)
		case 3:
			fmt.Println("Mark availability feature coming soon.")
		case 4:
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func ManagerDashboard(user *models.User) {
	for {
		fmt.Println("\n--- Manager Dashboard ---")
		fmt.Println("1. View Dashboard Summary")
		fmt.Println("2. Room Management")
		fmt.Println("3. View Bookings and Guests")
		fmt.Println("4. Manage Employees")
		fmt.Println("5. Assign Cleaning Tasks")
		fmt.Println("6. Assign Food Requests")
		fmt.Println("7. View Guest Service Requests")
		fmt.Println("8. Generate Reports")
		fmt.Println("9. Logout")
		fmt.Print("Select option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Dashboard summary feature coming soon.")
		case 2:
			fmt.Println("Room management feature coming soon.")
		case 3:
			fmt.Println("View bookings and guests feature coming soon.")
		case 4:
			fmt.Println("Manage employees feature coming soon.")
		case 5:
			fmt.Println("Assign cleaning tasks feature coming soon.")
		case 6:
			fmt.Println("Assign food requests feature coming soon.")
		case 7:
			fmt.Println("View guest service requests feature coming soon.")
		case 8:
			fmt.Println("Generate reports feature coming soon.")
		case 9:
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}
