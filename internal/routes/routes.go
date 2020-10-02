package routes

import (
	"github.com/code7unner/leadersofdigital2020-backend/internal/controller"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"github.com/go-chi/chi"
)

func InitRoutes(storage db.Storage) func(r chi.Router) {
	return func(r chi.Router) {
		c := controller.NewController(storage)

		// User requests
		r.HandleFunc("/user/create", c.User.CreateUser)
		r.HandleFunc("/user/delete", c.User.DeleteUser)

		// Product requests
		r.HandleFunc("/products/get", c.Products.GetProducts)
	}
}