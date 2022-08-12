package errors

import "net/http"

type RestError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	ErrorType  string `json:"error_type"`
}

// NewBadRequestError Create a new bad request error.
func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		ErrorType:  "bad_request",
	}
}

// NewInternalServerError Create a new internal server error.
func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		ErrorType:  "server_error",
	}
}

// NewNotFoundError Create a new not found error.
func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusNotFound,
		ErrorType:  "not_found",
	}
}
