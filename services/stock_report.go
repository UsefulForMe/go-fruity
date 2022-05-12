package services

import (
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
)

type StockReportService interface {
	GetAllStockReport(req dto.GetAllStockReportRequest) (*dto.GetAllStockReportResponse, *errs.AppError)
	SaveStockReport(req dto.SaveStockReportRequest) (*dto.SaveStockReportResponse, *errs.AppError)
	UpdateStockReport(req dto.UpdateStockReportRequest) (*dto.UpdateStockReportResponse, *errs.AppError)
	// GetStockById(req dto.FindStockByIdRequest) (*dto.FindStockByIdResponse, *errs.AppError)
}

type stockReportService struct {
	stockReportRepository repository.StockReportRepository
}

func NewStockReportService(stockReportRepository repository.StockReportRepository) stockReportService {
	return stockReportService{
		stockReportRepository: stockReportRepository,
	}
}

func (s stockReportService) GetAllStockReport(req dto.GetAllStockReportRequest) (*dto.GetAllStockReportResponse, *errs.AppError) {
	stockreports, err := s.stockReportRepository.Find()
	if err != nil {
		return nil, err
	}
	return &dto.GetAllStockReportResponse{
		StockReports: stockreports}, nil
}

func (s stockReportService) SaveStockReport(req dto.SaveStockReportRequest) (*dto.SaveStockReportResponse, *errs.AppError) {
	stockreport := models.StockReport{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Content:   req.Content,
	}
	newStockReport, err := s.stockReportRepository.Save(&stockreport)
	if err != nil {
		return nil, err
	}
	return &dto.SaveStockReportResponse{
		StockReport: *newStockReport}, nil
}

func (s stockReportService) UpdateStockReport(req dto.UpdateStockReportRequest) (*dto.UpdateStockReportResponse, *errs.AppError) {
	stockReport, err := s.stockReportRepository.FindByID(req.StockReportID)
	if err != nil {
		return nil, err
	}
	stockReport.Content = req.Content
	stockReport.Quantity = req.Quantity
	updatedStockReport, err := s.stockReportRepository.Update(stockReport)
	if err != nil {
		return nil, err
	}
	return &dto.UpdateStockReportResponse{
		StockReport: *updatedStockReport}, nil
}

// func (s stockReportService) GetStockReportById(req dto.FindStockReportByIdRequest) (*dto.FindStockByIdResponse, *errs.AppError) {
// 	stockReport, err := s.stockReportRepository.FindByID(req.StockID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &dto.FindStockByIdResponse{
// 		Stock: *stock}, nil
// }

// func (s stockService) GetStockByProductId(req dto.FindStockByProductIdRequest) (*dto.FindStockByProductIdResponse, *errs.AppError) {
// 	stock, err := s.stockRepository.FindByProductID(req.ProductID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &dto.FindStockByProductIdResponse{
// 		Stock: *stock}, nil
// }
