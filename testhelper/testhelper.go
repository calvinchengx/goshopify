package testhelper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/calvinchengx/goshopify/mock"
)

// ParseFile parses the given json file and returns its bytes and map
func ParseFile(filePath string) (b []byte, result *map[string]interface{}) {
	b, _ = ioutil.ReadFile(filePath)
	_ = json.Unmarshal([]byte(b), result)
	return b, result
}

// CreateMockHTTPClient creates a mock http client
func CreateMockHTTPClient(b []byte, statusCode int) *mock.HTTPClient {
	r := ioutil.NopCloser(bytes.NewReader([]byte(b)))
	mockHTTPClient := &mock.HTTPClient{}
	mockHTTPClient.DoFn = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: statusCode,
			Body:       r,
		}, nil
	}
	return mockHTTPClient
}
