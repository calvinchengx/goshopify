package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const vaultURL = "https://elb.deposit.shopifycs.com/sessions"

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
	HTTPClient := &http.Client{}
	VaultURL := vaultURL
	Client := &Client{APIKey, APISecret, BaseURL, VaultURL, HTTPClient}
	return &Shopify{domain, apiVersion, Client}
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client represents
type Client struct {
	APIKey     string
	APISecret  string
	BaseURL    string
	VaultURL   string
	HTTPClient HTTPClient
}

// PostToVault posts to credit card vault
func (c *Client) PostToVault(b []byte) (map[string]interface{}, error) {
	url := c.VaultURL

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.SetBasicAuth(c.APIKey, c.APISecret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

// Post executes a HTTP POST request with a json []byte payload to target resource and returns a response
// e.g. resource = "/customers.json"
func (c *Client) Post(resource string, b []byte) (map[string]interface{}, error) {
	url := c.BaseURL + resource

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.SetBasicAuth(c.APIKey, c.APISecret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

// Get executes a HTTP GET request with a json []byte payload to target resource and returns a response
func (c *Client) Get(resource string) (map[string]interface{}, error) {
	url := c.BaseURL + resource

	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(c.APIKey, c.APISecret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}
