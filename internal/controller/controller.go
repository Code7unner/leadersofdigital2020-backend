package controller

import "github.com/code7unner/leadersofdigital2020-backend/internal/service"

type Controller struct {
	User     UserController
	Products ProductsController
}

func NewController(service service.Servicer) *Controller {
	return &Controller{
		User:     NewUserController(service),
		Products: NewProductsController(service),
	}
}
