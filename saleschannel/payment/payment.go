package payment

import (
	"encoding/json"
	"fmt"

	"github.com/calvinchengx/goshopify/client"
)

// New returns our payment service
func New(Shopify *client.Shopify) *Service {
	return &Service{Shopify}
}

// Service is the implementation to access our service
type Service struct {
	Shopify *client.Shopify
}

// CreditCard models the credit card object
type CreditCard struct {
	Number            string `json:"number"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Month             string `json:"month"`
	Year              string `json:"year"`
	VerificationValue string `json:"verififcation_value"`
}

// Payment models the payment details
type Payment struct {
	RequestDetails *RequestDetails `json:"request_details"`
	Amount         string          `json:"amount"`
	SessionID      string          `json:"session_id"`
	UniqueToken    string          `json:"unique_token"`
}

// RequestDetails models the http request details in the payment object
type RequestDetails struct {
	IPAddress      string `json:"ip_address"`
	AcceptLanguage string `json:"accept_language"`
	UserAgent      string `json:"user_agent"`
}

// Store saves Credit Card to vault and returns a session ID
func (s *Service) Store(cc *CreditCard) (map[string]interface{}, error) {
	payload := &map[string]*CreditCard{
		"credit_card": cc,
	}
	b, _ := json.Marshal(payload)
	return s.Shopify.Client.PostToVault(b)
}

// CreatePayment creates a payment on a checkout using the token by the checkout creation
func (s *Service) CreatePayment(token string, p *Payment) (map[string]interface{}, error) {
	payload := &map[string]*Payment{
		"payment": p,
	}
	b, _ := json.Marshal(payload)
	resource := fmt.Sprintf("/checkouts/%s/payments.json", token)
	return s.Shopify.Client.Post(resource, b)
}
