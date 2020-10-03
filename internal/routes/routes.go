package routes

import (
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/controller"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"github.com/go-chi/chi"
)

func InitRoutes(config *configs.Config, storages ...db.Storage) func(r chi.Router) {
	return func(r chi.Router) {
		c := controller.NewController(config, storages...)

		// User requests
		r.Post("/user/create", c.User.Create)
		r.Get("/user/delete/{id}", c.User.Delete)
		r.Get("/user/{id}", c.User.GetUser)

		// Product requests
		r.Get("/products/get", c.Products.GetProducts)
		r.Post("/products/create", c.Products.Create)
		r.Get("/products/delete/{id}", c.Products.Delete)
		r.Get("/products/get/{type}", c.Products.GetProductsByType)
		r.Get("/products/get/{order_id}", c.Products.GetProductsByOrderId)

		// Order requests
		r.Post("/order/create", c.Order.Create)
		r.Post("/order/delete/{id}", c.Order.Delete)
		r.Get("/order/{id}", c.Order.GetById)
		r.Get("/order/{courier_id}", c.Order.GetByCourierId)

		// Store requests
		r.Post("/store/create", c.Store.Create)
		r.Get("/store/delete/{id}", c.Store.Delete)
	}
}
