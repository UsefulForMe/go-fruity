package models

import "github.com/google/uuid"

type Stock struct {
	CommonModelFields
	ProductID uuid.UUID `json:"product_id" gorm:"type:uuid"`
	Product   Product   `json:"product"`
	Quantity  int       `json:"quantity" gorm:"type:int;default:0"`
}
