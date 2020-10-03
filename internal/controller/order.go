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

type OrderController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	GetByCourierId(w http.ResponseWriter, r *http.Request)
}

type orderController struct {
	orderStorage db.Storage
	config       *configs.Config
}

func NewOrderController(orderStorage db.Storage, config *configs.Config) OrderController {
	return &orderController{orderStorage, config}
}

func (c *orderController) Delete(w http.ResponseWriter, r *http.Request) {
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

	storage := c.orderStorage.(*db.OrderStorage)
	if err := storage.DeleteById(userId); err != nil {
		utils.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, nil)
}

func (c *orderController) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.ErrorHandler(w, errors.New("id query param is not valid"), http.StatusBadRequest)
		return
	}

	orderId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	storage := c.orderStorage.(*db.OrderStorage)
	order, err := storage.SelectById(orderId)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, order)
}

func (c *orderController) GetByCourierId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "courier_id")
	if id == "" {
		utils.ErrorHandler(w, errors.New("courier_id query param is not valid"), http.StatusBadRequest)
		return
	}

	courierId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	storage := c.orderStorage.(*db.OrderStorage)
	orders, err := storage.SelectByCourier(courierId)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, orders)
}

func (c *orderController) Create(w http.ResponseWriter, r *http.Request) {
	var order db.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	storage := c.orderStorage.(*db.OrderStorage)
	if err := storage.Insert(order); err != nil {
		utils.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, nil)
}
