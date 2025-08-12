package errors

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/theothertomelliott/acyclic"
)

var _ Error = &HTTPError{}

// HTTPError is a struct that represents an HTTP error
// and it implements the Error interface.
type HTTPError struct {
	StatusCode int    `json:"statusCode"`
	Name       string `json:"name"`
	Message    string `json:"message"`
	Err        error  `json:"error"`
}

// Error returns a string concatenation of the name
// and message. If th Err property is not nil, it
// will be returned in the message.
func (e *HTTPError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("%s: %s", e.Name, e.Message)
	}

	return fmt.Sprintf(
		"%s: %s (%s)",
		e.Name,
		e.Message,
		e.Err.Error(),
	)
}

// JSON returns the bytes of the JSON representation
// of the error.
//
// If the Err property is not marshalable, it will
// be replaced with a marshalable error indicating
// why it is not marshalable.
//
// If the Err property is not an implementation of
// the Error interface, it will be replaced with a wrapped
// version of it to ensure that the error is marshallable
// and content will be shown.
func (e HTTPError) JSON() []byte {
	if e.Err != nil {
		if _, ok := e.Err.(Error); !ok {
			e.Err = Wrap(e.Err)
		}
	}

	if err := acyclic.Check(e); err != nil {
		e.Err = New(fmt.Sprintf(
			"HTTPError.Err contains a circular reference (%s), original: %s",
			err.Error(),
			e.Err.Error(),
		))
	}

	b, err := json.Marshal(e)
	if err != nil {
		e.Err = New(fmt.Sprintf(
			"HTTPError.Err is not marshallable (%s), original: %s",
			err.Error(),
			e.Err.Error(),
		))

		b, _ = json.Marshal(e)
	}

	return b
}

// NewHTTPError creates a new HTTPError.
func NewHTTPError(statusCode int, name, message string, err error) error {
	return &HTTPError{
		StatusCode: statusCode,
		Name:       name,
		Message:    message,
		Err:        err,
	}
}

// NewBadRequestError creates a new HTTPError with a 400 status code.
func NewBadRequestError(message string, err error) error {
	return NewHTTPError(http.StatusBadRequest, "bad_request_error", message, err)
}

// NewNotFoundError creates a new HTTPError with a 404 status code.
func NewNotFoundError(message string, err error) error {
	return NewHTTPError(http.StatusNotFound, "not_found_error", message, err)
}

// NewInternalServerError creates a new HTTPError with a 500 status code.
func NewInternalServerError(message string, err error) error {
	return NewHTTPError(http.StatusInternalServerError, "internal_server_error", message, err)
}

// NewUnauthorizedError creates a new HTTPError with a 401 status code.
func NewUnauthorizedError(message string, err error) error {
	return NewHTTPError(http.StatusUnauthorized, "unauthorized_error", message, err)
}

// NewForbiddenError creates a new HTTPError with a 403 status code.
func NewForbiddenError(message string, err error) error {
	return NewHTTPError(http.StatusForbidden, "forbidden_error", message, err)
}

// NewConflictError creates a new HTTPError with a 409 status code.
func NewConflictError(message string, err error) error {
	return NewHTTPError(http.StatusConflict, "conflict_error", message, err)
}

// NewTooManyRequestsError creates a new HTTPError with a 429 status code.
func NewTooManyRequestsError(message string, err error) error {
	return NewHTTPError(http.StatusTooManyRequests, "too_many_requests_error", message, err)
}

// NewBadGatewayError creates a new HTTPError with a 502 status code.
func NewBadGatewayError(message string, err error) error {
	return NewHTTPError(http.StatusBadGateway, "bad_gateway_error", message, err)
}

// NewServiceUnavailableError creates a new HTTPError with a 503 status code.
func NewServiceUnavailableError(message string, err error) error {
	return NewHTTPError(http.StatusServiceUnavailable, "service_unavailable_error", message, err)
}

// NewGatewayTimeoutError creates a new HTTPError with a 504 status code.
func NewGatewayTimeoutError(message string, err error) error {
	return NewHTTPError(http.StatusGatewayTimeout, "gateway_timeout_error", message, err)
}
