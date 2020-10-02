package controller

import (
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"net/http"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userStorage db.Storage
}

func NewUserController(userStorage db.Storage) UserController {
	return &userController{userStorage}
}

func (c *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create User"))
	//storage := c.userStorage.(*db.UserStorage)
}

func (t *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete User"))
}