package presenter

import (
	"github.com/gofiber/fiber/v2"
	"synapsis-challenge/internal/entities"
)

// User is the presenter object which will be passed in the response by Handler
type User struct {
	ID       int    `json:"id" bson:"_id,omitempty"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// AuthSuccesResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func AuthSuccesResponse(data *entities.User, token string) *fiber.Map {
	user := User{
		ID:       data.ID,
		Username: data.Username,
		Token:    token,
	}
	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}

// AuthErrorResponse is the ErrorResponse that will be passed in the response by Handler
func AuthErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
