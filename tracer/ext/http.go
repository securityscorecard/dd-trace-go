package ext

// HTTP meta constants.
const (
	HTTPType   = "http"
	HTTPMethod = "http.method"
	HTTPCode   = "http.status_code"
	HTTPURL    = "http.url"
)

// Distributed tracing headers
const (
	HTTPTraceIDHeader  = "x-datadog-trace-id"
	HTTPParentIDHeader = "x-datadog-parent-id"
)
