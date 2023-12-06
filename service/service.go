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
	return fmt.Sprintf("error: %v - %v", svc.message, svc.appErr)
}

func NewError(appError error, message any) error {

	switch desc := message.(type) {
	case int:
		return &ServiceError{appErr: appError, message: strconv.Itoa(desc)}
	case error:
		return &ServiceError{appErr: appError, message: desc.Error()}
	default:
		return &ServiceError{appErr: appError, message: desc.(string)}
	}
}
