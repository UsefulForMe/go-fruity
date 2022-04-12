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
	Price       float32        `json:"price" gorm:"type:numeric"`
	OldPrice    *float32       `json:"old_price"  gorm:"type:numeric"`
	CategoryID  *uuid.UUID     `json:"category_id"`
	Unit        *string        `json:"unit"`
	Tags        pq.StringArray `json:"tags" gorm:"type:varchar(100)[]" sql:"default: '{}'"`
	Instruction *string        `json:"instruction"`
	Origin      *string        `json:"origin"`
	Packs       pq.StringArray `json:"packs,omitempty" gorm:"type:varchar(100)[]" sql:"default: '{}'"`
	Percent     *float32       `json:"percent,omitempty" gorm:"<-:false;->;-:migration "`
	SellerID    *uuid.UUID     `json:"seller_id"`
	Seller      *Seller        `json:"seller"`
}
