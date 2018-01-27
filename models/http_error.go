package models

import "net/http"

// HTTPError error type for handling error in http delivery
type HTTPError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// Error return information about the error
func (e *HTTPError) Error() string {
	return e.Message
}

// NewErrorInternalServer returns Internal Server Error template
func NewErrorInternalServer(msg string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
	}
}

// NewErrorNotFound returns Not Found Error template
func NewErrorNotFound(msg string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
	}
}

// NewErrorUnprocessableEntity returns Unprocessable Entity Error template
func NewErrorUnprocessableEntity(msg string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    msg,
	}
}

// NewErrorUnauthorized returns Unauthorized Error template
func NewErrorUnauthorized(msg string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusUnauthorized,
		Message:    msg,
	}
}
