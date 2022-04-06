package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Product struct {
	CommonModelFields
	ImageURL    string         `json:"image_url"`
	ImageURLS   pq.StringArray `json:"image_urls" gorm:"type:varchar(100)[]"`
	Description string         `json:"description"`
	Name        string         `json:"name"`
	Price       float64        `json:"price"`
	OldPrice    *float64       `json:"old_price"`
	CategoryID  *uuid.UUID     `json:"category_id"`
	Unit        *string        `json:"unit"`
	Tags        pq.StringArray `json:"tags" gorm:"type:varchar(100)[]" sql:"default: '{}'"`
	Instruction *string        `json:"instruction"`
	Origin      *string        `json:"origin"`
	Packs       pq.StringArray `json:"packs,omitempty" gorm:"type:varchar(100)[]" sql:"default: '{}'"`
}
