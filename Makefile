.PHONY: help run build test fmt vet lint

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-8s\033[0m %s\n", $$1, $$2}'

run: ## Run the application
	go run ./cmd/server

build: ## Compile the application
	go build ./...

test: ## Run the tests
	go test ./...

fmt: ## Format the code
	go fmt ./...

vet: ## Report suspicious code
	go vet ./...

lint: ## Run golangci-lint
	golangci-lint run
