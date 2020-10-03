package routes

import (
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/controller"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"github.com/go-chi/chi"
)

func InitRoutes(storage db.Storage, config *configs.Config) func(r chi.Router) {
	return func(r chi.Router) {
		c := controller.NewController(storage, config)

		// User requests
		r.HandleFunc("/user/create", c.User.CreateUser)
		r.HandleFunc("/user/register", c.User.Register)
		r.HandleFunc("/user/{id}", c.User.GetUser)

		// Product requests
		r.HandleFunc("/products/get", c.Products.GetProducts)
		r.HandleFunc("/products/create", c.Products.Create)
		r.HandleFunc("/products/get/{type}", c.Products.GetProductsByType)
		r.HandleFunc("/products/get/{order_id}", c.Products.GetProductsByOrderId)

		// Order requests
		r.HandleFunc("/order/create", c.Order.Create)
		r.HandleFunc("/order/{id}", c.Order.GetById)
		r.HandleFunc("/order/{courier_id}", c.Order.GetByCourierId)

		// Store requests
		r.HandleFunc("/store/create", c.Store.Create)
	}
}
