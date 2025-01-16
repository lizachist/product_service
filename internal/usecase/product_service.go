package usecase

import (
	"log/slog"
	"product_service/internal/domain"
)

type ProductService struct {
	productRepo  domain.ProductRepository
	categoryRepo domain.CategoryRepository
	logger       *slog.Logger
}

func NewProductService(
	productRepo domain.ProductRepository,
	categoryRepo domain.CategoryRepository, logger *slog.Logger) domain.ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		logger:       logger,
	}
}

func (p *ProductService) Create(product *domain.Product) error {
	p.logger.Info("Creating new product", "name", product.Name, "categoryID", *product.CategoryID)
	_, err := p.categoryRepo.GetByID(*product.CategoryID)
	if err != nil {
		p.logger.Error("Invalid category ID", "error", err, "categoryID", *product.CategoryID)
		return domain.ErrInvalidCategoryID
	}
	err = p.productRepo.Create(product)
	if err != nil {
		p.logger.Error("Failed to create product", "error", err)
		return err
	}
	p.logger.Info("Product created successfully", "productID", product.ID)
	return nil
}

func (p *ProductService) Update(product *domain.Product) error {
	p.logger.Info("Updating product", "productID", product.ID)
	existingProduct, err := p.productRepo.GetByID(product.ID)
	if err != nil {
		p.logger.Error("Failed to get existing product", "error", err, "productID", product.ID)
		return err
	}

	if product.Name != "" {
		existingProduct.Name = product.Name
	}
	if product.Price != 0 {
		existingProduct.Price = product.Price
	}
	if product.Description != nil {
		existingProduct.Description = product.Description
	}
	existingProduct.IsActive = product.IsActive

	return p.productRepo.Update(existingProduct)
}

func (p *ProductService) Delete(product *domain.Product) error {
	return p.productRepo.Delete(product)
}

func (p *ProductService) GetByID(id int) (*domain.Product, error) {
	return p.productRepo.GetByID(id)
}

func (p *ProductService) GetAll() ([]*domain.Product, error) {
	return p.productRepo.GetAll()
}
