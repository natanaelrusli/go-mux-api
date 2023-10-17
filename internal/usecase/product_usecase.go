package usecase

import (
	"errors"

	"github.com/natanaelrusli/go-mux-api/internal/model"
	"github.com/natanaelrusli/go-mux-api/internal/repository"
)

type ProductUsecase interface {
	GetList() ([]model.Product, error)
	CreateOne(product model.Product) (model.Product, error)
}

type productUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
	}
}

func (u *productUsecase) GetList() ([]model.Product, error) {
	res, err := u.productRepository.GetProducts(1, 5)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *productUsecase) CreateOne(product model.Product) (model.Product, error) {
	res, err := u.productRepository.AddProduct(product)

	if product.Name == "" {
		return model.Product{}, errors.New("name field is required")
	} else if product.Price == 0 {
		return model.Product{}, errors.New("price is required and have to be greater than 0")
	}

	if err != nil {
		return model.Product{}, err
	}

	return res, nil
}
