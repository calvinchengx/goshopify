package shopify

import (
	"github.com/calvinchengx/goshopify/client"
	"github.com/calvinchengx/goshopify/customer"
)

// NewService instantiates all our services that make up our shopify service
func NewService(apiKey, apiSecret, domain string) *Service {
	c := client.New(apiKey, apiSecret, domain)
	CustomerSvc := customer.New(c)
	return &Service{CustomerSvc}
}

// Service is a definition of all our services
type Service struct {
	CustomerSvc *customer.Service
}
