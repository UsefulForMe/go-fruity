package services

import (
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
)

type PaymentService interface {
	CreatePayment(req dto.CreatePaymentRequest) (*dto.CreatePaymentResponse, *errs.AppError)
	MyPayments(req dto.MyPaymentsRequest) (*dto.MyPaymentsResponse, *errs.AppError)
}

type DefaultPaymentService struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentService(paymentRepository repository.PaymentRepository) DefaultPaymentService {
	return DefaultPaymentService{
		paymentRepository: paymentRepository,
	}
}

func (s DefaultPaymentService) CreatePayment(req dto.CreatePaymentRequest) (*dto.CreatePaymentResponse, *errs.AppError) {

	payment := models.Payment{
		Name:      req.Name,
		Provider:  req.Provider,
		AccountNo: req.AccountNo,
		UserID:    req.UserID,
		Logo:      req.Logo,
	}

	newPayment, err := s.paymentRepository.Save(payment)
	if err != nil {
		return nil, err
	}
	return &dto.CreatePaymentResponse{
		Payment: *newPayment,
	}, nil

}

func (s DefaultPaymentService) MyPayments(req dto.MyPaymentsRequest) (*dto.MyPaymentsResponse, *errs.AppError) {

	payments, err := s.paymentRepository.FindByUserID(req.UserID)
	if err != nil {
		return nil, err
	}
	return &dto.MyPaymentsResponse{
		Payments: payments,
	}, nil

}
