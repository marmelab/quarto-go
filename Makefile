MAKEFLAGS += --silent

.PHONY: help install run lint

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

BIN = docker run \
	-it \
	--rm \
	-v "${PWD}:/go" \
	-p "8080:8080" \
	--name quarto-go \
	quarto-go

install: ## Install docker environnement
	docker build --tag=quarto-go .
	$(BIN) go get github.com/gorilla/mux
	$(BIN) go get github.com/ahl5esoft/golang-underscore

run: ## Start the game
	 $(BIN) go run ./src/quarto/main.go

test: ## Test the code
	$(BIN) go test -v ./src/tests

lint: ## Check the code syntax and rules
	$(BIN) gofmt -w ./src

.DEFAULT_GOAL := help
