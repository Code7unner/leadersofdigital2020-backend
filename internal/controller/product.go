package controller

import (
	"github.com/code7unner/leadersofdigital2020-backend/internal/service"
	"net/http"
)

type ProductsController interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
}

type productsController struct {
	service service.Servicer
}

func NewProductsController(service service.Servicer) ProductsController {
	return &productsController{service}
}

func (p productsController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get products"))
}
