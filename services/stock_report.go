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
	GetStockReportById(req dto.FindStockReportByIdRequest) (*dto.FindStockReportByIdResponse, *errs.AppError)
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
		StockItems: req.StockItems,
		Content:    req.Content,
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
	stockReport.StockItems = req.StockItems
	updatedStockReport, err := s.stockReportRepository.Update(stockReport)
	if err != nil {
		return nil, err
	}
	return &dto.UpdateStockReportResponse{
		StockReport: *updatedStockReport}, nil
}

func (s stockReportService) GetStockReportById(req dto.FindStockReportByIdRequest) (*dto.FindStockReportByIdResponse, *errs.AppError) {
	stockReport, err := s.stockReportRepository.FindByID(req.StockReportID)
	if err != nil {
		return nil, err
	}
	return &dto.FindStockReportByIdResponse{
		StockReport: *stockReport}, nil
}
