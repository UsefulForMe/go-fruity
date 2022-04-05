package dto

import (
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
)

type CreateProductRequest struct {
	ImageURL    string     `json:"image_url"`
	ImageURLS   []string   `json:"image_urls"`
	Description string     `json:"description"`
	Name        string     `json:"name"`
	Price       float64    `json:"price"`
	CategoryID  *uuid.UUID `json:"category_id"`
	OldPrice    *float64   `json:"old_price"`
	Unit        *string    `json:"unit"`
	Tags        []string   `json:"tags"`
	Instruction *string    `json:"instruction"`
	Origin      *string    `json:"origin"`
	Packs       []string   `json:"packs"`
}

type CreateProductResponse struct {
	Product models.Product `json:"product"`
}

type GetProductsRequest struct {
}
type GetProductsResponse struct {
	Products []models.Product `json:"products"`
}

type GetProductRequest struct {
	ID uuid.UUID `json:"id"`
}
type GetProductResponse struct {
	Product models.Product `json:"product"`
}
