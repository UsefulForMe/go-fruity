package dto

import (
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
)

type CreateUserAddressRequest struct {
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	IsDefault   bool      `json:"is_default" gorm:"default:false"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(100);not null"`
	FullName    string    `json:"full_name" gorm:"type:varchar(100)"`
	Address     string    `json:"address" gorm:"type:varchar(100); not null"`
	Longitude   float32   `json:"longitude" gorm:"type:numeric; not null"`
	Latitude    float32   `json:"latitude" gorm:"type:numeric; not null"`
	Note        string    `json:"note" gorm:"type:varchar(255);"`
	Type        string    `json:"type" gorm:"type:varchar(100);"`
}

type CreateUserAddressResponse struct {
	UserAddress models.UserAddress `json:"user_address"`
}

type MyAddressesRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type MyAddressesResponse struct {
	UserAddresses []models.UserAddress `json:"user_addresses"`
}
