package models

import "github.com/google/uuid"

type StockReport struct {
	CommonModelFields
	StockItems []StockItem `json:"stock_items" gorm:"foreignkey:StockReportID"`
	Content    string      `json:"content" gorm:"type:text"`
}

type StockItem struct {
	CommonModelFields
	StockReportID uuid.UUID `json:"order_id" gorm:"type:uuid"`
	ProductID     uuid.UUID `json:"product_id" gorm:"type:uuid"`
	Product       Product   `json:"product"`
	Quantity      int       `json:"quantity" gorm:"type:int"`
	Note          string    `json:"note" gorm:"type:varchar(100)"`
}
