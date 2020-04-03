package client_test

import (
	"fmt"
	"testing"

	"github.com/calvinchengx/goshopify/client"
	"github.com/stretchr/testify/assert"
)

func TestShopifyNewClient(t *testing.T) {
	assert := assert.New(t)

	s := client.New("key", "secret", "test")

	assert.NotNil(s)

	expected := fmt.Sprintf(`https://%s.myshopify.com/admin/api/%s`, "test", "2020-04")

	assert.Equal(expected, s.URL())
}
