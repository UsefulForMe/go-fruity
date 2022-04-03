package services

import (
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/models"
	"github.com/UsefulForMe/go-ecommerce/repository"
	"github.com/google/uuid"
)

type CategoryService interface {
	GetAllCategories() (*dto.CategoryListResponse, *errs.AppError)
	GetAllParentCategories() (*dto.CategoryListResponse, *errs.AppError)
	GetAllChildCategories(parentID uuid.UUID) (*dto.CategoryListResponse, *errs.AppError)

	CreateCategory(req dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, *errs.AppError)
}

type DefaultCategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return DefaultCategoryService{
		repo: repo,
	}
}

func (s DefaultCategoryService) GetAllCategories() (*dto.CategoryListResponse, *errs.AppError) {
	categories, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return &dto.CategoryListResponse{
		Categories: categories,
	}, nil
}

func (s DefaultCategoryService) GetAllParentCategories() (*dto.CategoryListResponse, *errs.AppError) {
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
	return &dto.CategoryListResponse{
		Categories: parentCategories,
	}, nil
}

func (s DefaultCategoryService) GetAllChildCategories(parentID uuid.UUID) (*dto.CategoryListResponse, *errs.AppError) {
	categories, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	var childCategories []models.Category
	for _, category := range categories {
		if category.ParentID != nil && *category.ParentID == parentID {
			childCategories = append(childCategories, category)
		}
	}
	return &dto.CategoryListResponse{
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
