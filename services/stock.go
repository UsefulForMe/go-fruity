package services

import (
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
)

type StockService interface {
	GetAllStock(req dto.GetAllStockRequest) (*dto.GetAllStockResponse, *errs.AppError)
	SaveStock(req dto.SaveStockRequest) (*dto.SaveStockResponse, *errs.AppError)
	UpdateStock(req dto.UpdateStockRequest) (*dto.UpdateStockResponse, *errs.AppError)
	GetStockById(req dto.FindStockByIdRequest) (*dto.FindStockByIdResponse, *errs.AppError)
	GetStockByProductId(req dto.FindStockByProductIdRequest) (*dto.FindStockByProductIdResponse, *errs.AppError)
}

type stockService struct {
	stockRepository repository.StockRepository
}

func NewStockService(stockRepository repository.StockRepository) stockService {
	return stockService{
		stockRepository: stockRepository,
	}
}

func (s stockService) GetAllStock(req dto.GetAllStockRequest) (*dto.GetAllStockResponse, *errs.AppError) {
	stocks, err := s.stockRepository.Find()
	if err != nil {
		return nil, err
	}
	return &dto.GetAllStockResponse{
		Stocks: stocks}, nil
}

func (s stockService) SaveStock(req dto.SaveStockRequest) (*dto.SaveStockResponse, *errs.AppError) {
	stock := models.Stock{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
	newStock, err := s.stockRepository.Save(&stock)
	if err != nil {
		return nil, err
	}
	return &dto.SaveStockResponse{
		Stock: *newStock}, nil
}

func (s stockService) UpdateStock(req dto.UpdateStockRequest) (*dto.UpdateStockResponse, *errs.AppError) {
	stock, err := s.stockRepository.FindByID(req.StockID)
	if err != nil {
		return nil, err
	}
	stock.Quantity = req.Quantity
	updatedStock, err := s.stockRepository.Update(stock)
	if err != nil {
		return nil, err
	}
	return &dto.UpdateStockResponse{
		Stock: *updatedStock}, nil
}

func (s stockService) GetStockById(req dto.FindStockByIdRequest) (*dto.FindStockByIdResponse, *errs.AppError) {
	stock, err := s.stockRepository.FindByID(req.StockID)
	if err != nil {
		return nil, err
	}
	return &dto.FindStockByIdResponse{
		Stock: *stock}, nil
}

func (s stockService) GetStockByProductId(req dto.FindStockByProductIdRequest) (*dto.FindStockByProductIdResponse, *errs.AppError) {
	stock, err := s.stockRepository.FindByProductID(req.ProductID)
	if err != nil {
		return nil, err
	}
	return &dto.FindStockByProductIdResponse{
		Stock: *stock}, nil
}
