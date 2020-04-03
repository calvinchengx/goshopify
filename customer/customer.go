package customer

import (
	"encoding/json"
	"log"

	"github.com/calvinchengx/goshopify/client"
)

// New returns our customer's service
func New(Shopify *client.Shopify) *Service {
	return &Service{Shopify}
}

// Service is the implementation to access our service
type Service struct {
	Shopify *client.Shopify
}

type customerPayLoad struct {
	Customer *Customer `json:"customer"`
}

// Customer models a shopify customer
type Customer struct {
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	Email         string     `json:"email"`
	Phone         string     `json:"phone"`
	VerifiedEmail bool       `json:"verified_email"`
	Addresses     []*Address `json:"addresses"`
}

// Address models a shopify customer's address
type Address struct {
	Address1  string `json:"address1"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Phone     string `json:"phone"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Add makes a POST request to the customers.json url to create a customer
func (s *Service) Add(Customer *Customer) {
	payload := &customerPayLoad{Customer}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}
	s.Shopify.Client.Post("/customers.json", b)
}
