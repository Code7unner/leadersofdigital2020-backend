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

func NewController(config *configs.Config, storages ...db.Storage) *Controller {
	return &Controller{
		User:     NewUserController(storages[0], config),
		Products: NewProductsController(storages[1], config),
		Order:    NewOrderController(storages[2], config),
		Store:    NewStoreController(storages[3], config),
	}
}
