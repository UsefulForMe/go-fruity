package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(request *dto.CreateProductRequest) (*dto.CreateProductResponse, *errs.AppError)
	GetTopSaleProducts(request dto.GetTopSaleProductsRequest) (*dto.GetTopSaleProductsRespone, *errs.AppError)
	GetSaleOffProducts(request dto.GetProductsSaleOffRequest) (*dto.GetProductsSaleOffResponse, *errs.AppError)
	GetSaleShockProducts(request dto.GetProductsSaleShockRequest) (*dto.GetProductsSaleShockResponse, *errs.AppError)
	GetProducts(req dto.GetProductsRequest) (*dto.GetProductsResponse, *errs.AppError)
	GetProduct(dto.GetProductRequest) (*dto.GetProductResponse, *errs.AppError)
	GetProductsByIDS(req dto.GetProductsByIDsRequest) (*dto.GetProductsByIDsResponse, *errs.AppError)
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
	var price float32
	if s, err := strconv.ParseFloat(request.Price, 32); err == nil {
		price = float32(s)
	} else {
		return nil, errs.NewBadRequestError("Invalid price")
	}
	product, err := s.productRepository.Save(models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       price,
		CategoryID:  request.CategoryID,
		SellerID:    request.SellerID,
		OldPrice:    request.OldPrice,
		Unit:        request.Unit,
		Tags:        request.Tags,
		Instruction: request.Instruction,
		Origin:      request.Origin,
		ImageURL:    request.ImageURL,
		ImageURLS:   request.ImageURLS,
		Packs:       request.Packs,
	})
	postBody, _ := json.Marshal(map[string]string{
		"id":       product.ID.String(),
		"name":     product.Name,
		"price":    strconv.FormatFloat(float64(product.Price), 'f', 2, 32),
		"unit":     *product.Unit,
		"imageUrl": product.ImageURL,
		"origin":   *product.Origin,
	})
	// send this to index product on elastic search
	go func() {
		// set header

		req, _ := http.NewRequest("POST", "https://fruity.es.us-central1.gcp.cloud.es.io/product/_doc", bytes.NewBuffer(postBody))
		req.Header.Set("Authorization", "Basic ZWxhc3RpYzpWMTM4ZG9nN3RPN0JrTW9mN2hQZktKTVA=")
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		res, _ := client.Do(req)
		fmt.Println(res)
	}()

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
func (s DefaultProductService) GetTopSaleProducts(request dto.GetTopSaleProductsRequest) (*dto.GetTopSaleProductsRespone, *errs.AppError) {
	products, err := s.productRepository.FindTopSales(request.Limit)
	if err != nil {
		return nil, err
	}
	return &dto.GetTopSaleProductsRespone{Products: products}, nil
}
func (s DefaultProductService) GetSaleOffProducts(req dto.GetProductsSaleOffRequest) (*dto.GetProductsSaleOffResponse, *errs.AppError) {
	products, err := s.productRepository.FindSaleOff(req.Limit)
	if err != nil {
		return nil, err
	}
	return &dto.GetProductsSaleOffResponse{Products: products}, nil
}
func (s DefaultProductService) GetSaleShockProducts(req dto.GetProductsSaleShockRequest) (*dto.GetProductsSaleShockResponse, *errs.AppError) {
	products, err := s.productRepository.FindSaleShock(req.Limit)
	if err != nil {
		return nil, err
	}
	return &dto.GetProductsSaleShockResponse{Products: products}, nil
}
func (s DefaultProductService) GetProductsByIDS(req dto.GetProductsByIDsRequest) (*dto.GetProductsByIDsResponse, *errs.AppError) {
	ids := make([]uuid.UUID, 0)

	for _, id := range req.IDs {
		ids = append(ids, uuid.MustParse(id))
	}
	products, err := s.productRepository.FindByIDs(ids)
	if err != nil {
		return nil, err
	}
	return &dto.GetProductsByIDsResponse{Products: products}, nil
}
