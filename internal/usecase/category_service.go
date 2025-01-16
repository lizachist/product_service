package usecase

import (
	"product_service/internal/domain"
)

type CategoryService struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryService(categoryRepo domain.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *CategoryService) GetByID(id int) (*domain.Category, error) {
	return s.categoryRepo.GetByID(id)
}

func (s *CategoryService) GetAll() ([]*domain.Category, error) {
	return s.categoryRepo.GetAll()
}
