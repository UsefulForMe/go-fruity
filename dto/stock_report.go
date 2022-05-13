package dto

import (
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
)

type GetAllStockReportRequest struct{}

type GetAllStockReportResponse struct {
	StockReports []models.StockReport `json:"stock_reports"`
}

type SaveStockReportRequest struct {
	Content    string             `json:"content"`
	StockItems []models.StockItem `json:"stock_items"`
}

type SaveStockReportResponse struct {
	StockReport models.StockReport `json:"stock_report"`
}

type UpdateStockReportRequest struct {
	StockReportID uuid.UUID          `json:"stock_report_id"`
	Content       string             `json:"content"`
	StockItems    []models.StockItem `json:"stock_items"`
}

type UpdateStockReportResponse struct {
	StockReport models.StockReport `json:"stock_report"`
}

type FindStockReportByIdRequest struct {
	StockReportID uuid.UUID `json:"stock_report_id"`
}

type FindStockReportByIdResponse struct {
	StockReport models.StockReport `json:"stock_report"`
}
