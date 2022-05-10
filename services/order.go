package services

import (
	"time"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
)

type OrderService interface {
	GetAllOrders(req dto.GetAllOrdersRequest) (*dto.GetAllOrdersResponse, *errs.AppError)
	CreateOrder(req dto.CreateOrderRequest) (*dto.CreateOrderResponse, *errs.AppError)
	MyOrders(req dto.MyOrdersRequest) (*dto.MyOrdersResponse, *errs.AppError)
	GetOrderByID(req dto.GetOrderByIDRequest) (*dto.GetOrderByIDResponse, *errs.AppError)

	ChangeOrderStatus(req dto.ChangeOrderStatusRequest) (*dto.ChangeOrderStatusResponse, *errs.AppError)
}

type DefaultOrderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) DefaultOrderService {
	return DefaultOrderService{
		orderRepo: orderRepo,
	}
}

func (s DefaultOrderService) GetAllOrders(req dto.GetAllOrdersRequest) (*dto.GetAllOrdersResponse, *errs.AppError) {
	orders, err := s.orderRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return &dto.GetAllOrdersResponse{
		Orders: orders}, nil

}

func (s DefaultOrderService) CreateOrder(req dto.CreateOrderRequest) (*dto.CreateOrderResponse, *errs.AppError) {
	order := models.Order{
		UserID:           req.UserID,
		SellerID:         req.SellerID,
		PaymentID:        req.PaymentID,
		OrderItems:       req.OrderItems,
		ReceivedAt:       time.Time(req.ReceivedAt),
		Note:             req.Note,
		UserAddressID:    req.UserAddressID,
		ShippingFee:      req.ShippingFee,
		ShippingDistance: req.ShippingDistance,
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
	filterOrders := []models.Order{}
	if req.Status != "" {
		for _, order := range orders {
			if order.Status == req.Status {
				filterOrders = append(filterOrders, order)
			}
		}
	} else {
		filterOrders = orders
	}

	return &dto.MyOrdersResponse{
		Orders: filterOrders}, nil
}

func (s DefaultOrderService) GetOrderByID(req dto.GetOrderByIDRequest) (*dto.GetOrderByIDResponse, *errs.AppError) {
	order, err := s.orderRepo.FindByID(req.OrderID)
	if err != nil {
		return nil, err
	}
	return &dto.GetOrderByIDResponse{
		Order: *order}, nil
}

func (s DefaultOrderService) ChangeOrderStatus(req dto.ChangeOrderStatusRequest) (*dto.ChangeOrderStatusResponse, *errs.AppError) {
	order, err := s.orderRepo.ChangeOrderStatus(
		req.OrderID,
		req.Status,
		req.Note,
	)
	if err != nil {
		return nil, err
	}

	return &dto.ChangeOrderStatusResponse{
		Order: *order}, nil
}
