package request

type ProductRequestBody struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductsRequestBody struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ProductRequestParams struct {
	ProductID int `json:"id"`
}
