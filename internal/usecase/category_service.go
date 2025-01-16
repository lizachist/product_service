package usecase

import (
	"log/slog"
	"product_service/internal/domain"
)

type CategoryService struct {
	categoryRepo domain.CategoryRepository
	logger       *slog.Logger
}

func NewCategoryService(categoryRepo domain.CategoryRepository, logger *slog.Logger) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
		logger:       logger,
	}
}

func (s *CategoryService) GetByID(id int) (*domain.Category, error) {
	s.logger.Info("Getting category by ID", "categoryID", id)

	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		s.logger.Error("Failed to get category", "error", err, "categoryID", id)
		return nil, err
	}
	if category == nil {
		s.logger.Warn("Category not found", "categoryID", id)
		return nil, nil
	}

	s.logger.Info("Category retrieved successfully", "categoryID", id)
	return category, nil
}

func (s *CategoryService) GetAll() ([]*domain.Category, error) {
	s.logger.Info("Getting all categories")

	categories, err := s.categoryRepo.GetAll()
	if err != nil {
		s.logger.Error("Failed to get all categories", "error", err)
		return nil, err
	}

	s.logger.Info("All categories retrieved successfully", "count", len(categories))
	return categories, nil
}
