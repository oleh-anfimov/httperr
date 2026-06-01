package httperr

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteJSONErrorContentType(t *testing.T) {
	rec := httptest.NewRecorder()
	BadRequest(rec, errors.New("bad input"))

	if ct := rec.Header().Get("Content-Type"); ct != contentTypeJSON {
		t.Fatalf("Content-Type = %q, want %q", ct, contentTypeJSON)
	}
}

func TestWriteJSONErrorEscapesMessage(t *testing.T) {
	rec := httptest.NewRecorder()
	BadRequest(rec, errors.New(`invalid "name" field`))

	var resp errorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON body: %v", err)
	}
	if resp.Code != http.StatusBadRequest {
		t.Fatalf("code = %d, want %d", resp.Code, http.StatusBadRequest)
	}
	if resp.Error != `invalid "name" field` {
		t.Fatalf("error = %q, want quoted field message", resp.Error)
	}
}

func TestWriteJSONErrorNilError(t *testing.T) {
	rec := httptest.NewRecorder()
	BadRequest(rec, nil)

	var resp errorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON body: %v", err)
	}
	if resp.Error != "" {
		t.Fatalf("error = %q, want empty string", resp.Error)
	}
}

func TestStaticResponses(t *testing.T) {
	tests := []struct {
		name   string
		fn     func(http.ResponseWriter)
		status int
		errMsg string
	}{
		{"InternalServerError", InternalServerError, http.StatusInternalServerError, ErrInternalServer.Error()},
		{"NotImplemented", NotImplemented, http.StatusNotImplemented, ErrNotImplemented.Error()},
		{"BadGateway", BadGateway, http.StatusBadGateway, ErrBadGateway.Error()},
		{"ServiceUnavailable", ServiceUnavailable, http.StatusServiceUnavailable, ErrServiceUnavailable.Error()},
		{"GatewayTimeout", GatewayTimeout, http.StatusGatewayTimeout, ErrGatewayTimeout.Error()},
		{"HTTPVersionNotSupported", HTTPVersionNotSupported, http.StatusHTTPVersionNotSupported, ErrHTTPVersionNotSupported.Error()},
		{"NotFound", NotFound, http.StatusNotFound, ErrNotFound.Error()},
		{"RequestTimeout", RequestTimeout, http.StatusRequestTimeout, ErrRequestTimeout.Error()},
		{"TooManyRequests", TooManyRequests, http.StatusTooManyRequests, ErrTooManyRequests.Error()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			tt.fn(rec)

			if rec.Code != tt.status {
				t.Fatalf("status = %d, want %d", rec.Code, tt.status)
			}
			if ct := rec.Header().Get("Content-Type"); ct != contentTypeJSON {
				t.Fatalf("Content-Type = %q, want %q", ct, contentTypeJSON)
			}

			var resp errorResponse
			if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
				t.Fatalf("invalid JSON body: %v", err)
			}
			if resp.Code != tt.status {
				t.Fatalf("code = %d, want %d", resp.Code, tt.status)
			}
			if resp.Error != tt.errMsg {
				t.Fatalf("error = %q, want %q", resp.Error, tt.errMsg)
			}
		})
	}
}

func TestNewWrappers(t *testing.T) {
	tests := []struct {
		name   string
		fn     func(http.ResponseWriter)
		status int
	}{
		{"Unauthorized", func(w http.ResponseWriter) { Unauthorized(w, ErrUnauthorized) }, http.StatusUnauthorized},
		{"Conflict", func(w http.ResponseWriter) { Conflict(w, ErrConflict) }, http.StatusConflict},
		{"UnprocessableEntity", func(w http.ResponseWriter) { UnprocessableEntity(w, ErrUnprocessable) }, http.StatusUnprocessableEntity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			tt.fn(rec)

			if rec.Code != tt.status {
				t.Fatalf("status = %d, want %d", rec.Code, tt.status)
			}
			if !strings.Contains(rec.Body.String(), `"code":`) {
				t.Fatalf("body %q missing code field", rec.Body.String())
			}
		})
	}
}

func TestErrResponse(t *testing.T) {
	rec := httptest.NewRecorder()
	ErrResponse(rec, http.StatusTeapot, errors.New("short and stout"))

	if rec.Code != http.StatusTeapot {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusTeapot)
	}

	var resp errorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON body: %v", err)
	}
	if resp.Code != http.StatusTeapot {
		t.Fatalf("code = %d, want %d", resp.Code, http.StatusTeapot)
	}
	if resp.Error != "short and stout" {
		t.Fatalf("error = %q, want %q", resp.Error, "short and stout")
	}
}
