package errors

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Message    string
	Err        error
}

func (err *HttpError) Error() string {
	return fmt.Sprintf("%s - StatusCode: %d - Message: %s", err.Error(), err.StatusCode, err.Message)
}

func NewHttpError(err error) *HttpError {
	return &HttpError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Error",
		Err:        err,
	}
}
