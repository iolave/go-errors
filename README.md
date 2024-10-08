# Errors package for golang

> [!WARNING]
> All versions released prior to `v1.0.0` are to be considered [breaking changes](https://semver.org/#how-do-i-know-when-to-release-100)

Yet another `errors` package, implementing error handling primitives with error wrapping and error tracing.

## Installation

The recommended way to install go-errors is:

```
go get -u github.com/iolave/go-errors
```

## Examples

The `Kind` type allows you to create new errors containing the stack trace and also check if an error is of a particular kind.

```go
var ErrExample = errors.NewKind("example")

err := ErrExample.New()
if ErrExample.Is(err) {
	fmt.Printf("%+v\n", err)
}

// Example Output:
// example
//
// gopkg.in/src-d/errors%2v0_test.ExampleError_Format
//         /home/mcuadros/workspace/go/src/gopkg.in/src-d/errors.v0/example_test.go:60
// testing.runExample
//         /usr/lib/go/src/testing/example.go:114
// testing.RunExamples
//         /usr/lib/go/src/testing/example.go:38
// testing.(*M).Run
//         /usr/lib/go/src/testing/testing.go:744
// main.main
//         github.com/pkg/errors/_test/_testmain.go:106
// runtime.main
//         /usr/lib/go/src/runtime/proc.go:183
// runtime.goexit
//         /usr/lib/go/src/runtime/asm_amd64.s:2086
```


### Error with format

```go
var ErrMaxLimitReached = errors.NewKind("max. limit reached: %d")

err := ErrMaxLimitReached.New(42)
if ErrMaxLimitReached.Is(err) {
    fmt.Println(err)
}

// Output: max. limit reached: 42
```


### Error wrapping

```go
var ErrNetworking = errors.NewKind("network error")

err := ErrNetworking.Wrap(io.EOF)
if ErrNetworking.Is(err) {
    fmt.Println(err)
}

// Output: network error: EOF
```

You can find these examples and many more in the [examples](example_test.go) file.


## License

Apache License 2.0, see [LICENSE](LICENSE)
