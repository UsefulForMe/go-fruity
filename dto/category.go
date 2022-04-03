package dto

import "github.com/UsefulForMe/go-ecommerce/models"

type CategoryListResponse struct {
	Categories []models.Category `json:"categories"`
}

type CategoryListRequest struct {
	ParentID *string `json:"parent_id"`
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
