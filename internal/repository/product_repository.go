package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/natanaelrusli/go-mux-api/internal/dto/request"
	"github.com/natanaelrusli/go-mux-api/internal/dto/response"
	"github.com/natanaelrusli/go-mux-api/internal/model"
	"github.com/natanaelrusli/go-mux-api/internal/utils"
)

type ProductRepository interface {
	GetProducts(page int, limit int) (response.ProductResult, error)
	AddProduct(product request.ProductRequestBody) (model.Product, error)
	HandlePagination(table string, limit, page int) *utils.Pagination
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) HandlePagination(table string, limit, page int) *utils.Pagination {
	var (
		tmpl        = utils.Pagination{}
		recordCount int
	)

	sqlTable := fmt.Sprintf("SELECT count(id) FROM %s", table)
	row := r.db.QueryRow(sqlTable)

	row.Scan(&recordCount)

	total := (recordCount / limit)

	remainder := (recordCount % limit)
	if remainder == 0 {
		tmpl.TotalPage = total
	} else {
		tmpl.TotalPage = total + 1
	}

	tmpl.CurrentPage = page
	tmpl.RecordPerPage = limit

	if page <= 0 {
		tmpl.Next = page + 1
	} else if page < tmpl.TotalPage {
		tmpl.Previous = page - 1
		tmpl.Next = page + 1
	} else if page == tmpl.TotalPage {
		tmpl.Previous = page - 1
		tmpl.Next = 0
	}

	return &tmpl
}

func (r *productRepository) GetProducts(page, limit int) (response.ProductResult, error) {
	var (
		record   = model.Product{}
		products = []model.Product{}
		res      = response.ProductResult{}
	)

	offset := limit * (page - 1)
	pagination := *r.HandlePagination("products", limit, page)

	rows, err := r.db.Query("SELECT id, name, price FROM products order by id asc limit $1 offset $2", limit, offset)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&record.ID, &record.Name, &record.Price); err != nil {
			return res, err
		}
		products = append(products, record)
	}

	if len(products) == 0 {
		return res, nil
	}

	if err := rows.Err(); err != nil {
		return res, err
	}

	res = response.ProductResult{
		Products:   products,
		Pagination: pagination,
	}

	return res, nil
}

func (r *productRepository) AddProduct(product request.ProductRequestBody) (model.Product, error) {
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
