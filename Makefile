SHELL := /bin/bash

.PHONY: proto dev-up dev-down migrate-up migrate-down run run-gateway lint tidy test easyp-generate easyp-mod-vendor

proto: easyp-generate

easyp-generate:
	@chmod +x ./easyp
	./easyp generate

easyp-mod-vendor:
	@chmod +x ./easyp
	./easyp mod vendor

dev-up:
	docker compose -f deploy/docker-compose.yml up -d
	@chmod +x scripts/dev-wait.sh && scripts/dev-wait.sh

dev-down:
	docker compose -f deploy/docker-compose.yml down -v

migrate-up:
	go run ./cmd/migrate --direction=up

migrate-down:
	go run ./cmd/migrate --direction=down

run:
	go run ./cmd/agentmgr

run-gateway:
	HTTP_ADDR=":8082" go run ./cmd/agentmgr

lint:
	@echo "add golangci-lint later"

tidy:
	go mod tidy

test:
	GRPC_ADDR=":8080" \
	PG_CONN="postgres://agent:agent@localhost:5432/agentdb?sslmode=disable" \
	MIGRATIONS_DIR="migrations" \
	go test -race -count=1 ./...