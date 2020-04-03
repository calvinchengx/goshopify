package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Shopify models the shopify client parameters
// https://{apiKey}:{apiSecret}@{domain}.myshopify.com/admin/api/{api-version}/{resource}.json
type Shopify struct {
	domain     string
	apiVersion string
	Client     *Client
}

// New instantiates a shopify client
func New(APIKey, APISecret, domain string) *Shopify {
	// default version
	apiVersion := "2020-04"
	BaseURL := fmt.Sprintf(`https://%s.myshopify.com/admin/api/%s`, domain, apiVersion)
	Client := &Client{APIKey, APISecret, BaseURL}
	return &Shopify{domain, apiVersion, Client}
}

// Client represents
type Client struct {
	APIKey    string
	APISecret string
	BaseURL   string
}

// Post executes a HTTP POST request with a json []byte payload to target resource and returns a response
// e.g. resource = "/customers.json"
func (c *Client) Post(resource string, b []byte) map[string]interface{} {
	url := c.BaseURL + resource

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.SetBasicAuth(c.APIKey, c.APISecret)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}

// Get executes a HTTP GET request with a json []byte payload to target resource and returns a response
func (c *Client) Get(resource string) map[string]interface{} {
	url := c.BaseURL + resource

	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(c.APIKey, c.APISecret)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}
