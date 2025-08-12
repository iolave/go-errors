package errors

import (
	"errors"
	"testing"
)

// DummyError is a dummy error
type DummyError struct{}

func (e DummyError) Error() string {
	return "dummy error"
}

func TestToError(t *testing.T) {
	t.Parallel()

	t.Run("should panic when error is not an Error", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("ToError() should panic when error is not an Error")
			}
		}()

		err := errors.New("error")
		if err := ToError(err); err != nil {
			t.Fatalf("ToError() = %v, want panic", err)
		}
	})
}
