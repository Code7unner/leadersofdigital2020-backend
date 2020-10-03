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

type StoreController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type storeController struct {
	storeStorage db.Storage
	config       *configs.Config
}

func NewStoreController(storeStorage db.Storage, config *configs.Config) StoreController {
	return &storeController{storeStorage, config}
}

func (c *storeController) Delete(w http.ResponseWriter, r *http.Request) {
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

	storage := c.storeStorage.(*db.StoreStorage)
	if err := storage.DeleteById(userId); err != nil {
		utils.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, nil)
}

func (c *storeController) Create(w http.ResponseWriter, r *http.Request) {
	var store db.Store
	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	storage := c.storeStorage.(*db.StoreStorage)
	if err := storage.Insert(store); err != nil {
		utils.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessHandler(w, http.StatusOK, nil)
}
