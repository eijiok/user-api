package errors

import (
	"fmt"
	"net/http"
	"strings"
)

type HttpError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Err        *error `json:"error"`
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

type ValidationFieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (v *ValidationFieldError) Error() string {
	return fmt.Sprintf("Error %s : %s. ", v.Field, v.Message)
}

type ValidationError struct {
	Errs []error `json:"details"`
}

func (v *ValidationError) Append(error error) {
	if error == nil {
		return
	}
	v.Errs = append(v.Errs, error)
}

func (v *ValidationError) Error() string {
	var sb strings.Builder
	for _, err := range v.Errs {
		sb.WriteString(err.Error())
	}
	return sb.String()
}

func (v *ValidationError) HasErrors() bool {
	return len(v.Errs) > 0
}
