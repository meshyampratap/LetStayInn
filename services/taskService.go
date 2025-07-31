package services

import (
	"fmt"
	"hotel-management-cli/models"
	"hotel-management-cli/utils"
)

func ViewAssignedTasks(user *models.User) {
	var tasks []models.Task
	utils.ReadJSON("data/tasks.json", &tasks)
	fmt.Println("\n--- Assigned Tasks ---")
	for _, t := range tasks {
		if t.AssignedTo == user.ID {
			fmt.Printf("Task ID: %s, Type: %s, Status: %s, Details: %s\n", t.ID, t.Type, t.Status, t.Details)
		}
	}
}

func UpdateTaskStatus(user *models.User) {
	var tasks []models.Task
	utils.ReadJSON("data/tasks.json", &tasks)
	fmt.Println("\n--- Update Task Status ---")
	for _, t := range tasks {
		if t.AssignedTo == user.ID {
			fmt.Printf("Task ID: %s, Status: %s\n", t.ID, t.Status)
		}
	}
	var tid, status string
	fmt.Print("Enter Task ID: ")
	fmt.Scanln(&tid)
	fmt.Print("Enter new status (Pending/In Progress/Done): ")
	fmt.Scanln(&status)
	updated := false
	for i, t := range tasks {
		if t.ID == tid && t.AssignedTo == user.ID {
			tasks[i].Status = status
			tasks[i].UpdatedAt = t.UpdatedAt
			updated = true
			break
		}
	}
	if updated {
		utils.WriteJSON("data/tasks.json", tasks)
		fmt.Println("Task status updated.")
	} else {
		fmt.Println("Task not found.")
	}
}
