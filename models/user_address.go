package models

import (
	"github.com/google/uuid"
)

type UserAddress struct {
	CommonModelFields
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	IsDefault   bool      `json:"is_default" gorm:"default:false"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(100);not null"`
	FullName    string    `json:"full_name" gorm:"type:varchar(100)"`
	Address     string    `json:"address" gorm:"type:varchar(100); not null"`
	Longitude   float32   `json:"longitude" gorm:"type:numeric; not null"`
	Latitude    float32   `json:"latitude" gorm:"type:numeric; not null"`
	Note        string    `json:"note" gorm:"type:varchar(255);"`
}
