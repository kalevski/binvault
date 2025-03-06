package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type JSONResult struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

type OperationResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func JSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	var result = JSONResult{
		Data:  data,
		Error: nil,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(result)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var result = JSONResult{
		Data:  nil,
		Error: ResponseError{Message: message},
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(result)
}
