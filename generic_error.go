package errors

import (
	"encoding/json"
	"fmt"

	"github.com/theothertomelliott/acyclic"
)

var _ Error = &GenericError{}

type GenericError struct {
	// Name is the name of the error
	Name string `json:"name"`

	// Message is the message of the error
	Message string `json:"message"`

	// Original is an optional original error
	Original error `json:"original,omitempty"`
}

// Error returns a string concatenation of the name, message
// and original error.
func (e GenericError) Error() string {
	if e.Original == nil {
		return fmt.Sprintf("%s: %s", e.Name, e.Message)
	}

	return fmt.Sprintf("%s: %s (%s)", e.Name, e.Message, e.Original.Error())
}

// JSON returns the JSON representation of the error
//
// If the Orginal property is not marshallable, it will be
// replaced with a marshallable error indicating why the
// original error is not marshallable and it's original error
// returned by the Error() method.
func (e GenericError) JSON() []byte {
	if err := acyclic.Check(e); err != nil {
		e.Original = New(fmt.Sprintf(
			"GenericError.Original contains a circular reference (%s), original: %s",
			err.Error(),
			e.Original.Error(),
		))
	}

	b, err := json.Marshal(e)
	if err != nil {
		e.Original = New(fmt.Sprintf(
			"GenericError.Original is not marshallable (%s), original: %s",
			err.Error(),
			e.Original.Error(),
		))

		b, _ = json.Marshal(e)
	}

	return b
}

// Wrap wraps an error into a GenericError. It sets the name
// of the error to "error", the message to the original
// error.Error() value and the original property to the
// original error.
func Wrap(err error) error {
	return &GenericError{
		Name:     "error",
		Message:  err.Error(),
		Original: err,
	}
}

// New creates a new GenericError with it's
// name set to "error", the message set to
// the given message.
func New(msg string) error {
	return &GenericError{
		Name:    "error",
		Message: msg,
	}
}

// NewWithName creates a new GenericError with it's
// name set to the given name and the message set to
// the given message.
func NewWithName(name, msg string) error {
	return &GenericError{
		Name:    name,
		Message: msg,
	}
}

// NewWithNameAndErr creates a new GenericError with it's
// name set to the given name, the message set to
// the given message and the original property set to the given
// original error.
func NewWithNameAndErr(name, msg string, orig error) error {
	return &GenericError{
		Name:     name,
		Message:  msg,
		Original: orig,
	}
}
