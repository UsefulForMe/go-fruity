package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `json:"user_name" gorm:"type:varchar(100);unique_index"`
	Password    string `json:"password" gorm:"type:varchar(100)"`
	FirstName   string `json:"first_name" gorm:"type:nvarchar(100)"`
	LastName    string `json:"last_name" gorm:"type:nvarchar(100)"`
	Email       string `json:"email" gorm:"type:varchar(100);unique_index"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(100)"`
}
