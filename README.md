[![Maintainability](https://api.codeclimate.com/v1/badges/3929911652ee2902830b/maintainability)](https://codeclimate.com/github/calvinchengx/goshopify/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/3929911652ee2902830b/test_coverage)](https://codeclimate.com/github/calvinchengx/goshopify/test_coverage)

# golang SDK for Shopify Admin APIs

A golang library that interacts with Shopify Admin REST APIs.

## Tests and coverage

```bash
go test -coverprofile c.out ./...
go tool cover -html=c.out

# or simply
./test.sh
```