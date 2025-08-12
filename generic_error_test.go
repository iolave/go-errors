package errors

import (
	"errors"
	"fmt"
	"testing"
)

type AnyError struct {
	Any any `json:"fn"`
}

func (e *AnyError) Error() string {
	return "error"
}

func TestGenericError_Error(t *testing.T) {
	t.Parallel()

	t.Run("should return error message without original", func(t *testing.T) {
		want := "name: message"
		got := ToError(&GenericError{
			Name:    "name",
			Message: "message",
		}).Error()

		if got != want {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})

	t.Run("should return error message with original", func(t *testing.T) {
		want := "name: message (error)"
		got := ToError(&GenericError{
			Name:     "name",
			Message:  "message",
			Original: &AnyError{},
		})

		if got.Error() != want {
			t.Fatalf("\n got:  %v\n want: %v", got.Error(), want)
		}
	})
}

func TestGenericError_JSON(t *testing.T) {
	t.Parallel()

	t.Run("should return json error without original", func(t *testing.T) {
		want := `{"name":"error","message":"error"}`
		err := &GenericError{
			Name:    "error",
			Message: "error",
		}
		got := string(err.JSON())

		if got != want {
			t.Errorf("JSON() = %v, want %v", got, want)
		}
	})

	t.Run("should return json error when original is not marshalable", func(t *testing.T) {
		anyErr := &AnyError{
			Any: func() {},
		}

		want := fmt.Sprintf(`{"name":"error","message":"error","original":{"name":"error","message":"GenericError.Original is not marshallable (json: unsupported type: func()), original: %s"}}`, anyErr.Error())
		got := string(ToError(GenericError{
			Name:     "error",
			Message:  "error",
			Original: anyErr,
		}).JSON())

		if got != want {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})

	t.Run("should return json error when original contains a circular reference", func(t *testing.T) {
		anyErr := &AnyError{}
		anyErr.Any = anyErr

		want := fmt.Sprintf(`{"name":"error","message":"error","original":{"name":"error","message":"GenericError.Original contains a circular reference (cycle found: [Original Any]), original: %s"}}`, anyErr.Error())
		got := string(ToError(GenericError{
			Name:     "error",
			Message:  "error",
			Original: anyErr,
		}).JSON())

		if got != want {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestWrap(t *testing.T) {
	t.Parallel()

	t.Run("should wrap an error", func(t *testing.T) {
		toWrap := errors.New("normal error")
		want := string(GenericError{
			Name:     "error",
			Message:  toWrap.Error(),
			Original: toWrap,
		}.JSON())
		got := string((Wrap(toWrap)).(*GenericError).JSON())
		if got != want {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})

}

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		err := GenericError{
			Name:    "error",
			Message: "message",
		}
		want := string(err.JSON())
		got, ok := New(err.Message).(*GenericError)
		if !ok {
			t.Fatalf("expected error to be of type GenericError, got %T", got)
		}

		if want != string(got.JSON()) {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestNewWithName(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		err := GenericError{
			Name:    "name",
			Message: "message",
		}
		got, ok := NewWithName(err.Name, err.Message).(*GenericError)
		if !ok {
			t.Fatalf("expected error to be of type GenericError, got %T", got)
		}
		want := string(err.JSON())

		if want != string(got.JSON()) {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}

func TestNewWithNameAndErr(t *testing.T) {
	t.Parallel()

	t.Run("should create a new error", func(t *testing.T) {
		want := &GenericError{
			Name:     "name",
			Message:  "message",
			Original: errors.New("original error"),
		}
		got, ok := NewWithNameAndErr(want.Name, want.Message, want.Original).(*GenericError)
		if !ok {
			t.Fatalf("expected error to be of type GenericError, got %T", got)
		}

		if string(want.JSON()) != string(got.JSON()) {
			t.Fatalf("\n got:  %v\n want: %v", got, want)
		}
	})
}
