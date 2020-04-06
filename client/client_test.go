package client_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
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
	wd, err := os.Getwd()
	filePath := path.Join(wd, "..", "customer", "testdata", "addcustomer.json")
	b, err := ioutil.ReadFile(filePath)
	var c customer.Customer
	err = json.Unmarshal([]byte(b), &c)
	if err != nil {
		log.Fatal(err)
	}
	payload := &customer.Payload{Customer: &c}

	// test data
	filePath = path.Join(wd, "..", "customer", "testdata", "addcustomer_res.json")
	b, err = ioutil.ReadFile(filePath)
	r := ioutil.NopCloser(bytes.NewReader([]byte(b)))
	mockHTTPClient := &mock.HTTPClient{}
	mockHTTPClient.DoFn = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 201,
			Body:       r,
		}, nil
	}
	s.Client.HTTPClient = mockHTTPClient

	b, err = json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	result, err := s.Client.Post("/customers.json", b)

	// assertion
	customerResult := result["customer"].(map[string]interface{})
	assert.Equal("Steve", customerResult["first_name"])
	assert.Nil(err)
}

func TestShopifyClientPostInvalid(t *testing.T) {
	assert := assert.New(t)

	s := client.New("key", "secret", "test")

	// test data
	wd, err := os.Getwd()
	filePath := path.Join(wd, "..", "customer", "testdata", "addcustomer_invalid.json")
	b, err := ioutil.ReadFile(filePath)
	var c customer.Customer
	err = json.Unmarshal([]byte(b), &c)
	if err != nil {
		log.Fatal(err)
	}
	payload := &customer.Payload{Customer: &c}

	filePath = path.Join(wd, "..", "customer", "testdata", "addcustomer_res_failed.json")
	b, err = ioutil.ReadFile(filePath)
	r := ioutil.NopCloser(bytes.NewReader([]byte(b)))
	mockHTTPClient := &mock.HTTPClient{}
	mockHTTPClient.DoFn = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 201,
			Body:       r,
		}, nil
	}
	s.Client.HTTPClient = mockHTTPClient

	b, err = json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	result, err := s.Client.Post("/customers.json", b)

	errors := result["errors"].(map[string]interface{})
	base := errors["base"].([]interface{})
	assert.Equal("Customer must have a name, phone number or email address", base[0])
}

func TestShopifyClientPostFailed(t *testing.T) {
	assert := assert.New(t)
	s := client.New("key", "secret", "test")
	mockHTTPClient := &mock.HTTPClient{}
	mockHTTPClient.DoFn = func(*http.Request) (*http.Response, error) {
		return nil, errors.New("Connection failed")
	}
	s.Client.HTTPClient = mockHTTPClient
	result, err := s.Client.Post("/customers.json", nil)
	assert.Nil(result)
	assert.Equal("Connection failed", err.Error())

	result, err = s.Client.Get("/customers.json")
	assert.Nil(result)
	assert.Equal("Connection failed", err.Error())
}

func TestShopifyClientPostToVault(t *testing.T) {
	assert := assert.New(t)

	s := client.New("key", "secret", "test")

	// test data
	wd, err := os.Getwd()
	filePath := path.Join(wd, "..", "saleschannel", "payment", "testdata", "req_creditcard.json")
	b, err := ioutil.ReadFile(filePath)
	var c customer.Customer
	err = json.Unmarshal([]byte(b), &c)
	if err != nil {
		log.Fatal(err)
	}
	payload := &customer.Payload{Customer: &c}

	// test data
	filePath = path.Join(wd, "..", "saleschannel", "payment", "testdata", "res_token.json")
	b, err = ioutil.ReadFile(filePath)
	r := ioutil.NopCloser(bytes.NewReader([]byte(b)))
	mockHTTPClient := &mock.HTTPClient{}
	mockHTTPClient.DoFn = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	s.Client.HTTPClient = mockHTTPClient

	b, err = json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	result, err := s.Client.PostToVault(b)

	// assertion
	assert.Equal("83hru3obno3hu434b3u", result["id"])
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
