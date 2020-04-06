package payment_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/calvinchengx/goshopify/client"
	"github.com/calvinchengx/goshopify/mock"
	"github.com/calvinchengx/goshopify/saleschannel/payment"
	"github.com/stretchr/testify/assert"
)

func TestShopifyStoreCreditCard(t *testing.T) {
	assert := assert.New(t)

	shopifyClient := client.New("key", "secret", "test")
	svc := payment.New(shopifyClient)

	// test data
	wd, _ := os.Getwd()

	// test data
	filePath := path.Join(wd, "testdata", "req_creditcard.json")
	b, _ := ioutil.ReadFile(filePath)
	var creditCard payment.CreditCard
	_ = json.Unmarshal([]byte(b), &creditCard)

	// test data
	filePath = path.Join(wd, "testdata", "res_token.json")
	b, _ = ioutil.ReadFile(filePath)
	var tokenRes map[string]interface{}
	_ = json.Unmarshal([]byte(b), &tokenRes)

	// prepare our mock
	r := ioutil.NopCloser(bytes.NewReader([]byte(b)))
	mockHTTPClient := &mock.HTTPClient{}
	mockHTTPClient.DoFn = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	svc.Shopify.Client.HTTPClient = mockHTTPClient

	result, err := svc.Store(&creditCard)
	assert.Equal("83hru3obno3hu434b3u", result["id"])
	assert.Nil(err)
}
