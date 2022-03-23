package dto

import (
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
)

type GetAllUserResponse struct {
	Users []models.User `json:"users"`
}

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type CreateUserRequest struct {
	PhoneNumber string `json:"phone_number"`
}
