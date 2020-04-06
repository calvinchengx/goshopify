package payment

import (
	"encoding/json"

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

// Payload encapsulates our credit card object
type Payload struct {
	CreditCard *CreditCard `json:"credit_card"`
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

// Store saves Credit Card to vault
func (s *Service) Store(CreditCard *CreditCard) (map[string]interface{}, error) {
	payload := &Payload{CreditCard}
	b, _ := json.Marshal(payload)
	return s.Shopify.Client.PostToVault(b)
}
