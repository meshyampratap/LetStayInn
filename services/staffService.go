package services

import (
	"fmt"
	"hotel-management-cli/models"
	"hotel-management-cli/utils"
)

func ViewAssignedServiceRequests(user *models.User, reqType string) {
	var tasks []models.Task
	utils.ReadJSON("data/tasks.json", &tasks)
	fmt.Printf("\n--- Assigned %s Requests ---\n", reqType)
	for _, t := range tasks {
		if t.AssignedTo == user.ID && t.Type == reqType {
			fmt.Printf("Task ID: %s | Status: %s | Details: %s\n", t.ID, t.Status, t.Details)
		}
	}
}

func UpdateAssignedServiceRequestStatus(user *models.User, reqType string) {
	var tasks []models.Task
	utils.ReadJSON("data/tasks.json", &tasks)
	fmt.Printf("\n--- Update %s Request Status ---\n", reqType)
	for _, t := range tasks {
		if t.AssignedTo == user.ID && t.Type == reqType {
			fmt.Printf("Task ID: %s | Status: %s\n", t.ID, t.Status)
		}
	}
	fmt.Print("Enter Task ID: ")
	var tid, status string
	fmt.Scanln(&tid)
	fmt.Print("Enter new status (Pending/In Progress/Done): ")
	fmt.Scanln(&status)
	updated := false
	for i, t := range tasks {
		if t.ID == tid && t.AssignedTo == user.ID && t.Type == reqType {
			tasks[i].Status = status
			tasks[i].UpdatedAt = t.UpdatedAt
			updated = true
			break
		}
	}
	if updated {
		utils.WriteJSON("data/tasks.json", tasks)
		fmt.Println("Request status updated.")
	} else {
		fmt.Println("Task not found.")
	}
}
