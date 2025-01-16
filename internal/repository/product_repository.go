package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"product_service/internal/domain"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(database *sql.DB) domain.ProductRepository {
	return &productRepository{db: database}
}

func (r *productRepository) Create(product *domain.Product) error {
	query := `
        INSERT INTO products (name, price, quantity, description, category_id, is_active)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(query,
		product.Name,
		product.Price,
		product.Quantity,
		product.Description,
		product.CategoryID,
		product.IsActive,
	).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)

	return err
}

func (r *productRepository) Update(product *domain.Product) error {
	query := `
        UPDATE products
        SET name=$1, price=$2, quantity=$3, description=$4, category_id=$5, is_active=$6, updated_at=NOW()
        WHERE id=$7`

	result, err := r.db.Exec(query,
		product.Name,
		product.Price,
		product.Quantity,
		product.Description,
		product.CategoryID,
		product.IsActive,
		product.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (r *productRepository) GetByID(id int) (*domain.Product, error) {
	product := &domain.Product{}

	query := `
        SELECT id, name, price, quantity, description, category_id, is_active, created_at, updated_at
		FROM products WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Quantity,
		&product.Description,
		&product.CategoryID,
		&product.IsActive,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return product, nil
}

func (r *productRepository) Delete(product *domain.Product) error {
	query := `DELETE FROM products WHERE id = $1`

	result, err := r.db.Exec(query, product.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (r *productRepository) GetAll() ([]*domain.Product, error) {
	query := `
        SELECT id, name, price, quantity, description, category_id, is_active, created_at, updated_at
        FROM products
        ORDER BY id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying products: %w", err)
	}
	defer rows.Close()

	var products []*domain.Product

	for rows.Next() {
		var product domain.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Quantity,
			&product.Description,
			&product.CategoryID,
			&product.IsActive,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning product row: %w", err)
		}

		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after scanning rows: %w", err)
	}

	return products, nil
}
