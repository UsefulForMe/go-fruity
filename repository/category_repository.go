package repository

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	List() ([]models.Category, *errs.AppError)
	Create(category *models.Category) (*models.Category, *errs.AppError)
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
	if err := r.db.Model(&categories).Preload("Parent").Find(&categories).Error; err != nil {
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
