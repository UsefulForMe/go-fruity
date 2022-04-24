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

	GetUserAddressByID(req dto.GetUserAddressByIDRequest) (*dto.GetUserAddressByIDResponse, *errs.AppError)
	UpdateUserAddress(req dto.UpdateUserAddressRequest) (*dto.UpdateUserAddressResponse, *errs.AppError)
	DeleteUserAddress(req dto.DeleteUserAddressRequest) *errs.AppError
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

func (s DefaultUserAddressService) GetUserAddressByID(req dto.GetUserAddressByIDRequest) (*dto.GetUserAddressByIDResponse, *errs.AppError) {
	userAddress, err := s.userAddressRepository.FindUserAddressByID(req.UserAddressID)
	if err != nil {
		return nil, err
	}
	return &dto.GetUserAddressByIDResponse{UserAddress: *userAddress}, nil
}

func (s DefaultUserAddressService) UpdateUserAddress(req dto.UpdateUserAddressRequest) (*dto.UpdateUserAddressResponse, *errs.AppError) {
	userAddress, err := s.userAddressRepository.FindUserAddressByID(req.UserAddressID)
	if err != nil {
		return nil, err
	}
	userAddress.IsDefault = req.IsDefault
	userAddress.PhoneNumber = req.PhoneNumber
	userAddress.FullName = req.FullName
	userAddress.Address = req.Address
	userAddress.Longitude = req.Longitude
	userAddress.Latitude = req.Latitude
	userAddress.Note = req.Note
	userAddress.UserID = req.UserID

	newUserAddress, err := s.userAddressRepository.Update(*userAddress)
	if err != nil {
		return nil, err
	}
	return &dto.UpdateUserAddressResponse{UserAddress: *newUserAddress}, nil
}

func (s DefaultUserAddressService) DeleteUserAddress(req dto.DeleteUserAddressRequest) *errs.AppError {
	err := s.userAddressRepository.Delete(req.UserAddressID)
	if err != nil {
		return err
	}
	return nil
}
