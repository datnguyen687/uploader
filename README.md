# Uploader

Backend HTTP service

## Features

- Backend REST API to upload a file to Google Cloud Storage

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
7. Replace `<GCS_KEY_PATH>` and `<GCS_BUCKET_NAME>` in .env file with path to gcs key:
```sh
% cat .env
PORT=8080
DEBUG=true
SWAGGER_ENABLED=true
GCS_KEY_PATH=<GCS_KEY_PATH>
GCS_BUCKET_NAME=<GCS_BUCKET_NAME>

% source .env
```
8. Run:
```sh
./bin/uploader
```

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
curl -X 'POST' \
  'http://localhost:8080/v1/user/batch' \
  -H 'accept: application/json' \
  -H 'Content-Type: multipart/form-data' \
  -F 'file=@payload.json;type=application/json'
```
The response should be:
```json
{
  "code": 1,
  "message": "ok",
  "error": ""
}
```