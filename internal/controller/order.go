package controller

import (
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"net/http"
)

type OrderController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type orderController struct {
	orderStorage db.Storage
}

func NewOrderController(orderStorage db.Storage) OrderController {
	return &orderController{orderStorage}
}

func (c *orderController) Create(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
