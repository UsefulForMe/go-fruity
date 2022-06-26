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
	Price       string     `json:"price"`
	CategoryID  *uuid.UUID `json:"category_id"`
	SellerID    *uuid.UUID `json:"seller_id"`
	OldPrice    *float32   `json:"old_price"`
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
type GetTopSaleProductsRequest struct {
	Limit int `json:"limit"`
}
type GetTopSaleProductsRespone struct {
	Products []models.Product `json:"products"`
}

type GetProductsSaleOffRequest struct {
	Limit int `json:"limit"`
}
type GetProductsSaleOffResponse struct {
	Products []models.Product `json:"products"`
}
type GetProductsSaleShockRequest struct {
	Limit int `json:"limit"`
}
type GetProductsSaleShockResponse struct {
	Products []models.Product `json:"products"`
}
type GetProductsByIDsRequest struct {
	IDs []string `json:"ids"`
}
type GetProductsByIDsResponse struct {
	Products []models.Product `json:"products"`
}
