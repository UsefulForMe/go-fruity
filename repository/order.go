package repository

import (
	"time"

	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Save(order models.Order) (*models.Order, *errs.AppError)
	FindByUserID(userID uuid.UUID) ([]models.Order, *errs.AppError)

	FindByID(orderID uuid.UUID) (*models.Order, *errs.AppError)
}

type DefaultOrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) DefaultOrderRepository {
	return DefaultOrderRepository{
		db: db,
	}
}

func (repo DefaultOrderRepository) Save(order models.Order) (*models.Order, *errs.AppError) {
	tx := repo.db.Begin()
	if err := tx.Create(&order).Error; err != nil {
		logger.Error("Error while creating an order " + err.Error())
		tx.Rollback()
		return nil, errs.NewUnexpectedError("Unexpected error while creating an order " + err.Error())
	}
	track := models.Track{
		OrderID: order.ID,
		Note:    "",
		Time:    time.Now(),
	}
	if err := tx.Create(&track).Error; err != nil {
		logger.Error("Error while creating an track " + err.Error())
		tx.Rollback()
		return nil, errs.NewUnexpectedError("Unexpected error while creating an track " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		logger.Error("Error while commiting an order " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error while commiting an order " + err.Error())
	}

	return &order, nil
}

func (repo DefaultOrderRepository) FindByUserID(userID uuid.UUID) ([]models.Order, *errs.AppError) {
	var orders []models.Order
	if err := repo.db.Model(&orders).Where("user_id = ?", userID).Preload("OrderItems.Product").Preload("Payment").Preload("Seller").Preload("User").Preload("Tracks").Find(&orders).Error; err != nil {
		logger.Error("Error while finding orders by user id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error while finding orders by user id " + err.Error())
	}
	return orders, nil
}

func (repo DefaultOrderRepository) FindByID(orderID uuid.UUID) (*models.Order, *errs.AppError) {
	var order models.Order
	if err := repo.db.Model(&order).Where("id = ?", orderID).Preload("OrderItems.Product").Preload("Payment").Preload("Seller").Preload("User").Preload("Tracks").Find(&order).Error; err != nil {
		logger.Error("Error while finding order by id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error while finding order by id " + err.Error())
	}
	return &order, nil
}
