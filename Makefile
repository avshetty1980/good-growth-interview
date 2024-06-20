include .env

run: build
	@./bin/api

build:
	@go build -o ${BINARY} bin/api

test:
	@go test -v ./...