# ==================================================================================== #
#  MAKE FILE INSTRUCTION BY MAZELI VICTOR CHIDUBEM
#  01/05/2024
# ==================================================================================== #

# Variables
APP_NAME := grpc-cli-quiz

.PHONY: help
help: ## Display this help message
	@echo "Usage: make <target>"
	@echo
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# ==================================================================================== #
# CODE QUALITY CHECK
# ==================================================================================== #

.PHONY: lint
lint:
	golangci-lint run

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix

.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

.PHONY: mod-verify
mod-verify:
	go mod verify

# ==================================================================================== #
# PACKAGE DOWNLOAD
# ==================================================================================== #

.PHONY: mod-download
mod-download:
	go mod download

# ==================================================================================== #
# PROTO GENERATE
# ==================================================================================== #

proto-generate:
	protoc --go_out=./proto --go-grpc_out=./proto/ proto/quiz.proto

# ==================================================================================== #
# START PROGRAME
# ==================================================================================== #
start:
	go run . start-quiz