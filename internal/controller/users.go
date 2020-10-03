package controller

import (
	"encoding/json"
	"errors"
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"github.com/code7unner/leadersofdigital2020-backend/utils"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	// TODO:
	Login(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userStorage db.Storage
	config      *configs.Config
}

func NewUserController(userStorage db.Storage, config *configs.Config) UserController {
	return &userController{userStorage, config}
}

func (c *userController) Create(w http.ResponseWriter, r *http.Request) {
	var user db.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	storage := c.userStorage.(*db.UserStorage)
	if err := storage.Insert(user); err != nil {
		utils.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, nil)
}

func (c *userController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.ErrorHandler(w, errors.New("id query param is not valid"), http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	storage := c.userStorage.(*db.UserStorage)
	if err := storage.DeleteById(userId); err != nil {
		utils.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, nil)
}

func (c *userController) Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login User"))
}

func (c *userController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.ErrorHandler(w, errors.New("id query param is not valid"), http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	storage := c.userStorage.(*db.UserStorage)
	user, err := storage.GetUserById(userId)
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, user)
}
