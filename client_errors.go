package httperr

import (
	"encoding/json"
	"net/http"
)

const contentTypeJSON = "application/json"

type errorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

var (
	notFoundBody        []byte
	requestTimeoutBody  []byte
	tooManyRequestsBody []byte
)

func init() {
	notFoundBody = mustMarshal(http.StatusNotFound, ErrNotFound.Error())
	requestTimeoutBody = mustMarshal(http.StatusRequestTimeout, ErrRequestTimeout.Error())
	tooManyRequestsBody = mustMarshal(http.StatusTooManyRequests, ErrTooManyRequests.Error())
}

func mustMarshal(code int, msg string) []byte {
	body, err := json.Marshal(errorResponse{Code: code, Error: msg})
	if err != nil {
		panic(err)
	}
	return body
}

func writeStaticJSON(w http.ResponseWriter, code int, body []byte) {
	w.Header().Set("Content-Type", contentTypeJSON)
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func writeJSONError(w http.ResponseWriter, code int, msg string) {
	body, err := json.Marshal(errorResponse{Code: code, Error: msg})
	if err != nil {
		writeStaticJSON(w, http.StatusInternalServerError, internalServerErrorBody)
		return
	}
	writeStaticJSON(w, code, body)
}

func errorMessage(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

// ErrResponse writes a JSON error response with the given status code and
// message from err. If err is nil, the error field is empty.
func ErrResponse(w http.ResponseWriter, code int, err error) {
	writeJSONError(w, code, errorMessage(err))
}

// BadRequest writes a 400 Bad Request JSON error response using err.Error() as
// the message. If err is nil, the error field is empty.
func BadRequest(w http.ResponseWriter, err error) {
	writeJSONError(w, http.StatusBadRequest, errorMessage(err))
}

// NotFound writes a 404 Not Found JSON error response with ErrNotFound.
func NotFound(w http.ResponseWriter) {
	writeStaticJSON(w, http.StatusNotFound, notFoundBody)
}

// Forbidden writes a 403 Forbidden JSON error response using err.Error() as the
// message. If err is nil, the error field is empty.
func Forbidden(w http.ResponseWriter, err error) {
	writeJSONError(w, http.StatusForbidden, errorMessage(err))
}

// RequestTimeout writes a 408 Request Timeout JSON error response with
// ErrRequestTimeout.
func RequestTimeout(w http.ResponseWriter) {
	writeStaticJSON(w, http.StatusRequestTimeout, requestTimeoutBody)
}

// Unauthorized writes a 401 Unauthorized JSON error response using err.Error()
// as the message. If err is nil, the error field is empty.
func Unauthorized(w http.ResponseWriter, err error) {
	writeJSONError(w, http.StatusUnauthorized, errorMessage(err))
}

// Conflict writes a 409 Conflict JSON error response using err.Error() as the
// message. If err is nil, the error field is empty.
func Conflict(w http.ResponseWriter, err error) {
	writeJSONError(w, http.StatusConflict, errorMessage(err))
}

// UnprocessableEntity writes a 422 Unprocessable Entity JSON error response
// using err.Error() as the message. If err is nil, the error field is empty.
func UnprocessableEntity(w http.ResponseWriter, err error) {
	writeJSONError(w, http.StatusUnprocessableEntity, errorMessage(err))
}

// TooManyRequests writes a 429 Too Many Requests JSON error response with
// ErrTooManyRequests.
func TooManyRequests(w http.ResponseWriter) {
	writeStaticJSON(w, http.StatusTooManyRequests, tooManyRequestsBody)
}
