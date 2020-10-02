package controller

import (
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"net/http"
)

type OrderController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type orderController struct {
	orderStorage db.Storage
	config       *configs.Config
}

func NewOrderController(orderStorage db.Storage, config *configs.Config) OrderController {
	return &orderController{orderStorage, config}
}

func (c *orderController) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create order"))
}
