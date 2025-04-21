package helpers

import (
	"encoding/json"
	"net/http"
)

type ErrorResult struct {
	Message string `json:"message"`
}

type JSONResponse struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

type OperationResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SendJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	var result = JSONResponse{
		Data:  data,
		Error: nil,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(result)
}

func SendError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var result = JSONResponse{
		Data:  nil,
		Error: ErrorResult{Message: message},
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(result)
}
