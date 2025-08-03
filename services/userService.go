package services

import (
   "bufio"
   "fmt"
   "hotel-management-cli/models"
   "hotel-management-cli/utils"
   "os"
   "strings"
   "time"
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
		   RequestService(user, "food")
	   case 6:
		   RequestService(user, "cleaning")
	   case 7:
		   SubmitFeedback(user)
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
		   ViewAssignedServiceRequests(user, "food")
	   case 2:
		   UpdateAssignedServiceRequestStatus(user, "food")
	   case 3:
		   toggleAvailability(user)
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
		   ViewAssignedServiceRequests(user, "cleaning")
	   case 2:
		   UpdateAssignedServiceRequestStatus(user, "cleaning")
	   case 3:
		   toggleAvailability(user)
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
		   for {
			   fmt.Println("\n--- Room Management ---")
			   fmt.Println("1. List Rooms")
			   fmt.Println("2. Add Room")
			   fmt.Println("3. Update Room")
			   fmt.Println("4. Delete Room")
			   fmt.Println("5. Back")
			   fmt.Print("Select option: ")
			   var rchoice int
			   fmt.Scanln(&rchoice)
			   switch rchoice {
			   case 1:
				   ListRooms()
			   case 2:
				   AddRoom()
			   case 3:
				   UpdateRoom()
			   case 4:
				   DeleteRoom()
			   case 5:
				   break
			   default:
				   fmt.Println("Invalid option.")
			   }
			   if rchoice == 5 {
				   break
			   }
		   }
	   case 3:
		   ListBookingsAndGuests()
	   case 4:
		   for {
			   fmt.Println("\n--- Employee Management ---")
			   fmt.Println("1. List Employees")
			   fmt.Println("2. Update Employee Availability")
			   fmt.Println("3. Delete Employee")
			   fmt.Println("4. Back")
			   fmt.Print("Select option: ")
			   var echoice int
			   fmt.Scanln(&echoice)
			   switch echoice {
			   case 1:
				   ListEmployees()
			   case 2:
				   UpdateEmployeeAvailability()
			   case 3:
				   DeleteEmployee()
			   case 4:
				   break
			   default:
				   fmt.Println("Invalid option.")
			   }
			   if echoice == 4 {
				   break
			   }
		   }
	   case 5:
		   AssignTaskToEmployee("cleaning")
	   case 6:
		   AssignTaskToEmployee("food")
	   case 7:
		   ListServiceRequests()
	   case 8:
		   GenerateReport()
	   case 9:
		   fmt.Println("Logging out...")
		   return
	   default:
		   fmt.Println("Invalid option.")
	   }
	}
}

// toggleAvailability allows staff to mark themselves available/unavailable
func toggleAvailability(user *models.User) {
   var users []models.User
   utils.ReadJSON("data/users.json", &users)
   for i, u := range users {
	   if u.ID == user.ID {
		   users[i].Available = !u.Available
		   user.Available = users[i].Available // update current session
		   utils.WriteJSON("data/users.json", users)
		   status := "unavailable"
		   if users[i].Available {
			   status = "available"
		   }
		   fmt.Printf("You are now marked as %s.\n", status)
		   return
	   }
   }
   fmt.Println("User not found.")
}
