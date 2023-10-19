package handler

import (
	"encoding/json"
	"net/http"

	"github.com/natanaelrusli/go-mux-api/internal/dto/request"
	"github.com/natanaelrusli/go-mux-api/internal/dto/response"
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
	p := new(response.ProductResult)
	req := new(request.ProductsRequestBody)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		utils.RespondWithJSON(w, nil, cerrors.NewBadFormattedRequest().Message, http.StatusInternalServerError)
		return
	}

	res, err := h.productUsecase.GetList(req.Page, req.Limit)

	if err != nil {
		utils.RespondWithJSON(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	p.Products = res.Products
	p.Pagination = res.Pagination

	utils.RespondWithJSON(w, p, "success", http.StatusOK)
}

func (h *ProductHandler) CreateOneProduct(w http.ResponseWriter, r *http.Request) {
	p := new(request.ProductRequestBody)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		utils.RespondWithJSON(w, nil, cerrors.NewBadFormattedRequest().Message, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	res, err := h.productUsecase.CreateOne(*p)

	if err != nil {
		utils.RespondWithJSON(w, nil, cerrors.NewBadFormattedRequest().Message, http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, res, "success", http.StatusCreated)
}
