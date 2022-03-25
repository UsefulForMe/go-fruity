package models

import (
	"github.com/google/uuid"
)

type UserAddress struct {
	CommonModelFields
	UserID      uuid.UUID `json:"user_id" gorm:"primary_key;type:uuid;not null"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(100);unique_index;not null"`
	FullName    string    `json:"full_name" gorm:"type:varchar(100)"`
	Address     string    `json:"address" gorm:"type:varchar(100)"`
}
