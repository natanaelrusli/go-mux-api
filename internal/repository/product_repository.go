package repository

import (
	"database/sql"
	"errors"

	"github.com/natanaelrusli/go-mux-api/internal/model"
)

type ProductRepository interface {
	GetProducts(start, count int) ([]model.Product, error)
	AddProduct(product model.Product) (model.Product, error)
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

func (r *productRepository) AddProduct(product model.Product) (model.Product, error) {
	var p model.Product

	row := r.db.QueryRow(
		`INSERT INTO products (name, price) VALUES($1, $2) 
		RETURNING id, name, price`,
		product.Name,
		product.Price,
	)

	if row.Err() != nil {
		switch row.Err() {
		default:
			return model.Product{}, errors.New("internal server error")
		}
	}

	row.Scan(&p.ID, &p.Name, &p.Price)

	return p, nil
}
