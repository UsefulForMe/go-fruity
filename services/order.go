package services

import (
	"time"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
)

type OrderService interface {
	CreateOrder(req dto.CreateOrderRequest) (*dto.CreateOrderResponse, *errs.AppError)
	MyOrders(req dto.MyOrdersRequest) (*dto.MyOrdersResponse, *errs.AppError)
}

type DefaultOrderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) DefaultOrderService {
	return DefaultOrderService{
		orderRepo: orderRepo,
	}
}

func (s DefaultOrderService) CreateOrder(req dto.CreateOrderRequest) (*dto.CreateOrderResponse, *errs.AppError) {
	order := models.Order{
		UserID:        req.UserID,
		SellerID:      req.SellerID,
		PaymentID:     req.PaymentID,
		OrderItems:    req.OrderItems,
		ReceivedAt:    time.Time(req.ReceivedAt),
		Note:          req.Note,
		UserAddressID: req.UserAddressID,
	}
	newOrder, err := s.orderRepo.Save(order)
	if err != nil {
		return nil, err
	}

	return &dto.CreateOrderResponse{
		Order: *newOrder}, nil
}

func (s DefaultOrderService) MyOrders(req dto.MyOrdersRequest) (*dto.MyOrdersResponse, *errs.AppError) {
	orders, err := s.orderRepo.FindByUserID(req.UserID)
	if err != nil {
		return nil, err
	}
	return &dto.MyOrdersResponse{
		Orders: orders}, nil
}
