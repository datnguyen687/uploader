init:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/vektra/mockery/v2/.../

swagger:
	swag init

build:
	go build -o ./bin/ecommerce ./main.go

up:
	docker-compose up --build -d

down:
	docker-compose down

mockery:
	mockery --output ./mocks --all

test:
	go test -v ./...