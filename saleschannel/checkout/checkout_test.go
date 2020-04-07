package checkout_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/calvinchengx/goshopify/client"
	"github.com/calvinchengx/goshopify/mock"
	"github.com/calvinchengx/goshopify/saleschannel/checkout"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestShopifyCheckoutCreateInvalid(t *testing.T) {
	assert := assert.New(t)

	shopifyClient := client.New("key", "secret", "test")
	svc := checkout.New(shopifyClient)

	// test data
	wd, err := os.Getwd()
	filePath := path.Join(wd, "testdata", "req_checkout_invalid.json")
	b, err := ioutil.ReadFile(filePath)
	var checkoutReq checkout.Checkout
	err = json.Unmarshal([]byte(b), &checkoutReq)
	if err != nil {
		log.Fatal(err)
	}

	// test data
	filePath = path.Join(wd, "testdata", "res_checkout_invalid.json")
	b, err = ioutil.ReadFile(filePath)
	var checkoutRes map[string]interface{}
	err = json.Unmarshal([]byte(b), &checkoutRes)

	// prepare our mock
	r := ioutil.NopCloser(bytes.NewReader([]byte(b)))
	mockHTTPClient := &mock.HTTPClient{}
	mockHTTPClient.DoFn = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 422,
			Body:       r,
		}, nil
	}
	svc.Shopify.Client.HTTPClient = mockHTTPClient

	result, err := svc.Create(&checkoutReq)
	assert.Equal(checkoutRes, result)
	assert.Nil(err)
}

func TestShopifyCheckoutCreateValid(t *testing.T) {
	assert := assert.New(t)

	shopifyClient := client.New("key", "secret", "test")
	svc := checkout.New(shopifyClient)

	// test data
	wd, err := os.Getwd()
	filePath := path.Join(wd, "testdata", "req_checkout_valid_nolineitems.json")
	b, err := ioutil.ReadFile(filePath)
	var checkoutReq checkout.Checkout
	err = json.Unmarshal([]byte(b), &checkoutReq)
	if err != nil {
		log.Fatal(err)
	}

	// test data
	filePath = path.Join(wd, "testdata", "res_checkout_valid_created.json")
	b, err = ioutil.ReadFile(filePath)
	var checkoutRes map[string]interface{}
	err = json.Unmarshal([]byte(b), &checkoutRes)

	// prepare our mock
	r := ioutil.NopCloser(bytes.NewReader([]byte(b)))
	mockHTTPClient := &mock.HTTPClient{}
	mockHTTPClient.DoFn = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 201,
			Body:       r,
		}, nil
	}
	svc.Shopify.Client.HTTPClient = mockHTTPClient

	result, err := svc.Create(&checkoutReq)
	assert.Equal(checkoutRes, result)
	assert.Nil(err)
}

func TestShopifyCheckoutCreateValidWithLineItems(t *testing.T) {
	assert := assert.New(t)

	shopifyClient := client.New("key", "secret", "test")
	svc := checkout.New(shopifyClient)

	// test data
	wd, err := os.Getwd()
	filePath := path.Join(wd, "testdata", "req_checkout_valid_withlineitems.json")
	b, err := ioutil.ReadFile(filePath)
	var checkoutReq map[string]interface{}
	err = json.Unmarshal([]byte(b), &checkoutReq)

	// test data
	filePath = path.Join(wd, "testdata", "res_checkout_valid_created_withlineitems.json")
	b, err = ioutil.ReadFile(filePath)
	var checkoutRes map[string]interface{}
	err = json.Unmarshal([]byte(b), &checkoutRes)

	// prepare our mock
	r := ioutil.NopCloser(bytes.NewReader([]byte(b)))
	mockHTTPClient := &mock.HTTPClient{}
	mockHTTPClient.DoFn = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 201,
			Body:       r,
		}, nil
	}
	svc.Shopify.Client.HTTPClient = mockHTTPClient

	var checkoutRequest *checkout.Checkout
	mapstructure.Decode([]byte(b), checkoutRequest)
	result, err := svc.Create(checkoutRequest)
	assert.Equal(checkoutRes, result)
	checkout := checkoutRes["checkout"].(map[string]interface{})
	// this is the token used for
	token := checkout["token"]
	assert.Equal("660b5050744ca776869234e2c54e6133", token)
	assert.Nil(err)
}

func TestShopifyCheckoutComplete(t *testing.T) {
	assert := assert.New(t)

	shopifyClient := client.New("key", "secret", "test")
	svc := checkout.New(shopifyClient)

	// test data
	wd, _ := os.Getwd()

	// test data
	filePath := path.Join(wd, "testdata", "res_checkout_complete.json")
	b, _ := ioutil.ReadFile(filePath)
	var checkoutRes map[string]interface{}
	_ = json.Unmarshal([]byte(b), &checkoutRes)

	// prepare our mock
	r := ioutil.NopCloser(bytes.NewReader([]byte(b)))
	mockHTTPClient := &mock.HTTPClient{}
	mockHTTPClient.DoFn = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 202,
			Body:       r,
		}, nil
	}
	svc.Shopify.Client.HTTPClient = mockHTTPClient

	result, err := svc.Complete("sometoken")
	assert.NotNil(result)
	assert.Nil(err)
}
