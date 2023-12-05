package service

import "fmt"

type serviceError struct {
	appErr string // programmer-defined message
	svcErr error  // service error e.g. http, grpc, database, etc
}

func (svc *serviceError) Error() string {
	return fmt.Sprintf("error: %v - %v", svc.svcErr, svc.appErr)
}

func NewError(message string, svc error) error {
	return &serviceError{appErr: message, svcErr: fmt.Errorf("service-error type: %w", svc)}
}
