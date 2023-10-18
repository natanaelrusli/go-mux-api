package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/natanaelrusli/go-mux-api/internal/db"
	"github.com/natanaelrusli/go-mux-api/internal/handler"
	"github.com/natanaelrusli/go-mux-api/internal/repository"
	"github.com/natanaelrusli/go-mux-api/internal/usecase"
)

func main() {
	database := db.NewPostgresDB()

	r := mux.NewRouter()

	productRepository := repository.NewProductRepository(database)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)

	r.HandleFunc("/product", productHandler.CreateOneProduct).Methods("POST")
	r.HandleFunc("/products", productHandler.GetProductList).Methods("GET")

	srv := &http.Server{
		Addr:    ":8020",
		Handler: r,
	}

	log.Println("Running HTTP server at: ", srv.Addr)
	if err := http.ListenAndServe(srv.Addr, r); err != nil {
		log.Fatal(err.Error())
	}
}
