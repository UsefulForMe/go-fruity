package services

import (
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/google/uuid"
)

type SellerService interface {
	CreateSeller(req dto.CreateSellerRequest) (*dto.CreateSellerResponse, *errs.AppError)
	GetAllSeller(req dto.GetAllSellerRequest) (*dto.GetAllSellerResponse, *errs.AppError)
	GetSellerByID(req dto.GetSellerByIDRequest) (*dto.GetSellerByIDResponse, *errs.AppError)
	GetProductsBySellerID(req dto.GetProductsBySellerIDRequest) (*dto.GetProductsBySellerIDResponse, *errs.AppError)
}

type DefaultSellerService struct {
	repo repository.SellerRepository
}

func NewSellerService(repo repository.SellerRepository) DefaultSellerService {
	return DefaultSellerService{
		repo: repo,
	}
}

func (s DefaultSellerService) CreateSeller(req dto.CreateSellerRequest) (*dto.CreateSellerResponse, *errs.AppError) {
	seller := models.Seller{
		Name:          req.Name,
		Logo:          req.Logo,
		Banner:        req.Banner,
		Type:          req.Type,
		PhoneNumber:   req.PhoneNumber,
		Description:   req.Description,
		HeadQuarter:   req.HeadQuarter,
		Rating:        req.Rating,
		AvailableTime: req.AvailableTime,
		Note:          req.Note,
		Email:         req.Email,
		TotalVote:     req.TotalVote,
	}

	newSeller, err := s.repo.Save(seller)

	if err != nil {
		return nil, err
	}

	return &dto.CreateSellerResponse{
		Seller: *newSeller,
	}, nil
}

func (s DefaultSellerService) GetAllSeller(req dto.GetAllSellerRequest) (*dto.GetAllSellerResponse, *errs.AppError) {

	var sellers []models.Seller
	var err *errs.AppError

	if len(req.IDs) > 0 {
		ids := make([]uuid.UUID, 0)
		for _, id := range req.IDs {
			ids = append(ids, uuid.MustParse(id))
		}

		sellers, err = s.repo.FindByIDs(ids)
	} else {
		sellers, err = s.repo.FindAll()
	}

	if err != nil {
		return nil, err
	}

	return &dto.GetAllSellerResponse{
		Sellers: sellers,
	}, nil
}

func (s DefaultSellerService) GetSellerByID(req dto.GetSellerByIDRequest) (*dto.GetSellerByIDResponse, *errs.AppError) {
	seller, err := s.repo.FindByID(req.ID)

	if err != nil {
		return nil, err
	}

	return &dto.GetSellerByIDResponse{
		Seller: *seller,
	}, nil
}

func (s DefaultSellerService) GetProductsBySellerID(req dto.GetProductsBySellerIDRequest) (*dto.GetProductsBySellerIDResponse, *errs.AppError) {
	products, err := s.repo.ProductBySeller(req.SellerID)

	if err != nil {
		return nil, err
	}

	return &dto.GetProductsBySellerIDResponse{
		Products: products,
	}, nil
}
