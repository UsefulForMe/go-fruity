package repository

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockReportRepository interface {
	Find() ([]models.StockReport, *errs.AppError)
	Save(stockReport *models.StockReport) (*models.StockReport, *errs.AppError)
	Update(stockReport *models.StockReport) (*models.StockReport, *errs.AppError)
	FindByID(id uuid.UUID) (*models.StockReport, *errs.AppError)
}

type DefaultStockReportRepository struct {
	db *gorm.DB
}

func NewStockReportRepository(db *gorm.DB) DefaultStockReportRepository {
	return DefaultStockReportRepository{
		db: db,
	}
}

func (r DefaultStockReportRepository) Find() ([]models.StockReport, *errs.AppError) {
	var stockReports []models.StockReport
	if err := r.db.Model(&stockReports).Preload("Product").Find(&stockReports).Error; err != nil {
		logger.Error("Error when find all stockReports " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find all stockReports " + err.Error())
	}
	return stockReports, nil
}

func (r DefaultStockReportRepository) Save(stockReport *models.StockReport) (*models.StockReport, *errs.AppError) {
	if err := r.db.Create(&stockReport).Error; err != nil {
		logger.Error("Error when create stock " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when create stock " + err.Error())
	}
	return stockReport, nil
}

func (r DefaultStockReportRepository) Update(stockReport *models.StockReport) (*models.StockReport, *errs.AppError) {
	if err := r.db.Save(&stockReport).Error; err != nil {
		logger.Error("Error when update stockReport " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when update stockReport " + err.Error())
	}
	return stockReport, nil
}

func (r DefaultStockReportRepository) FindByID(id uuid.UUID) (*models.StockReport, *errs.AppError) {
	var stockReport models.StockReport
	if err := r.db.Model(&stockReport).Where("id = ?", id).First(&stockReport).Error; err != nil {
		logger.Error("Error when find stockReport by id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find stockReport by id " + err.Error())
	}
	return &stockReport, nil
}
