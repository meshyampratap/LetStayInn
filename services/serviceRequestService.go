package services

import (
	"fmt"
	"hotel-management-cli/models"
	"hotel-management-cli/utils"
	"time"
)

func ListServiceRequests() {
	var requests []models.ServiceRequest
	var users []models.User
	utils.ReadJSON("data/service_requests.json", &requests)
	utils.ReadJSON("data/users.json", &users)
	fmt.Println("\n--- Guest Service Requests ---")
	for _, r := range requests {
		var uname string
		for _, u := range users {
			if u.ID == r.UserID {
				uname = u.Name
				break
			}
		}
		fmt.Printf("ID: %s | Guest: %s | Type: %s | Status: %s | Created: %s\n", r.ID, uname, r.Type, r.Status, r.CreatedAt)
	}
}

func AssignTaskToEmployee(taskType string) {
	var users []models.User
	var requests []models.ServiceRequest
	var tasks []models.Task
	utils.ReadJSON("data/users.json", &users)
	utils.ReadJSON("data/service_requests.json", &requests)
	utils.ReadJSON("data/tasks.json", &tasks)
	fmt.Printf("\n--- Assign %s Task ---\n", taskType)
	for _, r := range requests {
		if r.Type == taskType && r.Status == "pending" {
			fmt.Printf("Request ID: %s | Guest: %s\n", r.ID, r.UserID)
		}
	}
	fmt.Print("Enter Request ID to assign: ")
	var rid string
	fmt.Scanln(&rid)
	var req *models.ServiceRequest
	for i, r := range requests {
		if r.ID == rid && r.Type == taskType && r.Status == "pending" {
			req = &requests[i]
			break
		}
	}
	if req == nil {
		fmt.Println("Request not found.")
		return
	}
	fmt.Println("Available employees:")
	for _, u := range users {
		if ((taskType == "food" && u.Role == models.RoleKitchenStaff) || (taskType == "cleaning" && u.Role == models.RoleCleaningStaff)) && u.Available {
			fmt.Printf("ID: %s | Name: %s\n", u.ID, u.Name)
		}
	}
	fmt.Print("Enter Employee ID to assign: ")
	var eid string
	fmt.Scanln(&eid)
	assigned := false
	for _, u := range users {
		if u.ID == eid && u.Available && ((taskType == "food" && u.Role == models.RoleKitchenStaff) || (taskType == "cleaning" && u.Role == models.RoleCleaningStaff)) {
			task := models.Task{
				ID:         fmt.Sprintf("t%d", len(tasks)+1),
				Type:       taskType,
				AssignedTo: eid,
				BookingID:  "",
				Status:     "Pending",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
				Details:    "Assigned from service request " + req.ID,
			}
			tasks = append(tasks, task)
			req.Status = "assigned"
			assigned = true
			break
		}
	}
	if assigned {
		utils.WriteJSON("data/tasks.json", tasks)
		utils.WriteJSON("data/service_requests.json", requests)
		fmt.Println("Task assigned!")
	} else {
		fmt.Println("Employee not found or not available.")
	}
}
