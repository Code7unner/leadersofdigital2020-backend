package controller

import (
	"encoding/json"
	"errors"
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"github.com/code7unner/leadersofdigital2020-backend/utils"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type ProductsController interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProductsByType(w http.ResponseWriter, r *http.Request)
	GetProductsByOrderId(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type productsController struct {
	productsStorage db.Storage
	config          *configs.Config
}

func NewProductsController(productsStorage db.Storage, config *configs.Config) ProductsController {
	return &productsController{productsStorage, config}
}

func (c *productsController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.ErrorHandler(w, errors.New("id query param is not valid"), http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	storage := c.productsStorage.(*db.ProductStorage)
	if err := storage.DeleteById(userId); err != nil {
		utils.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, nil)
}

func (c *productsController) GetProductsByOrderId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "order_id")
	if id == "" {
		utils.ErrorHandler(w, errors.New("order_id query param is not valid"),http.StatusBadRequest)
		return
	}

	orderId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	storage := c.productsStorage.(*db.ProductStorage)
	products, err := storage.GetProductsByOrder(orderId)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, products)
}

func (c *productsController) GetProductsByType(w http.ResponseWriter, r *http.Request) {
	productType := chi.URLParam(r, "type")
	if productType == "" {
		utils.ErrorHandler(w, errors.New("type query param is not valid"),http.StatusBadRequest)
		return
	}

	storage := c.productsStorage.(*db.ProductStorage)
	products, err := storage.SelectByType(productType)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, products)
}

func (c *productsController) Create(w http.ResponseWriter, r *http.Request) {
	var product db.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	storage := c.productsStorage.(*db.ProductStorage)
	if err := storage.Insert(product); err != nil {
		utils.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, nil)
}

func (c *productsController) GetProducts(w http.ResponseWriter, r *http.Request) {
	storage := c.productsStorage.(*db.ProductStorage)
	products, err := storage.GetAllProducts()
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, products)
}
