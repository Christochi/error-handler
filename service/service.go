package service

import "fmt"

type serviceError struct {
	appErr error // error returned by the application, external package or go lib
	svcErr error // generic error returned from the service e.g. http, grpc, database, etc
}

func (svc *serviceError) Error() string {
	return fmt.Sprintf("error: %v - %v", svc.svcErr, svc.appErr)
}

func NewError(appError error, svc error) error {
	return &serviceError{appErr: appError, svcErr: svc}
}
