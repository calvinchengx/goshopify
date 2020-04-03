package customer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
func (s *Service) Add(customer *Customer) {
	url := s.Shopify.URL() + "/customers.json"

	fmt.Println(customer)
	payload := make(map[string]*Customer)
	payload["customer"] = customer
	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.SetBasicAuth(s.Shopify.APIKey, s.Shopify.APISecret)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// bodyText, err := ioutil.ReadAll(resp.Body)
	// str := string(bodyText)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println(&result)
}
