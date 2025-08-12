package errors

// Error is an interface for errors
// that can be marshalled to JSON
type Error interface {
	// ensures that the error implements the error interface
	error

	// JSON returns the JSON representation of the error
	// and it must ensure that the returned value is
	// is a valid JSON even if the error is not marshallable
	// by handling the error returned by json.Marshal
	JSON() []byte
}

// ToError asserts that the error is an Error
// and returns it. Otherwise, it panics.
//
// This function is useful when you want to
// ensure that the error is an Error or you
// know beforehand that the error is an Error
// and you want to use Error methods.
func ToError(err error) Error {
	e, ok := err.(Error)
	if !ok {
		panic("given err is not an Error")
	}

	return e
}
