package models

import "github.com/google/uuid"

type StockReport struct {
	CommonModelFields
	ProductID uuid.UUID `json:"product_id" gorm:"type:uuid"`
	Product   Product   `json:"product"`
	Content   string    `json:"content" gorm:"type:text"`
	Quantity  int       `json:"quantity" gorm:"type:int;default:0"`
}
