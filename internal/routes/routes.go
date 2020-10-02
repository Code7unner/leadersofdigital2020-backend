package routes

import (
	"github.com/code7unner/leadersofdigital2020-backend/internal/controller"
	"github.com/code7unner/leadersofdigital2020-backend/internal/service"
	"github.com/go-chi/chi"
)

func TestRoutes(service *service.Service) func(r chi.Router) {
	return func(r chi.Router) {
		testController := controller.NewTestController(service)

		r.HandleFunc("/test", testController.TestCtrl)
	}
}