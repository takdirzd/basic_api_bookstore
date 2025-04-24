package utils

import (
	"encoding/json"
	"net/http"

	"bookstore_backend/model"
)

func JSONResponse(w http.ResponseWriter, status int, method, description string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := model.Response{
		Status:      status,
		Method:      method,
		Description: description,
	}

	json.NewEncoder(w).Encode(resp)
}
