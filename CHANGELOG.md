# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed
- `Wrap` panic when wrapping nil error.


## [v1.0.0]

### Added
- `Error` interface with `JSON()` method.
- `ToError` function to assert that an error implements `Error`.
- `GenericError` struct with `New`, `NewWithName`, `NewWithNameAndErr` constructors.
- `HTTPError` struct with the following constructors:
    - `NewHTTPError`,
    - `NewBadRequestError`,
    - `NewNotFoundError`,
    - `NewInternalServerError`,
    - `NewUnauthorizedError`,
    - `NewForbiddenError`,
    - `NewConflictError`,
    - `NewTooManyRequestsError`,
    - `NewBadGatewayError`,
    - `NewServiceUnavailableError`,
    - `NewGatewayTimeoutError`.

[unreleased]: https://github.com/iolave/go-proxmox/compare/v1.0.0...master
[v1.0.0]: https://github.com/iolave/go-errors/commits/v1.0.0

