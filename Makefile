SHELL := /bin/bash

APP := agentmgr

PROTO_DIR := api/proto
BIN_DIR := bin

.PHONY: proto dev-up dev-down migrate-up migrate-down run lint tidy

proto:
	@chmod +x scripts/gen-proto.sh
	@scripts/gen-proto.sh

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