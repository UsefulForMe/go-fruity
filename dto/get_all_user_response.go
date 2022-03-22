package dto

import "github.com/UsefulForMe/go-ecommerce/models"

type GetAllUserResponse struct {
	Users []models.User `json:"users"`
}
