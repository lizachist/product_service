package domain

import "time"

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryRepository interface {
	GetAll() ([]*Category, error)
	GetByID(id int) (*Category, error)
}

type CategoryService interface {
	GetAll() ([]*Category, error)
	GetByID(id int) (*Category, error)
}
