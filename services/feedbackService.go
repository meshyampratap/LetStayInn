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

func SubmitFeedback(user *models.User) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- Submit Feedback ---")
	fmt.Print("Enter your feedback: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)
	if message == "" {
		fmt.Println("Feedback cannot be empty.")
		return
	}
	var feedbacks []models.Feedback
	utils.ReadJSON("data/feedback.json", &feedbacks)
	id := fmt.Sprintf("f%d", len(feedbacks)+1)
	feedback := models.Feedback{
		ID:        id,
		UserID:    user.ID,
		Message:   message,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	feedbacks = append(feedbacks, feedback)
	utils.WriteJSON("data/feedback.json", feedbacks)
	fmt.Println("Thank you for your feedback!")
}

func RequestService(user *models.User, reqType string) {
	var requests []models.ServiceRequest
	utils.ReadJSON("data/service_requests.json", &requests)
	id := fmt.Sprintf("sr%d", len(requests)+1)
	request := models.ServiceRequest{
		ID:        id,
		UserID:    user.ID,
		Type:      reqType,
		Status:    "pending",
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	requests = append(requests, request)
	utils.WriteJSON("data/service_requests.json", requests)
	fmt.Printf("%s request submitted!\n", strings.Title(reqType))
}
