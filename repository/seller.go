package repository

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SellerRepository interface {
	Save(seller models.Seller) (*models.Seller, *errs.AppError)
	FindAll() ([]models.Seller, *errs.AppError)
	FindByID(id uuid.UUID) (*models.Seller, *errs.AppError)
	FindByIDs(ids []uuid.UUID) ([]models.Seller, *errs.AppError)

	ProductBySeller(sellerID uuid.UUID) ([]models.Product, *errs.AppError)
}

type DefaultSellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) DefaultSellerRepository {
	return DefaultSellerRepository{
		db: db,
	}
}

func (r DefaultSellerRepository) Save(seller models.Seller) (*models.Seller, *errs.AppError) {
	if err := r.db.Create(&seller).Error; err != nil {
		logger.Error("Error when create seller " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when create seller " + err.Error())
	}
	return &seller, nil
}

func (r DefaultSellerRepository) FindAll() ([]models.Seller, *errs.AppError) {
	var sellers []models.Seller

	if err := r.db.Find(&sellers).Error; err != nil {
		logger.Error("Error when find seller " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find seller " + err.Error())
	}
	return sellers, nil
}

func (r DefaultSellerRepository) FindByID(id uuid.UUID) (*models.Seller, *errs.AppError) {
	var seller models.Seller
	if err := r.db.Where("id = ?", id).First(&seller).Error; err != nil {
		logger.Error("Error when find seller " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find seller " + err.Error())
	}
	return &seller, nil
}

func (r DefaultSellerRepository) FindByIDs(ids []uuid.UUID) ([]models.Seller, *errs.AppError) {
	var sellers []models.Seller
	if err := r.db.Where("id IN (?)", ids).Find(&sellers).Error; err != nil {
		logger.Error("Error when find seller " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find seller " + err.Error())
	}
	return sellers, nil
}

func (r DefaultSellerRepository) ProductBySeller(sellerID uuid.UUID) ([]models.Product, *errs.AppError) {
	var products []models.Product
	if err := r.db.Where("seller_id = ?", sellerID).Preload("Seller").Find(&products).Error; err != nil {
		logger.Error("Error when find product by seller " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find product by seller " + err.Error())
	}
	return products, nil
}
