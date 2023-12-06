# error-handler

## Description
Generic error package that can be used to handle errors in your applications, services or any other "adapter" your service may use

## Setup
- import the package using `go get` or `go mod tidy` commands
  - using `go get`: on the terminal of your project directory, paste: `go get github.com/Christochi/error-handler`
  - paste **github.com/Christochi/error-handler/service** on the import block of your code, and on the terminal, run `go mod tidy`

### Usage Example
~~~
func openFile() error {
	if _, err := os.Open("no-file"); err != nil {
		return service.NewError(err, "file does not exist")
	}

	return nil
}

~~~
err := openFile()

var svcErr *service.ServiceError
if errors.As(err, &svcErr){
   // your implementation
}
~~~
~~~
