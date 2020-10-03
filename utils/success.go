package utils

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func SuccessHandler(w http.ResponseWriter, code int, data interface{}) {
	var successResponse = SuccessResponse{code,data}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(successResponse)
}
