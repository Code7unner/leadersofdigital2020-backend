package controller

import (
	"encoding/json"
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/auth"
	"github.com/code7unner/leadersofdigital2020-backend/internal/controller/model"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"github.com/code7unner/leadersofdigital2020-backend/utils"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userStorage db.Storage
	config      *configs.Config
}

func NewUserController(userStorage db.Storage, config *configs.Config) UserController {
	return &userController{userStorage, config}
}

func (c *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create User"))
}

func (t *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete User"))
}

func (c *userController) Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login User"))
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

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(auth.AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	})
}
