package controller

import (
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"net/http"
)

type ProductsController interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
}

type productsController struct {
	productsStorage db.Storage
}

func NewProductsController(productsStorage db.Storage) ProductsController {
	return &productsController{productsStorage}
}

func (c *productsController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get products"))

	//storage := c.productsStorage.(*db.ProductStorage)
}
