package model

import (
	"errors"
	"net/http"
)

// Define error constants
var (
	ErrInvalidInput        = NewAPIError("Invalid input data", http.StatusBadRequest)
	ErrUnauthorized        = NewAPIError("Unauthorized access", http.StatusUnauthorized)
	ErrForbidden           = NewAPIError("Forbidden access", http.StatusForbidden)
	ErrNotFound            = NewAPIError("Resource not found", http.StatusNotFound)
	ErrInternalServerError = NewAPIError("Something went wrong, please try again later", http.StatusInternalServerError)
	ErrEmailAlreadyExists  = NewAPIError("Email already exists", http.StatusBadRequest)
	ErrNoContent           = NewAPIError("No information found", http.StatusNoContent)

	ErrNoUpdates = errors.New("no updates found")
	ErrNoDelete  = errors.New("no user found to delete")
)

// APIError struct to represent errors
type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewAPIError(message string, code int) *APIError {
	return &APIError{
		Message: message,
		Code:    code,
	}
}
