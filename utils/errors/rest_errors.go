package errors

import "net/http"

type RestError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	ErrorType  string `json:"error_type"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		ErrorType:  "bad_request",
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		ErrorType:  "bad_request",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusNotFound,
		ErrorType:  "not_found",
	}
}
