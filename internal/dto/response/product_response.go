package response

import (
	"github.com/natanaelrusli/go-mux-api/internal/model"
	"github.com/natanaelrusli/go-mux-api/internal/utils"
)

type ProductResult struct {
	Products   []model.Product  `json:"products"`
	Pagination utils.Pagination `json:"pagination"`
}

type ProductResponse struct {
	Product model.Product `json:"product"`
}

type ProductsResponse struct {
	Products []model.Product `json:"products"`
}
