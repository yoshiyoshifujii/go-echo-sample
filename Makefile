SHELL := /bin/bash

.PHONY: help generate run fmt vet test tidy build

help:
	@echo "make generate  - OpenAPI からコード生成 (oapi-codegen)"
	@echo "make run       - サーバー起動"
	@echo "make fmt       - go fmt"
	@echo "make vet       - go vet"
	@echo "make test      - go test ./..."
	@echo "make tidy      - go mod tidy"
	@echo "make build     - go build ./cmd/api"

generate:
	go generate ./internal/api

run:
	go run ./cmd/api

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

tidy:
	go mod tidy

build:
	go build ./cmd/api
