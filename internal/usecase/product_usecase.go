package usecase

import (
	"errors"

	"github.com/natanaelrusli/go-mux-api/internal/dto/request"
	"github.com/natanaelrusli/go-mux-api/internal/dto/response"
	"github.com/natanaelrusli/go-mux-api/internal/model"
	"github.com/natanaelrusli/go-mux-api/internal/repository"
)

type ProductUsecase interface {
	GetList(page, limit int) (response.ProductResult, error)
	CreateOne(product request.ProductRequestBody) (model.Product, error)
}

type productUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
	}
}

func (u *productUsecase) GetList(page, limit int) (response.ProductResult, error) {
	var res response.ProductResult

	res, err := u.productRepository.GetProducts(page, limit)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *productUsecase) CreateOne(data request.ProductRequestBody) (model.Product, error) {
	product, err := u.productRepository.AddProduct(data)

	if product.Name == "" {
		return model.Product{}, errors.New("name field is required")
	}

	if product.Price == 0 {
		return model.Product{}, errors.New("price is required and have to be greater than 0")
	}

	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}
