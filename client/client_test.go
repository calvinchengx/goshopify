package client_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/calvinchengx/goshopify/client"
	"github.com/calvinchengx/goshopify/customer"
	"github.com/calvinchengx/goshopify/mock"
	"github.com/stretchr/testify/assert"
)

func TestShopifyNewClient(t *testing.T) {
	assert := assert.New(t)

	s := client.New("key", "secret", "test")

	assert.NotNil(s)

	expected := fmt.Sprintf(`https://%s.myshopify.com/admin/api/%s`, "test", "2020-04")

	assert.Equal(expected, s.Client.BaseURL)

}

func TestShopifyClientPost(t *testing.T) {
	assert := assert.New(t)

	s := client.New("key", "secret", "test")

	// test data
	c := &customer.Customer{
		Email:     "testemail@example.com",
		FirstName: "John",
		LastName:  "Freud",
	}
	payload := &customer.Payload{Customer: c}

	// test data
	jsonData := `{"id":"122345","first_name":"John","last_name":"Freud","email":"testemail@example.com"}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))
	mockHTTPClient := &mock.HTTPClient{}
	mockHTTPClient.DoFn = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			Body: r,
		}, nil
	}
	s.Client.HTTPClient = mockHTTPClient

	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	result, err := s.Client.Post("/customers.json", b)

	// assertion
	assert.Equal("John", result["first_name"])
	assert.Nil(err)
}

func TestShopifyClientGet(t *testing.T) {
	assert := assert.New(t)

	s := client.New("key", "secret", "test")

	// test data
	jsonData := `{"customers": [ {"id":"122345","first_name":"John","last_name":"Freud","email":"testemail@example.com"} ] }`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))
	mockHTTPClient := &mock.HTTPClient{}
	mockHTTPClient.DoFn = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			Body: r,
		}, nil
	}
	s.Client.HTTPClient = mockHTTPClient

	result, err := s.Client.Get("/customers.json")

	// assertion
	resultList := result["customers"].([]interface{})
	result0 := resultList[0].(map[string]interface{})
	assert.Equal("John", result0["first_name"].(string))
	assert.Nil(err)
}
