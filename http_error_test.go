package errors

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"

	"github.com/iolave/go-errors/internal"
)

func TestHTTPError_Error(t *testing.T) {
	t.Parallel()

	t.Run("should return error message without original error", func(t *testing.T) {
		err := &HTTPError{
			Name:    internal.GenerateRandomString(10),
			Message: internal.GenerateRandomString(10),
		}
		want := fmt.Sprintf("%s: %s", err.Name, err.Message)
		got := err.Error()

		if got != want {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})

	t.Run("should return error message with original", func(t *testing.T) {
		err := &HTTPError{
			Name:    internal.GenerateRandomString(10),
			Message: internal.GenerateRandomString(10),
			Err:     &DummyError{},
		}
		want := fmt.Sprintf("%s: %s (%s)", err.Name, err.Message, err.Err.Error())
		got := err.Error()

		if got != want {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestHTTPError_JSON(t *testing.T) {
	t.Parallel()

	t.Run("should return json error without original", func(t *testing.T) {
		err := &HTTPError{
			StatusCode: rand.Int(),
			Name:       internal.GenerateRandomString(10),
			Message:    internal.GenerateRandomString(10),
		}
		want := fmt.Sprintf(
			`{"statusCode":%d,"name":"%s","message":"%s","error":null}`,
			err.StatusCode,
			err.Name,
			err.Message,
		)
		got := string(err.JSON())

		if got != want {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})

	t.Run("should return json error with original", func(t *testing.T) {
		dummyErr := &DummyError{}
		wrappedDummyErr := ToError(Wrap(dummyErr))
		err := &HTTPError{
			StatusCode: rand.Int(),
			Name:       internal.GenerateRandomString(10),
			Message:    internal.GenerateRandomString(10),
			Err:        dummyErr,
		}
		t.Log(string(wrappedDummyErr.JSON()))
		want := fmt.Sprintf(
			`{"statusCode":%d,"name":"%s","message":"%s","error":%s}`,
			err.StatusCode,
			err.Name,
			err.Message,
			wrappedDummyErr.JSON(),
		)
		got := string(err.JSON())

		if got != want {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})

	t.Run("should return json error when original is not marshalable", func(t *testing.T) {
		someErr := &AnyError{
			Any: func() {},
		}
		someErrExpected := New(fmt.Sprintf(
			"HTTPError.Err is not marshallable (json: unsupported type: func()), original: %s",
			Wrap(someErr).Error(),
		)).(Error)

		err := &HTTPError{
			StatusCode: rand.Int(),
			Name:       internal.GenerateRandomString(10),
			Message:    internal.GenerateRandomString(10),
			Err:        someErr,
		}

		want := fmt.Sprintf(
			`{"statusCode":%d,"name":"%s","message":"%s","error":%s}`,
			err.StatusCode,
			err.Name,
			err.Message,
			someErrExpected.JSON(),
		)

		got := string(err.JSON())

		if got != want {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})

	t.Run("should return json error when original contains a circular reference", func(t *testing.T) {
		someErr := &AnyError{}
		someErr.Any = someErr
		someErrExpected := New(fmt.Sprintf(
			"HTTPError.Err contains a circular reference (cycle found: [Err Original Any]), original: %s",
			Wrap(someErr).Error(),
		)).(Error)

		err := &HTTPError{
			StatusCode: rand.Int(),
			Name:       internal.GenerateRandomString(10),
			Message:    internal.GenerateRandomString(10),
			Err:        someErr,
		}

		want := fmt.Sprintf(
			`{"statusCode":%d,"name":"%s","message":"%s","error":%s}`,
			err.StatusCode,
			err.Name,
			err.Message,
			someErrExpected.JSON(),
		)

		got := string(err.JSON())

		if got != want {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestNewBadRequestError(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		err := &HTTPError{
			StatusCode: http.StatusBadRequest,
			Name:       "bad_request_error",
			Message:    internal.GenerateRandomString(10),
		}

		want := string(err.JSON())
		got := string(NewBadRequestError(err.Message, err.Err).(Error).JSON())

		if want != got {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestNewNotFoundError(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		err := &HTTPError{
			StatusCode: http.StatusNotFound,
			Name:       "not_found_error",
			Message:    internal.GenerateRandomString(10),
		}

		want := string(err.JSON())
		got := string(NewNotFoundError(err.Message, err.Err).(Error).JSON())

		if want != got {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestNewInternalServerError(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		err := &HTTPError{
			StatusCode: http.StatusInternalServerError,
			Name:       "internal_server_error",
			Message:    internal.GenerateRandomString(10),
		}

		want := string(err.JSON())
		got := string(NewInternalServerError(err.Message, err.Err).(Error).JSON())

		if want != got {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestNewUnauthorizedError(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		err := &HTTPError{
			StatusCode: http.StatusUnauthorized,
			Name:       "unauthorized_error",
			Message:    internal.GenerateRandomString(10),
		}

		want := string(err.JSON())
		got := string(NewUnauthorizedError(err.Message, err.Err).(Error).JSON())

		if want != got {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestNewForbiddenError(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		err := &HTTPError{
			StatusCode: http.StatusForbidden,
			Name:       "forbidden_error",
			Message:    internal.GenerateRandomString(10),
		}

		want := string(err.JSON())
		got := string(NewForbiddenError(err.Message, err.Err).(Error).JSON())

		if want != got {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestNewConflictError(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		err := &HTTPError{
			StatusCode: http.StatusConflict,
			Name:       "conflict_error",
			Message:    internal.GenerateRandomString(10),
		}

		want := string(err.JSON())
		got := string(NewConflictError(err.Message, err.Err).(Error).JSON())

		if want != got {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestNewTooManyRequestsError(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		err := &HTTPError{
			StatusCode: http.StatusTooManyRequests,
			Name:       "too_many_requests_error",
			Message:    internal.GenerateRandomString(10),
		}

		want := string(err.JSON())
		got := string(NewTooManyRequestsError(err.Message, err.Err).(Error).JSON())

		if want != got {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestNewBadGatewayError(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		err := &HTTPError{
			StatusCode: http.StatusBadGateway,
			Name:       "bad_gateway_error",
			Message:    internal.GenerateRandomString(10),
		}

		want := string(err.JSON())
		got := string(NewBadGatewayError(err.Message, err.Err).(Error).JSON())

		if want != got {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestNewServiceUnavailableError(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		err := &HTTPError{
			StatusCode: http.StatusServiceUnavailable,
			Name:       "service_unavailable_error",
			Message:    internal.GenerateRandomString(10),
		}

		want := string(err.JSON())
		got := string(NewServiceUnavailableError(err.Message, err.Err).(Error).JSON())

		if want != got {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestNewGatewayTimeoutError(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		err := &HTTPError{
			StatusCode: http.StatusGatewayTimeout,
			Name:       "gateway_timeout_error",
			Message:    internal.GenerateRandomString(10),
		}

		want := string(err.JSON())
		got := string(NewGatewayTimeoutError(err.Message, err.Err).(Error).JSON())

		if want != got {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}
