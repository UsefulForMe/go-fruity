package repository

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockRepository interface {
	Find() ([]models.Stock, *errs.AppError)
	Save(stock *models.Stock) (*models.Stock, *errs.AppError)
	Update(stock *models.Stock) (*models.Stock, *errs.AppError)

	FindByProductID(id uuid.UUID) (*models.Stock, *errs.AppError)
	FindByID(id uuid.UUID) (*models.Stock, *errs.AppError)
}

type DefaultStockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) DefaultStockRepository {
	return DefaultStockRepository{
		db: db,
	}
}

func (r DefaultStockRepository) Find() ([]models.Stock, *errs.AppError) {
	var stocks []models.Stock
	if err := r.db.Model(&stocks).Preload("Product").Find(&stocks).Error; err != nil {
		logger.Error("Error when find all stocks " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find all stocks " + err.Error())
	}
	return stocks, nil
}

func (r DefaultStockRepository) Save(stock *models.Stock) (*models.Stock, *errs.AppError) {
	if err := r.db.Create(&stock).Error; err != nil {
		logger.Error("Error when create stock " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when create stock " + err.Error())
	}
	return stock, nil
}

func (r DefaultStockRepository) Update(stock *models.Stock) (*models.Stock, *errs.AppError) {
	if err := r.db.Save(&stock).Error; err != nil {
		logger.Error("Error when update stock " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when update stock " + err.Error())
	}
	return stock, nil
}

func (r DefaultStockRepository) FindByProductID(id uuid.UUID) (*models.Stock, *errs.AppError) {
	var stock models.Stock
	if err := r.db.Model(&stock).Where("product_id = ?", id).First(&stock).Error; err != nil {
		logger.Error("Error when find stock by product id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find stock by product id " + err.Error())
	}
	return &stock, nil
}

func (r DefaultStockRepository) FindByID(id uuid.UUID) (*models.Stock, *errs.AppError) {
	var stock models.Stock
	if err := r.db.Model(&stock).Where("id = ?", id).First(&stock).Error; err != nil {
		logger.Error("Error when find stock by id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find stock by id " + err.Error())
	}
	return &stock, nil
}
