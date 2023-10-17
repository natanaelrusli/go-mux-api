package repository

import (
	"database/sql"

	"github.com/natanaelrusli/go-mux-api/internal/model"
)

type ProductRepository interface {
	GetProducts(start, count int) ([]model.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetProducts(start, count int) ([]model.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price FROM products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if len(products) == 0 {
		return []model.Product{}, nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
