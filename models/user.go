package models

import (
	"crypto/rand"
	"encoding/base64"
)

type User struct {
	ID       int      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	UserType string   `json:"userType"`
	Token    string   `json:"token"`
	Tickets  []Ticket `json:"tickets"` // A user has many tickets
}

func GenerateRandomToken() (string, error) {
	// Generate a random byte slice of a suitable length (e.g., 16 bytes).
	tokenBytes := make([]byte, 16)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64 to create the token.
	token := base64.StdEncoding.EncodeToString(tokenBytes)
	return token, nil
}
