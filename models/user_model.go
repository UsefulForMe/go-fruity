package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName    string `json:"full_name" gorm:"type:varchar(100)"`
	Email       string `json:"email" gorm:"type:varchar(100);unique_index"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(100)"`
}
