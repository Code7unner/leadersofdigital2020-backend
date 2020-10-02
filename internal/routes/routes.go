package routes

import (
	"github.com/code7unner/leadersofdigital2020-backend/internal/controller"
	"github.com/code7unner/leadersofdigital2020-backend/internal/service"
	"github.com/go-chi/chi"
)

func InitRoutes(service *service.Service) func(r chi.Router) {
	return func(r chi.Router) {
		c := controller.NewController(service)

		// User requests
		r.HandleFunc("/user/create", c.User.CreateUser)
		r.HandleFunc("/user/delete", c.User.DeleteUser)

		// Product requests
		r.HandleFunc("/products/get", c.Products.GetProducts)
	}
}