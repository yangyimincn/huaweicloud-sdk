package hwerror

import "fmt"

type ServerError struct {
	err        error
	name       string
	statusCode int
	message    string
}

func NewServerError(statusCode int, message string) ServerError {
	return ServerError{
		err:        fmt.Errorf("%s", message),
		name:       "ServerError",
		statusCode: statusCode,
		message:    message,
	}
}

func (e ServerError) Error() string {
	return fmt.Sprintf("huaweicloud: %s %s", e.name, e.err)
}

func (e ServerError) Name() string {
	return e.name
}

func (e ServerError) StatusCode() int {
	return e.statusCode
}

func (e ServerError) Message() string {
	return e.message
}