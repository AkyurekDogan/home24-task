package httpclient

import (
	"net/http"
)

// Doer defines the minimal interface needed to execute an HTTP request.
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}
