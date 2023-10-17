package handler

import (
	"net/http"

	"github.com/natanaelrusli/go-mux-api/internal/usecase"
	"github.com/natanaelrusli/go-mux-api/internal/utils"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}

func (h *ProductHandler) GetProductList(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, nil, "success", http.StatusOK)
}
