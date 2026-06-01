package httperr

import "errors"

var (
	// ErrInternalServer is returned by InternalServerError.
	ErrInternalServer = errors.New("internal server error")

	// ErrNotFound is returned by NotFound.
	ErrNotFound = errors.New("not found")

	// ErrRequestTimeout is returned by RequestTimeout.
	ErrRequestTimeout = errors.New("request timeout")

	// ErrUnauthorized is the default message for Unauthorized responses.
	ErrUnauthorized = errors.New("unauthorized")

	// ErrConflict is the default message for Conflict responses.
	ErrConflict = errors.New("conflict")

	// ErrUnprocessable is the default message for UnprocessableEntity responses.
	ErrUnprocessable = errors.New("unprocessable entity")

	// ErrTooManyRequests is returned by TooManyRequests.
	ErrTooManyRequests = errors.New("too many requests")

	// ErrNotImplemented is returned by NotImplemented.
	ErrNotImplemented = errors.New("not implemented")

	// ErrBadGateway is returned by BadGateway.
	ErrBadGateway = errors.New("bad gateway")

	// ErrServiceUnavailable is returned by ServiceUnavailable.
	ErrServiceUnavailable = errors.New("service unavailable")

	// ErrGatewayTimeout is returned by GatewayTimeout.
	ErrGatewayTimeout = errors.New("gateway timeout")

	// ErrHTTPVersionNotSupported is returned by HTTPVersionNotSupported.
	ErrHTTPVersionNotSupported = errors.New("http version not supported")
)
