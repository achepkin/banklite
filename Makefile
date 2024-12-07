.PHONY: up build run test lint
make all: build test lint

up:
	docker compose up  --remove-orphans --build

build:
	go build -race -o app cmd/main.go

run:
	go build -race -o app cmd/main.go && \
	HTTP_ADDR=:8080 \
	DEBUG_ERRORS=1 \
	./app

test:
	go test -race ./...

install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install go.uber.org/mock/mockgen@latest
	go install github.com/vektra/mockery/v2@latest

lint:
	golangci-lint run ./...

generate:
	mockery