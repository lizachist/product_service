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

	err = p.productRepo.Update(existingProduct)
	if err != nil {
		p.logger.Error("Failed to update product", "error", err, "productID", product.ID)
		return err
	}

	p.logger.Info("Product updated successfully", "productID", product.ID)
	return nil
}

func (p *ProductService) Delete(product *domain.Product) error {
	p.logger.Info("Deleting product", "productID", product.ID)

	err := p.productRepo.Delete(product)
	if err != nil {
		p.logger.Error("Failed to delete product", "error", err, "productID", product.ID)
		return err
	}

	p.logger.Info("Product deleted successfully", "productID", product.ID)
	return nil
}

func (p *ProductService) GetByID(id int) (*domain.Product, error) {
	p.logger.Info("Getting product by ID", "productID", id)

	product, err := p.productRepo.GetByID(id)
	if err != nil {
		p.logger.Error("Failed to get product", "error", err, "productID", id)
		return nil, err
	}

	p.logger.Info("Product retrieved successfully", "productID", id)
	return product, nil
}

func (p *ProductService) GetAll() ([]*domain.Product, error) {
	p.logger.Info("Getting all products")

	products, err := p.productRepo.GetAll()
	if err != nil {
		p.logger.Error("Failed to get all products", "error", err)
		return nil, err
	}

	p.logger.Info("All products retrieved successfully", "count", len(products))
	return products, nil
}
