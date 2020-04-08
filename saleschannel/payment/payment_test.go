package payment_test

import (
	"os"
	"path"
	"testing"

	"github.com/calvinchengx/goshopify/client"
	"github.com/calvinchengx/goshopify/saleschannel/payment"
	"github.com/calvinchengx/goshopify/testhelper"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestShopifyStoreCreditCard(t *testing.T) {
	assert := assert.New(t)

	shopifyClient := client.New("key", "secret", "test")
	svc := payment.New(shopifyClient)

	wd, _ := os.Getwd()

	// test data input
	filePath := path.Join(wd, "testdata", "req_creditcard.json")
	b, input := testhelper.ParseFile(filePath)

	// expected result
	filePath = path.Join(wd, "testdata", "res_sessionid.json")
	b, expected := testhelper.ParseFile(filePath)

	// prepare our mock http client
	svc.Shopify.Client.HTTPClient = testhelper.CreateMockHTTPClient(b, 200)

	var creditCard payment.CreditCard
	mapstructure.Decode(input, &creditCard)
	actual, err := svc.Store(&creditCard)

	assert.Equal(expected["id"], actual["id"])
	assert.Nil(err)
}

func TestShopifyCreatePaymentValid(t *testing.T) {
	assert := assert.New(t)

	shopifyClient := client.New("key", "secret", "test")
	svc := payment.New(shopifyClient)

	wd, _ := os.Getwd()

	// test data input
	filePath := path.Join(wd, "testdata", "req_payment_valid.json")
	b, input := testhelper.ParseFile(filePath)

	// expected result
	filePath = path.Join(wd, "testdata", "res_payment_valid.json")
	b, expected := testhelper.ParseFile(filePath)
	expectedPayment := expected["payment"].(map[string]interface{})
	expectedID := expectedPayment["id"].(float64)

	// prepare our mock http client
	svc.Shopify.Client.HTTPClient = testhelper.CreateMockHTTPClient(b, 202)

	// invoke our CreatePayment function
	var p payment.Payment
	mapstructure.Decode(input, &p)
	actual, err := svc.CreatePayment("somevalidtoken", &p)

	actualPayment := actual["payment"].(map[string]interface{})
	actualID := actualPayment["id"].(float64)

	assert.Equal(expectedPayment, actualPayment)
	assert.Equal(int(expectedID), int(actualID))
	assert.Nil(err)
}
