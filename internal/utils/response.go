package utils

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse represents the structure of a success response
type SuccessResponse struct {
	Status string      `json:"status"`
	Result interface{} `json:"result,omitempty"`
}

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Status string      `json:"status"`
	Error  interface{} `json:"error,omitempty"`
}

// SendJSONResponse sends a JSON response with the specified status code and data
func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// SendSuccessResponse sends a success response with the specified status code and result
func SendSuccessResponse(w http.ResponseWriter, statusCode int, result interface{}) {
	SendJSONResponse(w, statusCode, SuccessResponse{
		Status: "success",
		Result: result,
	})
}

// SendErrorResponse sends an error response with the specified status code and error
func SendErrorResponse(w http.ResponseWriter, statusCode int, err interface{}) {
	SendJSONResponse(w, statusCode, ErrorResponse{
		Status: "error",
		Error:  err,
	})
}
