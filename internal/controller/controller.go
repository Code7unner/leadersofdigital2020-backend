package controller

import (
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
)

type Controller struct {
	User     UserController
	Products ProductsController
	Order    OrderController
	Store    StoreController
}

func NewController(storage db.Storage, config *configs.Config) *Controller {
	return &Controller{
		User:     NewUserController(storage, config),
		Products: NewProductsController(storage, config),
		Order:    NewOrderController(storage, config),
		Store:    NewStoreController(storage, config),
	}
}
