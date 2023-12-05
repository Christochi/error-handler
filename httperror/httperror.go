package httperror

import "errors"

// generic http errors
var (
	ErrNotImplemented      = errors.New("not implemented")
	ErrNotFound            = errors.New("not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
)

type APIError struct {
	Message string
	Status  int
}
