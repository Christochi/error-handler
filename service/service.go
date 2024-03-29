// Error Interface

package service

import (
	"fmt"
	"strconv"
)

type ServiceError struct {
	appErr  error  // error returned by the application, adapter, external package or go lib
	message string // error description, generic error or error status codes
}

func (svc *ServiceError) Error() string {
	return fmt.Sprintf("Error: %v, Desc: %v", svc.appErr, svc.message)
}

func (svc *ServiceError) AppError() error {
	return svc.appErr
}

func (svc *ServiceError) ErrDesc() string {
	return svc.message
}

func NewError(appError error, message any) error {

	switch msg := message.(type) {
	case int:
		return &ServiceError{appErr: appError, message: strconv.Itoa(msg)}
	case error:
		return &ServiceError{appErr: appError, message: msg.Error()}
	default:
		return &ServiceError{appErr: appError, message: msg.(string)}
	}
}
