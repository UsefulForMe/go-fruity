package models

import "github.com/google/uuid"

type Payment struct {
	CommonModelFields
	Name      string    `json:"name" gorm:"type:varchar(100)"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid"`
	Provider  string    `json:"provider" gorm:"type:varchar(100)"`
	Status    string    `json:"status" gorm:"type:varchar(100);default:active"`
	AccountNo string    `json:"account_no" gorm:"type:varchar(100)"`
	Logo      string    `json:"logo" gorm:"type:varchar(100)"`
}