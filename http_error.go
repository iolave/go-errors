package errors

import (
	"fmt"
	"io"
	"net/http"
)

const (
	HTTP_ERR_BAD_REQUEST     = "bad_request_error"
	HTTP_ERR_INTERNAL_SERVER = "internal_server_error"
	HTTP_ERR_BAD_GATEWAY     = "bad_gateway_error"
	HTTP_ERR_UNAUTHORIZED    = "unauthorized_error"
	HTTP_ERR_FORBIDDEN       = "forbidden_error"
	HTTP_ERR_GATEWAY_TIMEOUT = "gateway_timeout_error"
	HTTP_ERR_NOT_FOUND       = "not_found_error"
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

// Cause returns the underlying cause of the error
func (err *HttpError) Cause() error {
	return err.cause
}

// Unwrap returns the underlying cause of the error.
// It implements a new optional error interface introduced in Go 1.13.
func (err *HttpError) Unwrap() error {
	return err.cause
}

// StackTrace returns an stack trace of the error
func (err *HttpError) StackTrace() StackTrace {
	return err.stack
}

// Format implements fmt.Formatter and can be formatted by the fmt package. The
// following verbs are supported
//
//	%s    print the error. If the error has a Cause it will be
//	      printed recursively
//	%v    see %s
//	%+v   extended format. Each Frame of the error's StackTrace will
//	      be printed in detail.
//
// TODO: add more info here
func (err *HttpError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, err.message+"\n")
			err.stack.Format(s, verb)
			return
		}

		fallthrough
	case 's':
		io.WriteString(s, err.Error())
	case 'q':
		fmt.Fprintf(s, "%q", err.Error())
	}
}

var BadRequestError = NewKind(HTTP_ERR_BAD_REQUEST)
var InternalServerError = NewKind(HTTP_ERR_INTERNAL_SERVER)
var BadGatewayError = NewKind(HTTP_ERR_BAD_GATEWAY)
var UnauthorizedError = NewKind(HTTP_ERR_UNAUTHORIZED)
var ForbiddenError = NewKind(HTTP_ERR_FORBIDDEN)
var GatewayTimeoutError = NewKind(HTTP_ERR_GATEWAY_TIMEOUT)
var NotFoundError = NewKind(HTTP_ERR_NOT_FOUND)

func NewBadRequestError(msg string) *HttpError {
	err := BadRequestError.newHttpError(msg, http.StatusBadRequest, nil)
	return err
}

func NewBadRequestErrorWithCause(msg string, cause error) *HttpError {
	err := BadRequestError.newHttpError(msg, http.StatusBadRequest, cause)
	return err
}

func NewInternalServerError(msg string) *HttpError {
	err := InternalServerError.newHttpError(msg, http.StatusInternalServerError, nil)
	return err
}

func NewInternalServerErrorWithCause(msg string, cause error) *HttpError {
	err := InternalServerError.newHttpError(msg, http.StatusInternalServerError, cause)
	return err
}

func NewBadGatewayError(msg string) *HttpError {
	err := BadGatewayError.newHttpError(msg, http.StatusBadGateway, nil)
	return err
}

func NewBadGatewayErrorWithCause(msg string, cause error) *HttpError {
	err := BadGatewayError.newHttpError(msg, http.StatusBadGateway, cause)
	return err
}

func NewUnauthorizedError(msg string) *HttpError {
	err := UnauthorizedError.newHttpError(msg, http.StatusUnauthorized, nil)
	return err
}

func NewUnauthorizedErrorWithCause(msg string, cause error) *HttpError {
	err := UnauthorizedError.newHttpError(msg, http.StatusUnauthorized, cause)
	return err
}

func NewForbiddenError(msg string) *HttpError {
	err := ForbiddenError.newHttpError(msg, http.StatusForbidden, nil)
	return err
}

func NewForbiddenErrorWithCause(msg string, cause error) *HttpError {
	err := ForbiddenError.newHttpError(msg, http.StatusForbidden, cause)
	return err
}

func NewGatewayTimeoutError(msg string) *HttpError {
	err := GatewayTimeoutError.newHttpError(msg, http.StatusGatewayTimeout, nil)
	return err
}

func NewGatewayTimeoutErrorWithCause(msg string, cause error) *HttpError {
	err := GatewayTimeoutError.newHttpError(msg, http.StatusGatewayTimeout, cause)
	return err
}

func NewNotFoundError(msg string) *HttpError {
	err := NotFoundError.newHttpError(msg, http.StatusNotFound, nil)
	return err
}

func NewNotFoundErrorWithCause(msg string, cause error) *HttpError {
	err := NotFoundError.newHttpError(msg, http.StatusNotFound, cause)
	return err
}
