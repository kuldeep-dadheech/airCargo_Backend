package errors

import "net/http"

type RestError struct{
	Message string 
	Status int
	Error string
} 
func InternalServerError(message string) *RestError{
	return &RestError{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "internal_server_error",
	}
}

func NewBadRequestError(message string) *RestError{
	return &RestError{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}
}
