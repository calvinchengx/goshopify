package main

import (
	"log"
	"os"

	shopify "github.com/calvinchengx/goshopify"
	"github.com/calvinchengx/goshopify/customer"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("SHOPIFY_API_KEY")
	apiSecret := os.Getenv("SHOPIFY_API_SECRET")
	domain := os.Getenv("SHOPIFY_DOMAIN")
	s := shopify.NewService(apiKey, apiSecret, domain)
	customer := &customer.Customer{
		FirstName:     "Steve",
		LastName:      "Lastnameson",
		Email:         "steve.lastnameson@example.com",
		VerifiedEmail: true,
		Phone:         "+15142546011",
		Addresses: []*customer.Address{
			&customer.Address{
				Address1:  "123 Oak St",
				City:      "Ottawa",
				Province:  "ON",
				Phone:     "555-1212",
				Zip:       "123 ABC",
				LastName:  "Lastnameson",
				FirstName: "Mother",
				Country:   "CA",
			},
		},
	}
	s.CustomerSvc.Add(customer)
}
