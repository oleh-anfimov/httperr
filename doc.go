// Package httperr provides helpers for writing JSON HTTP error responses.
//
// Each helper sets Content-Type to application/json, writes the HTTP status
// code, and sends a body of the form {"code": <status>, "error": "<message>"}.
package httperr
