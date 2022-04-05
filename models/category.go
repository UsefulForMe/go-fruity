package models

import "github.com/google/uuid"

type Category struct {
	CommonModelFields
	Name     string     `json:"name"`
	ImageURL string     `json:"image_url"`
	Parent   *Category  `json:"parent,omitempty"`
	ParentID *uuid.UUID `json:"parent_id,omitempty"`
	Products []Product  `json:"products,omitempty" gorm:"foreignKey:CategoryID" `
}
