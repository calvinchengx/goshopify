package client

import "fmt"

// Shopify models the shopify client parameters
// https://{apiKey}:{apiSecret}@{domain}.myshopify.com/admin/api/{api-version}/{resource}.json
type Shopify struct {
	APIKey     string
	APISecret  string
	domain     string
	apiVersion string
}

// New instantiates a shopify client
func New(APIKey, APISecret, domain string) *Shopify {
	// default version
	apiVersion := "2020-04"
	return &Shopify{APIKey, APISecret, domain, apiVersion}
}

// URL returns the Shopify Admin URL we are targeting
func (s *Shopify) URL() string {
	url := fmt.Sprintf(`https://%s.myshopify.com/admin/api/%s`, s.domain, s.apiVersion)
	fmt.Println(url)
	return url
}
