package httpapi

import (
	"encoding/json"
	"net/http"
)

// APIErrorResponse Response wrapper for error responses
type APIErrorResponse struct {
	Message string `json:"errorMessage"`
	Code    int    `json:"errorCode"`
}

// Ok - handles sending http ok response
func Ok(w http.ResponseWriter, data interface{}) {
	sendResponse(w, http.StatusOK, data)
}

// Success - handles sending http success responses
func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	sendResponse(w, statusCode, data)
}

// Error - handles sending http error responses
func Error(w http.ResponseWriter, statusCode int, msg string) {
	var errorResponse APIErrorResponse
	errorResponse.Message = msg
	errorResponse.Code = statusCode
	sendResponse(w, statusCode, errorResponse)
}

func sendResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}
