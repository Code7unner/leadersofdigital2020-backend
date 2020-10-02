package controller

import (
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
)

type Controller struct {
	User         UserController
	Products     ProductsController
}

func NewController(storage db.Storage) *Controller {
	return &Controller{
		User:         NewUserController(storage),
		Products:     NewProductsController(storage),
	}
}
