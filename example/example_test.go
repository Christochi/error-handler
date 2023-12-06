// Test example
package example

import (
	"errors"
	"testing"

	"github.com/Christochi/error-handler/service"
)

func TestOpenFile(t *testing.T) {
	err := openFile()

	var svcErr *service.ServiceError
	if errors.As(err, &svcErr) {
		t.Errorf("want %v, got %v", nil, err)
	}
}

func TestGoToURL(t *testing.T) {
	err := goToUrl()

	var svcErr *service.ServiceError
	if errors.As(err, &svcErr) {
		t.Errorf("want %v, got %v", nil, err)
	}
}
