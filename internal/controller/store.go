package controller

import (
	"github.com/code7unner/leadersofdigital2020-backend/configs"
	"github.com/code7unner/leadersofdigital2020-backend/internal/db"
	"net/http"
)

type StoreController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type storeController struct {
	storeStorage db.Storage
	config       *configs.Config
}

func NewStoreController(storeStorage db.Storage, config *configs.Config) StoreController {
	return &storeController{storeStorage, config}
}

func (c *storeController) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create store"))
}
