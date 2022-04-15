package repository

import (
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Save(payment models.Payment) (*models.Payment, *errs.AppError)
	FindByUserID(userID uuid.UUID) ([]models.Payment, *errs.AppError)
}

type DefaultPaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) DefaultPaymentRepository {
	return DefaultPaymentRepository{
		db: db,
	}
}

func (repo DefaultPaymentRepository) Save(payment models.Payment) (*models.Payment, *errs.AppError) {
	if err := repo.db.Create(&payment).Error; err != nil {
		logger.Error("Error while creating an payment " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error while creating a payment " + err.Error())
	}
	return &payment, nil
}

func (repo DefaultPaymentRepository) FindByUserID(userID uuid.UUID) ([]models.Payment, *errs.AppError) {
	var payments []models.Payment
	if err := repo.db.Where("user_id = ?", userID).Find(&payments).Error; err != nil {
		logger.Error("Error while finding payments by user id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error while finding payments by user id " + err.Error())
	}
	return payments, nil
}
