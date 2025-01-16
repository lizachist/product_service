package dto

import (
	"product_service/internal/domain"
	"time"
)

type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Description *string   `json:"description,omitempty"`
	Category    string    `json:"category,omitempty"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewProductResponse(product *domain.Product, category *domain.Category) *ProductResponse {
	if product == nil {
		return nil
	}

	resp := &ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Description: product.Description,
		IsActive:    product.IsActive,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	if category != nil {
		resp.Category = category.Name
	}

	return resp
}
