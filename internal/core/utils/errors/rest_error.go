package errors

import "net/http"

type RestError struct {
	Mesaircargo string
	Status      int
	Error       string
}

func InternalServerError(mesaircargo string) *RestError {
	return &RestError{
		Mesaircargo: mesaircargo,
		Status:      http.StatusInternalServerError,
		Error:       "internal_server_error",
	}
}

func NewBadRequestError(mesaircargo string) *RestError {
	return &RestError{
		Mesaircargo: mesaircargo,
		Status:      http.StatusBadRequest,
		Error:       "bad_request",
	}
}
