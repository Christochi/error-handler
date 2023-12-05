package service

import "fmt"

type serviceError struct {
	appErr string // description of the error
	svcErr error  // generic errors retruned from the services e.g. http, grpc, database, etc
}

func (svc *serviceError) Error() string {
	return fmt.Sprintf("error: %v - %v", svc.svcErr, svc.appErr)
}

func NewError(message string, svc error) error {
	return &serviceError{appErr: message, svcErr: fmt.Errorf("service-error type: %w", svc)}
}
