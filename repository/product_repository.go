package repository

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product models.Product) (*models.Product, *errs.AppError)
	Find() ([]models.Product, *errs.AppError)
	FindByID(id uuid.UUID) (*models.Product, *errs.AppError)
}

type DefaultProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) DefaultProductRepository {
	return DefaultProductRepository{
		db: db,
	}
}

func (r DefaultProductRepository) Save(product models.Product) (*models.Product, *errs.AppError) {
	if err := r.db.Create(&product).Error; err != nil {
		logger.Error("Error when create product " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when create product " + err.Error())
	}
	return &product, nil
}

func (r DefaultProductRepository) Find() ([]models.Product, *errs.AppError) {
	var products []models.Product

	if err := r.db.Find(&products).Error; err != nil {
		logger.Error("Error when find product " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find product " + err.Error())
	}
	return products, nil
}

func (r DefaultProductRepository) FindByID(id uuid.UUID) (*models.Product, *errs.AppError) {
	var product models.Product
	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {

		if err == gorm.ErrRecordNotFound {

			return nil, errs.NewNotFoundError("Product not found")
		}
		logger.Error("Error when find product by id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find product by id " + err.Error())
	}

	return &product, nil
}
