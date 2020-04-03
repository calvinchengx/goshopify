package mock

import "net/http"

// HTTPClient is a mock of http.Client
type HTTPClient struct {
	DoFn func(*http.Request) (*http.Response, error)
}

// Do mocks the http
func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFn(req)
}
