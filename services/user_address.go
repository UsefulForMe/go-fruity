package services

import (
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
)

type UserAddressService interface {
	CreateUserAddress(req dto.CreateUserAddressRequest) (*dto.CreateUserAddressResponse, *errs.AppError)

	MyAddresses(req dto.MyAddressesRequest) (*dto.MyAddressesResponse, *errs.AppError)
}

type DefaultUserAddressService struct {
	userAddressRepository repository.UserAddressRepository
}

func NewUserAddressService(userAddressRepository repository.UserAddressRepository) DefaultUserAddressService {
	return DefaultUserAddressService{
		userAddressRepository: userAddressRepository,
	}
}

func (s DefaultUserAddressService) CreateUserAddress(req dto.CreateUserAddressRequest) (*dto.CreateUserAddressResponse, *errs.AppError) {

	userAddress := models.UserAddress{
		UserID:      req.UserID,
		IsDefault:   req.IsDefault,
		PhoneNumber: req.PhoneNumber,
		FullName:    req.FullName,
		Address:     req.Address,
		Longitude:   req.Longitude,
		Latitude:    req.Latitude,
		Note:        req.Note,
		Type:        req.Type,
	}
	newUserAddress, err := s.userAddressRepository.Save(userAddress)
	if err != nil {
		return nil, err
	}
	return &dto.CreateUserAddressResponse{UserAddress: *newUserAddress}, nil
}

func (s DefaultUserAddressService) MyAddresses(req dto.MyAddressesRequest) (*dto.MyAddressesResponse, *errs.AppError) {
	userAddresses, err := s.userAddressRepository.FindUserAddressesByUserID(req.UserID)
	if err != nil {
		return nil, err
	}
	return &dto.MyAddressesResponse{UserAddresses: userAddresses}, nil
}
