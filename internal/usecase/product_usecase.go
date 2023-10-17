package usecase

import (
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

	if err != nil {
		return model.Product{}, err
	}

	return res, nil
}
