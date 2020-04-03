package client

// IClient is the interface to access our http client
type IClient interface {
	Post(resource string, b []byte) map[string]interface{}
	Get(resource string) map[string]interface{}
}
