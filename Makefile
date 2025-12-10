SHELL := /bin/bash

.PHONY: generate run

generate:
	go generate ./internal/api

run:
	go run ./cmd/api
