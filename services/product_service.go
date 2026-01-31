package services

import (
	"errors"
	"task-session-1/models"
	"task-session-1/repositories"
)

type ProductService struct {
	productRepo  *repositories.ProductRepository
	categoryRepo *repositories.CategoryRepository
}

func NewProductService(
	productRepo *repositories.ProductRepository,
	categoryRepo *repositories.CategoryRepository,
) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.productRepo.GetAll()
}

func (s *ProductService) Create(data *models.Product) error {
	if data.CategoryID == 0 {
		return errors.New("category cannot empty!")
	}

	category, err := s.categoryRepo.GetByID(data.CategoryID)
	if err != nil {
		return err
	}

	if category == nil {
		return errors.New("Category not found!")
	}

	return s.productRepo.Create(data)
}
