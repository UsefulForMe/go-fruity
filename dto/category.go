package dto

import (
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
)

type ListCategoryResponse struct {
	Categories []models.Category `json:"categories"`
}
type ListCategoryRequest struct {
	ParentID *uuid.UUID `json:"parent_id"`
}

//---------------------------------------------
type CreateCategoryRequest struct {
	Name     string  `json:"name"`
	ImageURL string  `json:"image_url"`
	ParentID *string `json:"parent_id"`
}
type CreateCategoryResponse struct {
	Category models.Category `json:"category"`
}

//---------------------------------------------

type GetProductsByCategoryRequest struct {
	CategoryID uuid.UUID `json:"category_id"`
}

type GetProductsByCategoryResponse struct {
	Products []models.Product `json:"products"`
}
