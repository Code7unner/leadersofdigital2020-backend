package controller

import (
	"github.com/code7unner/leadersofdigital2020-backend/internal/service"
	"net/http"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service service.Servicer
}

func NewUserController(service service.Servicer) UserController {
	return &userController{service}
}

func (t *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create User"))
}

func (t *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete User"))
}