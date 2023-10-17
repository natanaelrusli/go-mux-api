package handler

import (
	"encoding/json"
	"net/http"

	"github.com/natanaelrusli/go-mux-api/internal/model"
	"github.com/natanaelrusli/go-mux-api/internal/usecase"
	"github.com/natanaelrusli/go-mux-api/internal/utils"
	cerrors "github.com/natanaelrusli/go-mux-api/internal/utils/errors"
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
	res, err := h.productUsecase.GetList()

	if err != nil {
		utils.RespondWithJSON(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, res, "success", http.StatusOK)
}

func (h *ProductHandler) CreateOneProduct(w http.ResponseWriter, r *http.Request) {
	var p model.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		return
	}
	defer r.Body.Close()

	res, err := h.productUsecase.CreateOne(p)

	if err != nil {
		utils.RespondWithJSON(w, nil, cerrors.NewBadFormattedRequest().Message, http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, res, "success", http.StatusCreated)
}
