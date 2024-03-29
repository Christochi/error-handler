/* Example of how to use the Error interface in service.go
   to handle errors generated by your service
*/

package example

import (
	"net/http"
	"os"
	"strconv"

	"github.com/Christochi/error-handler/service"
)

func openFile() error {
	if _, err := os.Open("no-file"); err != nil {
		return service.NewError(err, os.ErrNotExist) // application error, file does not exist
	}

	return nil
}

func goToUrl() error {
	_, err := http.Get("https://www.googles.com")
	if err != nil {
		return service.NewError(err, strconv.Itoa(http.StatusNotFound)+"-"+http.StatusText(http.StatusNotFound))
	}

	return nil
}
