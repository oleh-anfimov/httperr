package httperr

import "net/http"

var (
	internalServerErrorBody     []byte
	notImplementedBody          []byte
	badGatewayBody              []byte
	serviceUnavailableBody      []byte
	gatewayTimeoutBody          []byte
	httpVersionNotSupportedBody []byte
)

func init() {
	internalServerErrorBody = mustMarshal(http.StatusInternalServerError, ErrInternalServer.Error())
	notImplementedBody = mustMarshal(http.StatusNotImplemented, ErrNotImplemented.Error())
	badGatewayBody = mustMarshal(http.StatusBadGateway, ErrBadGateway.Error())
	serviceUnavailableBody = mustMarshal(http.StatusServiceUnavailable, ErrServiceUnavailable.Error())
	gatewayTimeoutBody = mustMarshal(http.StatusGatewayTimeout, ErrGatewayTimeout.Error())
	httpVersionNotSupportedBody = mustMarshal(http.StatusHTTPVersionNotSupported, ErrHTTPVersionNotSupported.Error())
}

// InternalServerError writes a 500 Internal Server Error JSON error response
// with ErrInternalServer.
func InternalServerError(w http.ResponseWriter) {
	writeStaticJSON(w, http.StatusInternalServerError, internalServerErrorBody)
}

// NotImplemented writes a 501 Not Implemented JSON error response with
// ErrNotImplemented.
func NotImplemented(w http.ResponseWriter) {
	writeStaticJSON(w, http.StatusNotImplemented, notImplementedBody)
}

// BadGateway writes a 502 Bad Gateway JSON error response with ErrBadGateway.
func BadGateway(w http.ResponseWriter) {
	writeStaticJSON(w, http.StatusBadGateway, badGatewayBody)
}

// ServiceUnavailable writes a 503 Service Unavailable JSON error response
// with ErrServiceUnavailable.
func ServiceUnavailable(w http.ResponseWriter) {
	writeStaticJSON(w, http.StatusServiceUnavailable, serviceUnavailableBody)
}

// GatewayTimeout writes a 504 Gateway Timeout JSON error response with
// ErrGatewayTimeout.
func GatewayTimeout(w http.ResponseWriter) {
	writeStaticJSON(w, http.StatusGatewayTimeout, gatewayTimeoutBody)
}

// HTTPVersionNotSupported writes a 505 HTTP Version Not Supported JSON error
// response with ErrHTTPVersionNotSupported.
func HTTPVersionNotSupported(w http.ResponseWriter) {
	writeStaticJSON(w, http.StatusHTTPVersionNotSupported, httpVersionNotSupportedBody)
}
