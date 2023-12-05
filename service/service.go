package service

import (
	"errors"
	"fmt"
)

// generic http errors
var (
	ErrNotImplemented      = errors.New("not implemented")
	ErrNotFound            = errors.New("not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrUnprocessableEntity = errors.New("syntax error in json")
)

type ServiceError struct {
	appErr error // error returned by the application, adapter, external package or go lib
	svcErr error // generic error returned from the service e.g. http, grpc, database, etc
}

func (svc *ServiceError) Error() string {
	return fmt.Sprintf("error: %v - %v", svc.svcErr, svc.appErr)
}

func NewError(appError error, svc error) error {
	return &ServiceError{appErr: appError, svcErr: svc}
}
