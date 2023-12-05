package service

import "fmt"

type ServiceError struct {
	appErr string // description of the error
	svcErr error  // generic error returned from the service e.g. http, grpc, database, etc
}

func (svc *ServiceError) Error() string {
	return fmt.Sprintf("error: %v - %v", svc.svcErr, svc.appErr)
}

func NewError(message string, svc error) error {
	return &ServiceError{appErr: message, svcErr: fmt.Errorf("service-error type: %w", svc)}
}
