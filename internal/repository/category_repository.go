package repository

import (
	"database/sql"
	"product_service/internal/domain"
)

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(database *sql.DB) domain.CategoryRepository {
	return &categoryRepository{db: database}
}

func (r *categoryRepository) GetAll() ([]*domain.Category, error) {
	query := `SELECT id, name, created_at, updated_at FROM category`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*domain.Category
	for rows.Next() {
		var category domain.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) GetByID(id int) (*domain.Category, error) {
	query := `SELECT id, name, created_at, updated_at FROM category WHERE id = $1`
	var category domain.Category
	err := r.db.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Возвращаем nil, nil, если категория не найдена
		}
		return nil, err
	}
	return &category, nil
}
