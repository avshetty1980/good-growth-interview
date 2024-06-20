include .env

run: build
	/bin/api

build:
	docker build -t good-growth-api .

test:
	go test -v ./...

up:
	docker-compose up --build -d --remove-orphans

down:
	docker-compose down