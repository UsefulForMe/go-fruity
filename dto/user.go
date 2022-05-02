package dto

import (
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
)

type GetAllUserResponse struct {
	Users []models.User `json:"users"`
}

type CreateUserResponse struct {
	User models.User `json:"user"`
}

type CreateUserRequest struct {
	PhoneNumber string `json:"phone_number"  binding:"required"`
	IdToken     string `json:"id_token" binding:"required"`
}

type UpdateFCMTokenRequest struct {
	Token  string    `json:"token" binding:"required"`
	UserID uuid.UUID `json:"user_id"`
}

type UpdateFCMTokenResponse struct {
	Success bool `json:"success"`
}

type UpdateUserInforRequest struct {
	UserID   uuid.UUID `json:"user_id"`
	FullName string    `json:"full_name"`
	Avatar   string    `json:"avatar"`
	Email    string    `json:"email"`
}

type UpdateUserInforResponse struct {
	User models.User `json:"user"`
}
