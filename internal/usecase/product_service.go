package usecase

import "product_service/internal/domain"

type ProductService struct {
	productRepo  domain.ProductRepository
	categoryRepo domain.CategoryRepository
}

func NewProductService(
	productRepo domain.ProductRepository,
	categoryRepo domain.CategoryRepository) domain.ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (p *ProductService) Create(product *domain.Product) error {
	_, err := p.categoryRepo.GetByID(product.CategoryID)
	if err != nil {
		return domain.ErrInvalidCategoryID
	}
	return p.productRepo.Create(product)
}

func (p *ProductService) Update(product *domain.Product) error {
	existingProduct, err := p.productRepo.GetByID(product.ID)
	if err != nil {
		return err
	}

	if product.Name != "" {
		existingProduct.Name = product.Name
	}
	if product.Price != 0 {
		existingProduct.Price = product.Price
	}
	if product.Description != "" {
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
