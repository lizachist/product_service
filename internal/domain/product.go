package domain

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Description *string   `json:"description"`
	CategoryID  *int      `json:"category_id"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductRepository interface {
	Create(product *Product) error
	Update(product *Product) error
	GetByID(id int) (*Product, error)
	Delete(product *Product) error
	GetAll() ([]*Product, error)
}

type ProductService interface {
	Create(product *Product) error
	Update(product *Product) error
	Delete(product *Product) error
	GetByID(id int) (*Product, error)
	GetAll() ([]*Product, error)
}
