package controller

import (
	"github.com/code7unner/leadersofdigital2020-backend/internal/service"
	"net/http"
)

type TestController interface {
	TestCtrl(w http.ResponseWriter, r *http.Request)
}

type testController struct {
	service service.Servicer
}

func NewTestController(service service.Servicer) TestController {
	return &testController{service}
}

func (t testController) TestCtrl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}
