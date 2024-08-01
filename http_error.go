package errors

import (
	"fmt"
	"net/http"
)

// Error represents an error of some Kind, implements the error interface
type HttpError struct {
	kind       *Kind      `json:"-"`
	name       string     `json:"name"`
	statusCode int        `json:"statusCode"`
	message    string     `json:"message"`
	cause      error      `json:"cause"`
	stack      StackTrace `json:"stack"`
}

func (err *HttpError) Error() string {
	if err.cause == nil {
		return err.message
	}

	return fmt.Sprintf(err.message, err.cause.Error())
}

var BadRequestError = NewKind("bad_request_error")
var InternalServerError = NewKind("internal_server_error")

func NewBadRequestError(msg string) *HttpError {
	err := BadRequestError.NewHttpError(msg, http.StatusBadRequest, nil)
	return err
}

func NewBadRequestErrorWithCause(msg string, cause error) *HttpError {
	err := BadRequestError.NewHttpError(msg, http.StatusBadRequest, cause)
	return err
}

func NewInternalServerError(msg string) *HttpError {
	err := InternalServerError.NewHttpError(msg, http.StatusInternalServerError, nil)
	return err
}

func NewInternalServerErrorWithCause(msg string, cause error) *HttpError {
	err := InternalServerError.NewHttpError(msg, http.StatusInternalServerError, cause)
	return err
}
