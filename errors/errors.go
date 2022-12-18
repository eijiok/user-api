package errors

import (
	"fmt"
	"net/http"
	"strings"
)

type HttpError struct {
	StatusCode int
	Message    string
	Err        error
}

func (err *HttpError) Error() string {
	return fmt.Sprintf("%s - StatusCode: %d - Message: %s", err.Err.Error(), err.StatusCode, err.Message)
}

func NewInternalServerError(err error) *HttpError {
	return &HttpError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Error",
		Err:        err,
	}
}

func NewInternalServerErrorWithMessage(err error, message ...string) *HttpError {
	return &HttpError{
		StatusCode: http.StatusInternalServerError,
		Message:    strings.Join(message, " "),
		Err:        err,
	}
}

func NewHttpError(err error, statusCode int, message string) *HttpError {
	return &HttpError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}
