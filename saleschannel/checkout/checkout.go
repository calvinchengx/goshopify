package checkout

import (
	"encoding/json"
	"fmt"

	"github.com/calvinchengx/goshopify/client"
)

// New returns our checkout's service
func New(Shopify *client.Shopify) *Service {
	return &Service{Shopify}
}

// Service is the implementation to access our service
type Service struct {
	Shopify *client.Shopify
}

// Payload encapsulates our checkout and associated line items
type Payload struct {
	Checkout *Checkout `json:"checkout"`
}

// Checkout models a checkout object
type Checkout struct {
	LineItems []*LineItem `json:"line_items"`
}

// LineItem models a line_item object list
type LineItem struct {
	VariantID int64 `json:"variant_id"`
	Quantity  int64 `json:"quantity"`
}

// Create a checkout
func (s *Service) Create(Checkout *Checkout) (map[string]interface{}, error) {
	payload := &Payload{Checkout}
	b, _ := json.Marshal(payload)
	return s.Shopify.Client.Post("/checkouts.json", b)
}

// Complete a checkout using the token returned by payment.Store
func (s *Service) Complete(token string) (map[string]interface{}, error) {
	resource := fmt.Sprintf("/checkouts/%s/complete.json", token)
	return s.Shopify.Client.Post(resource, []byte(`{}`))
}
