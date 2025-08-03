package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"hotel-management-cli/models"
)

func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

func CheckPassword(hash, password string) bool {
	return hash == HashPassword(password)
}

func FindUserByEmail(users []models.User, email string) *models.User {
	for _, u := range users {
		if u.Email == email {
			return &u
		}
	}
	return nil
}
