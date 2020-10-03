package controller

import (
	"encoding/json"
	"errors"
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/auth"
	"github.com/code7unner/leadersofdigital2020-backend/internal/controller/model"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"github.com/code7unner/leadersofdigital2020-backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	// TODO:
	DeleteUser(w http.ResponseWriter, r *http.Request)
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

func (c *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
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

func (t *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete User"))
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

func (c *userController) Register(w http.ResponseWriter, r *http.Request) {
	var registerUser model.RegisterUser

	if err := json.NewDecoder(r.Body).Decode(&registerUser); err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	expiresAt := time.Now().Add(time.Hour * 100).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &auth.TokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		RegisterUser: registerUser,
	}

	tokenString, err := token.SignedString([]byte(c.config.TokenSecret))
	if err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
	}

	data := auth.AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	}
	utils.SuccessHandler(w, http.StatusOK, data)
}
