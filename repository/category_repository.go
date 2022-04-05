package repository

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	List() ([]models.Category, *errs.AppError)
	Create(category *models.Category) (*models.Category, *errs.AppError)

	ListProducts(id uuid.UUID) ([]models.Product, *errs.AppError)
}

type DefaultCatoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return DefaultCatoryRepository{
		db: db,
	}
}

func (r DefaultCatoryRepository) List() ([]models.Category, *errs.AppError) {
	var categories []models.Category
	if err := r.db.Model(&categories).Preload("Parent").Order("name ASC").Find(&categories).Error; err != nil {
		logger.Error("Error when find all categories " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find all categories " + err.Error())
	}
	return categories, nil
}

func (r DefaultCatoryRepository) Create(category *models.Category) (*models.Category, *errs.AppError) {
	if err := r.db.Create(&category).Error; err != nil {
		logger.Error("Error when create category " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when create category " + err.Error())
	}
	return category, nil
}

func (r DefaultCatoryRepository) ListProducts(id uuid.UUID) ([]models.Product, *errs.AppError) {
	var products []models.Product
	if err := r.db.Model(&products).Where("category_id = ?", id).Find(&products).Error; err != nil {
		logger.Error("Error when find all products " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error when find all products " + err.Error())
	}
	return products, nil
}
