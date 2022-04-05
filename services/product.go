package services

import (
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
)

type ProductService interface {
	CreateProduct(request *dto.CreateProductRequest) (*dto.CreateProductResponse, *errs.AppError)

	GetProducts(req dto.GetProductsRequest) (*dto.GetProductsResponse, *errs.AppError)
	GetProduct(dto.GetProductRequest) (*dto.GetProductResponse, *errs.AppError)
}

type DefaultProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) DefaultProductService {
	return DefaultProductService{
		productRepository: productRepository,
	}
}

func (s DefaultProductService) CreateProduct(request *dto.CreateProductRequest) (*dto.CreateProductResponse, *errs.AppError) {
	product, err := s.productRepository.Save(models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		CategoryID:  request.CategoryID,
		OldPrice:    request.OldPrice,
		Unit:        request.Unit,
		Tags:        request.Tags,
		Instruction: request.Instruction,
		Origin:      request.Origin,
		ImageURL:    request.ImageURL,
		ImageURLS:   request.ImageURLS,
		Packs:       request.Packs,
	})
	if err != nil {
		return nil, err
	}
	return &dto.CreateProductResponse{
		Product: *product,
	}, nil
}

func (s DefaultProductService) GetProducts(req dto.GetProductsRequest) (*dto.GetProductsResponse, *errs.AppError) {
	products, err := s.productRepository.Find()
	if err != nil {
		return nil, err
	}
	return &dto.GetProductsResponse{Products: products}, nil
}

func (s DefaultProductService) GetProduct(req dto.GetProductRequest) (*dto.GetProductResponse, *errs.AppError) {
	product, err := s.productRepository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	return &dto.GetProductResponse{Product: *product}, nil
}
