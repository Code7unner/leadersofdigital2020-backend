package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func ErrorHandler(w http.ResponseWriter, err error, code int) {
	var errorResponse ErrorResponse
	errorResponse.Error = err.Error()
	errorResponse.Code = code

	w.WriteHeader(code)

	json.NewEncoder(w).Encode(errorResponse)
}