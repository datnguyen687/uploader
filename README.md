# Ecommerce

Backend HTTP service

## Features

- Backend REST API to handle ecommerce products

## Dependencies

- [golang](https://golang.org/) (>=1.7) to compile/run
- [make]() to run Makefile
- [swag](https://github.com/swaggo/swag) to generate swagger for go
- [mockery](https://github.com/vektra/mockery) to generate mocks for unittest
- [docker](https://www.docker.com/) to compile Dockerfile
- [docker-compose](https://docs.docker.com/compose/) to compile docker-compose.yaml

## Steps to run (without docker)
1. Install Golang using this [link](https://go.dev/)
2. Set up dependencies (installing mockery and swagger)
```sh
make init
```
3. To generate swagger:
```sh
make swagger
```
4. To generate mocks:
```sh
make mockery
```
5. To run tests:
```sh
make test
```
6. To build:
```sh
make build
```
7. Replace `<DB_...>` in .env file with real db connection info:
```sh
% cat .env
PORT=8080
DEBUG=true
SWAGGER_ENABLED=true

DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=dat.nguyen
DB_NAME=ecommerce
DB_PASSWORD=password

% source .env
```
8. Run:
```sh
./bin/ecommerce
```
9. To use docker-compose (optional):
```sh
make up // to start
make down // to stop
```
10. To inject data, you can use the script under `./db/init.sql`
## Steps to check if it runs Ok:
1. Check Swagger UI:
```sh
http://localhost:8080/internal/swagger/index.html
```
Note: this will only be valid if you set `SWAGGER_ENABLED` to true
2. Use curl:
```sh
curl -X 'GET' \
  'http://localhost:8080/v1/ping' \
  -H 'accept: application/json'
```

The response should be:
```json
{
  "code": 1,
  "message": "ok",
  "error": ""
}
```

```sh
curl -X 'GET' \
  'http://localhost:8080/v1/products' \
  -H 'accept: application/json'
```
The response should be:
```json
{
  "code": 1,
  "message": "ok",
  "error": "",
  "data": [
    {
      "ID": 2,
      "Name": "test name",
      "Brand": "test brand",
      "Price": 123,
      "CreatedAt": "2022-10-24T15:23:56.875491Z",
      "UpdatedAt": null
    },
    {
      "ID": 3,
      "Name": "test",
      "Brand": "test brand",
      "Price": 0.568505341224743,
      "CreatedAt": "2022-10-26T12:12:55.284157Z",
      "UpdatedAt": null
    },
    {
      "ID": 4,
      "Name": "test",
      "Brand": "test brand",
      "Price": 0.283186833389724,
      "CreatedAt": "2022-10-26T12:14:01.614275Z",
      "UpdatedAt": null
    },
    ...
    ]
}
```