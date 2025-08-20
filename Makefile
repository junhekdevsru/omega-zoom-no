SHELL := /bin/bash

APP := agentmgr

PROTO_DIR := api/proto
BIN_DIR := bin

.PHONY: proto dev-up dev-down migrate-up migrate-down run lint tidy easyp-generate easyp-mod-vendor

proto: easyp-generate

proto:
	@chmod +x ./easyp
	./easyp generate


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

lint:
	@echo "tip: добавь golangci-lint позже"

tidy:
	go mod tidy

.PHONY: lint lint-ci

lint:
	@command -v golangci-lint >/dev/null 2>&1 || { echo "Install golangci-lint: https://golangci-lint.run/welcome/install/"; exit 1; }
	golangci-lint run --timeout=5m

# как в CI
lint-ci:
	golangci-lint run --timeout=5m --out-format=github-actions

.PHONY: test
test:
	GRPC_ADDR=":8080" \
	PG_CONN="postgres://agent:agent@localhost:5432/agentdb?sslmode=disable" \
	MIGRATIONS_DIR="migrations" \
	go test -race -count=1 ./...

