package errors

import (
	"fmt"
	"net/http"
	"strings"
)

type HttpError struct {
	StatusCode int
	Message    string
	Err        *error
}

func (err *HttpError) Error() string {
	details := ""
	if err.Err != nil {
		err := *(err.Err)
		details = "- " + err.Error()
	}

	return fmt.Sprintf("%s - StatusCode: %d %s", err.Message, err.StatusCode, details)
}

func NewInternalServerError(err *error) *HttpError {
	return &HttpError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Error",
		Err:        err,
	}
}

func NewInternalServerErrorWithMessage(err *error, message ...string) *HttpError {
	return &HttpError{
		StatusCode: http.StatusInternalServerError,
		Message:    strings.Join(message, " "),
		Err:        err,
	}
}

func NewHttpError(statusCode int, message string, err *error) *HttpError {
	return &HttpError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}
