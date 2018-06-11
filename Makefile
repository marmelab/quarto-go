MAKEFLAGS += --silent

.PHONY: help install run lint

help: 
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

BIN = docker run \
	--interactive \
	--rm \
	-v "${PWD}:/code" \
	--name quarto-go \
	quarto-go

install: ## Install docker environnement
	docker build --tag=quarto-go .

run: ## Start the game
	 $(BIN) go run ./src/quarto.go

test: ## Test the code
	$(BIN) go test ./src/tests
lint: ## Check the code syntax and rules
	$(BIN) golint ./src

.DEFAULT_GOAL := help
