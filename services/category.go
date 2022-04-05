package services

import (
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/google/uuid"
)

type CategoryService interface {
	GetAllCategories() (*dto.ListCategoryResponse, *errs.AppError)
	GetAllParentCategories() (*dto.ListCategoryResponse, *errs.AppError)
	GetAllChildCategories(dto.ListCategoryRequest) (*dto.ListCategoryResponse, *errs.AppError)

	CreateCategory(req dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, *errs.AppError)

	GetProductsByCategory(dto.GetProductsByCategoryRequest) (*dto.GetProductsByCategoryResponse, *errs.AppError)
}

type DefaultCategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return DefaultCategoryService{
		repo: repo,
	}
}

func (s DefaultCategoryService) GetAllCategories() (*dto.ListCategoryResponse, *errs.AppError) {
	categories, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return &dto.ListCategoryResponse{
		Categories: categories,
	}, nil
}

func (s DefaultCategoryService) GetAllParentCategories() (*dto.ListCategoryResponse, *errs.AppError) {
	categories, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	var parentCategories []models.Category
	for _, category := range categories {
		if category.ParentID == nil {
			parentCategories = append(parentCategories, category)
		}
	}
	return &dto.ListCategoryResponse{
		Categories: parentCategories,
	}, nil
}

func (s DefaultCategoryService) GetAllChildCategories(req dto.ListCategoryRequest) (*dto.ListCategoryResponse, *errs.AppError) {
	categories, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	var childCategories []models.Category
	for _, category := range categories {
		if category.ParentID != nil && *category.ParentID == *req.ParentID {
			childCategories = append(childCategories, category)
		}
	}
	return &dto.ListCategoryResponse{
		Categories: childCategories,
	}, nil
}

func (s DefaultCategoryService) CreateCategory(req dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, *errs.AppError) {
	category := models.Category{
		Name:     req.Name,
		ImageURL: req.ImageURL,
	}
	if req.ParentID != nil {
		if id, err := uuid.Parse(*req.ParentID); err != nil {
			return nil, errs.NewBadRequestError("Invalid parent id")
		} else {
			category.ParentID = &id
		}
	}
	createdCategory, err := s.repo.Create(&category)
	if err != nil {
		return nil, err
	}
	return &dto.CreateCategoryResponse{
		Category: *createdCategory,
	}, nil
}

func (s DefaultCategoryService) GetProductsByCategory(req dto.GetProductsByCategoryRequest) (*dto.GetProductsByCategoryResponse, *errs.AppError) {
	products, err := s.repo.ListProducts(req.CategoryID)
	if err != nil {
		return nil, err
	}
	return &dto.GetProductsByCategoryResponse{Products: products}, nil

}
