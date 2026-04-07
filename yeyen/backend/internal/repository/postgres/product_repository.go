package postgres

import (
	"backend/internal/domain"
	"context"
	"database/sql"
	"errors"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
	query := `INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRowContext(ctx, query, product.Name, product.Price, product.Stock).Scan(&product.ID)
}

func (r *productRepository) GetByID(ctx context.Context, id int) (*domain.Product, error) {
	query := `SELECT id, name, price, stock FROM products WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var product domain.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetAll(ctx context.Context) ([]*domain.Product, error) {
	query := `SELECT id, name, price, stock FROM products`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (r *productRepository) Update(ctx context.Context, product *domain.Product) error {
	query := `UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, product.Name, product.Price, product.Stock, product.ID)
	return err
}

func (r *productRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
