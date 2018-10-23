# Simple Payments API
API is meant to be a part of bigger platform behind auth gateway

## Technical Choices
* Web Framework
  `go-swagger` provides api validation and error handling based API swagger specification
* Storage
  `cassandra` platform choice

## API Spec
api/swagger.yml
## Dev
There are separate cassandra keyspaces for dev and test envs
### Build
`make build`
Not all build requirements are coverd in `Makefile` rules
### Test
* Unit tests are not forced to be atomic. One has to take care of it on its own running `clearStorage()`
* Request to PaymentsSource is not mocked.
