package shopify

import (
	"github.com/calvinchengx/goshopify/client"
	"github.com/calvinchengx/goshopify/customer"
	"github.com/calvinchengx/goshopify/saleschannel/checkout"
	"github.com/calvinchengx/goshopify/saleschannel/payment"
)

// NewService instantiates all our services that make up our shopify service
func NewService(apiKey, apiSecret, domain string) *Service {
	c := client.New(apiKey, apiSecret, domain)
	CustomerSvc := customer.New(c)
	CheckoutSvc := checkout.New(c)
	PaymentSvc := payment.New(c)
	return &Service{CustomerSvc, CheckoutSvc, PaymentSvc}
}

// Service is a definition of all our services
type Service struct {
	CustomerSvc *customer.Service
	CheckoutSvc *checkout.Service
	PaymentSvc  *payment.Service
}
