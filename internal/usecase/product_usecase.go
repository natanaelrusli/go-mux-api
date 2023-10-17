package usecase

import "github.com/natanaelrusli/go-mux-api/internal/repository"

type ProductUsecase interface {
}

type productUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
	}
}
