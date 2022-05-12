package repository

import (
	"fmt"
	"time"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Save(order models.Order) (*models.Order, *errs.AppError)
	FindAll() ([]models.Order, *errs.AppError)

	FindByUserID(userID uuid.UUID) ([]models.Order, *errs.AppError)

	FindByID(orderID uuid.UUID) (*models.Order, *errs.AppError)

	ChangeOrderStatus(orderID uuid.UUID, status string, note string) (*models.Order, *errs.AppError)
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

	stocks := []models.Stock{}
	productIDs := []uuid.UUID{}
	for _, item := range order.OrderItems {
		productIDs = append(productIDs, item.ProductID)
	}

	if err := tx.Model(&stocks).Where("product_id in  ?", productIDs).Preload("Product").Find(&stocks).Error; err != nil {
		logger.Error("Error while finding stocks " + err.Error())
		tx.Rollback()
		return nil, errs.NewUnexpectedError("Unexpected error while finding stocks " + err.Error())
	}

	stockMap := map[uuid.UUID]models.Stock{}
	for _, stock := range stocks {
		stockMap[stock.ProductID] = stock
	}

	for _, item := range order.OrderItems {
		stock := stockMap[item.ProductID]
		if stock.Quantity < item.Quantity {
			tx.Rollback()
			return nil, errs.NewUnexpectedError(stock.Product.Name + " has only " + fmt.Sprint(stock.Quantity) + " left")
		}
		stock.Quantity = stock.Quantity - item.Quantity
		if err := tx.Save(&stock).Error; err != nil {
			logger.Error("Error while updating stock " + err.Error())
			tx.Rollback()
			return nil, errs.NewUnexpectedError("Unexpected error while updating stock " + err.Error())
		}
	}

	if err := tx.Commit().Error; err != nil {
		logger.Error("Error while commiting an order " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error while commiting an order " + err.Error())
	}

	return &order, nil
}

func (repo DefaultOrderRepository) FindAll() ([]models.Order, *errs.AppError) {
	var orders []models.Order
	if err := repo.db.Model(&orders).Order("created_at DESC").Preload("Seller").Preload("OrderItems.Product").Preload("OrderItems.Product.Seller").Preload("Payment").Preload("User").Preload("Tracks").Preload("UserAddress").Find(&orders).Error; err != nil {
		logger.Error("Error while finding all orders " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error while finding all orders " + err.Error())
	}

	// get total price
	for index := range orders {
		order := &orders[index]
		var totalPrice float32
		for _, orderItem := range order.OrderItems {
			totalPrice += float32(orderItem.Product.Price) * float32(orderItem.Quantity)
		}
		order.TotalPrice = totalPrice + float32(order.ShippingFee)
	}

	return orders, nil

}

func (repo DefaultOrderRepository) FindByUserID(userID uuid.UUID) ([]models.Order, *errs.AppError) {
	var orders []models.Order
	if err := repo.db.Model(&orders).Where("user_id = ?", userID).Order("created_at DESC").Preload("Seller").Preload("OrderItems.Product").Preload("OrderItems.Product.Seller").Preload("Payment").Preload("User").Preload("Tracks").Preload("UserAddress").Find(&orders).Error; err != nil {
		logger.Error("Error while finding orders by user id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error while finding orders by user id " + err.Error())
	}

	// get total price
	for index := range orders {
		order := &orders[index]
		var totalPrice float32
		for _, orderItem := range order.OrderItems {
			totalPrice += float32(orderItem.Product.Price) * float32(orderItem.Quantity)
		}
		order.TotalPrice = totalPrice + float32(order.ShippingFee)
	}

	return orders, nil
}

func (repo DefaultOrderRepository) FindByID(orderID uuid.UUID) (*models.Order, *errs.AppError) {
	var order models.Order
	if err := repo.db.Model(&order).Where("id = ?", orderID).Preload("Seller").Preload("OrderItems.Product").Preload("OrderItems.Product.Seller").Preload("Payment").Preload("User").Preload("Tracks").Preload("UserAddress").Find(&order).Error; err != nil {
		logger.Error("Error while finding order by id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error while finding order by id " + err.Error())
	}

	totalPrice := float32(0)
	for _, item := range order.OrderItems {
		totalPrice += float32(item.Product.Price) * float32(item.Quantity)
	}
	order.TotalPrice = totalPrice + float32(order.ShippingFee)
	return &order, nil
}

func (repo DefaultOrderRepository) ChangeOrderStatus(orderID uuid.UUID, status string, note string) (*models.Order, *errs.AppError) {
	tx := repo.db.Begin()

	var order models.Order
	if err := tx.Model(&order).Where("id = ?", orderID).Find(&order).Error; err != nil {
		logger.Error("Error while finding order by id " + err.Error())
		tx.Rollback()
		return nil, errs.NewUnexpectedError("Unexpected error while finding order by id " + err.Error())
	}

	if order.Status == status {
		return &order, nil
	}

	track := models.Track{
		OrderID: orderID,
		Note:    note,
		Time:    time.Now(),
		Status:  status,
	}
	if err := tx.Create(&track).Error; err != nil {
		logger.Error("Error while creating an track " + err.Error())
		tx.Rollback()
		return nil, errs.NewUnexpectedError("Unexpected error while creating an track " + err.Error())
	}

	if err := tx.Model(&order).Where("id = ?", orderID).Update("status", status).Error; err != nil {
		logger.Error("Error while updating order status " + err.Error())
		tx.Rollback()
		return nil, errs.NewUnexpectedError("Unexpected error while updating order status " + err.Error())
	}

	if status == dto.OrderStatusCancelled {
		stocks := []models.Stock{}
		productIDs := []uuid.UUID{}

		var items []models.OrderItem

		if err := tx.Model(&items).Where("order_id = ?", orderID).Find(&items).Error; err != nil {
			logger.Error("Error while finding order items by order id " + err.Error())
			tx.Rollback()
			return nil, errs.NewUnexpectedError("Unexpected error while finding order items by order id " + err.Error())
		}

		for _, item := range items {
			productIDs = append(productIDs, item.ProductID)
		}
		if err := tx.Model(&stocks).Where("product_id IN (?)", productIDs).Find(&stocks).Error; err != nil {
			logger.Error("Error while finding stocks by product ids " + err.Error())
			tx.Rollback()
			return nil, errs.NewUnexpectedError("Unexpected error while finding stocks by product ids " + err.Error())
		}
		stockMap := map[uuid.UUID]*models.Stock{}
		for _, stock := range stocks {
			stockMap[stock.ProductID] = &stock
		}

		for _, item := range items {
			stock, ok := stockMap[item.ProductID]
			if !ok {
				logger.Error("Error while finding stock by product id " + item.ProductID.String())
				tx.Rollback()
				return nil, errs.NewUnexpectedError("Unexpected error while finding stock by product id " + item.ProductID.String())
			}

			stock.Quantity += item.Quantity
			if err := tx.Save(stock).Error; err != nil {
				logger.Error("Error while saving stock " + err.Error())
				tx.Rollback()
				return nil, errs.NewUnexpectedError("Unexpected error while saving stock " + err.Error())
			}
		}

	}

	if err := tx.Commit().Error; err != nil {
		logger.Error("Error while commiting an order " + err.Error())
		tx.Rollback()
		return nil, errs.NewUnexpectedError("Unexpected error while commiting an order " + err.Error())
	}

	return &order, nil
}
