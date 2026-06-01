package httperr

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const legacyResponseErrorFormat = `{"code": %d, "error": "%s"}`

func writeLegacyNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = fmt.Fprintf(w, legacyResponseErrorFormat, http.StatusNotFound, ErrNotFound)
}

func BenchmarkNotFoundStatic(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		NotFound(rec)
	}
}

func BenchmarkNotFoundLegacySprintf(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		writeLegacyNotFound(rec)
	}
}

func BenchmarkBadRequestDynamic(b *testing.B) {
	errMsg := `invalid "name" field`
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		BadRequest(rec, benchmarkError{msg: errMsg})
	}
}

func BenchmarkBadRequestLegacySprintf(b *testing.B) {
	errMsg := `invalid "name" field`
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		rec.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(rec, legacyResponseErrorFormat, http.StatusBadRequest, benchmarkError{msg: errMsg})
	}
}

type benchmarkError struct {
	msg string
}

func (e benchmarkError) Error() string {
	return e.msg
}
