package ext

import "net/textproto"

// HTTP meta constants.
const (
	HTTPType   = "http"
	HTTPMethod = "http.method"
	HTTPCode   = "http.status_code"
	HTTPURL    = "http.url"
)

// Distributed tracing headers
var (
	HTTPTraceIDHeader  = textproto.CanonicalMIMEHeaderKey("X-Datadog-Trace-Id")
	HTTPParentIDHeader = textproto.CanonicalMIMEHeaderKey("X-Datadog-Parent-Id")
)
