package controller

import (
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"net/http"
)

type ProductsController interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
}

type productsController struct {
	productsStorage db.Storage
	config          *configs.Config
}

func NewProductsController(productsStorage db.Storage, config *configs.Config) ProductsController {
	return &productsController{productsStorage, config}
}

func (c *productsController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get products"))
}
