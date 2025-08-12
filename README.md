# go-errors

A simple error handling library for Go, designed for clean JSON serialization, safe error wrapping, and HTTP-friendly error responses.

## ✨ Features

- **Unified `Error` interface** — all errors can be serialized to valid JSON.
- **Generic errors** with names, messages, and original error support.
- **HTTP errors** with built-in status codes and structured responses.
- **Circular reference protection** when marshalling to JSON.
- **Automatic wrapping** of non-JSON-compatible errors.
- **Simple, idiomatic API** with no external dependencies (except for circular reference checking).

---

## 📦 Installation

```bash
go get github.com/iolave/go-errors
```

---

## 🚀 Quick Start

### Creating a generic error

```go
import errors "github.com/iolave/go-errors"

func main() {
    err := errors.New("configuration file missing").(errors.Error)
    fmt.Println(err.JSON())
    // {"name":"error","message":"configuration file missing"}
}
```

### Wrapping an existing error

```go
import (
    "fmt"
    errors "github.com/iolave/go-errors"
)

func main() {
    err := fmt.Errorf("configuration file missing")
    fmt.Println(errors.Wrap(err).(errors.Error).JSON())
    // {"name":"error","message":"configuration file missing","original": {}}
}
```

### Named errors

```go
import errors "github.com/iolave/go-errors"

func main() {
    err := errors.NewWithName("my_error", "configuration file missing").(errors.Error)
    fmt.Println(err.JSON())
    // {"name":"my_error","message":"configuration file missing"}
}
```

### HTTP errors

```go
import errors "github.com/iolave/go-errors"

func main() {
    err := errors.NewNotFoundError("user not found", nil).(errors.Error)
    fmt.Println(err.JSON())
    // {"statusCode":404,"name":"not_found_error","message":"user not found","error":null}
}
```

Available constructors:

- `NewBadRequestError(message string, err error) error`
- `NewNotFoundError(message string, err error) error`
- `NewInternalServerError(message string, err error) error`
- `NewUnauthorizedError(message string, err error) error`
- `NewForbiddenError(message string, err error) error`
- `NewConflictError(message string, err error) error`
- `NewTooManyRequestsError(message string, err error) error`
- `NewBadGatewayError(message string, err error) error`
- `NewServiceUnavailableError(message string, err error) error`
- `NewGatewayTimeoutError(message string, err error) error`

Each has a corresponding `func(message string, err error) error` signature.

---

## 📜 API Overview

### `Error` interface

All errors implement:
```go
type Error interface {
    error
    JSON() []byte
}
```

### Conversion

- `ToError(err error) Error` — asserts that an error implements Error (panics otherwise).

### Generic Errors

- `New(msg string) error` 
- `NewWithName(name, msg string) error` 
- `NewWithNameAndErr(name, msg string, orig error) error` 
- `Wrap(err error) error`

### HTTP Errors

- `NewHTTPError(statusCode int, name, message string, err error) error`

Convenience constructors listed above.

---

### 🛡️ JSON Safety
`errors` automatically:

- Detects circular references and replaces them with explanatory placeholder errors.
- Wraps non-`Error` types so their content is still available in serialized form.
- Ensures `JSON()` always returns valid JSON.

### 🤝 Contributing

Pull requests are welcome!

If you add features or change APIs, please include tests and update the README accordingly.
