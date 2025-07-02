# iyzipay-go Makefile

.PHONY: help build test test-race test-cover lint fmt vet clean examples deps security check-deps

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Build targets
build: ## Build the library
	@echo "Building iyzipay-go..."
	@go build -v ./...

examples: ## Build all examples
	@echo "Building examples..."
	@cd examples/basic_payment && go build -o ../../bin/basic_payment .
	@cd examples/threeds_payment && go build -o ../../bin/threeds_payment .
	@cd examples/checkout_form && go build -o ../../bin/checkout_form .
	@cd examples/card_management && go build -o ../../bin/card_management .
	@cd examples/refund_cancel && go build -o ../../bin/refund_cancel .
	@echo "Examples built in bin/ directory"

# Test targets
test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-race: ## Run tests with race detection
	@echo "Running tests with race detection..."
	@go test -race -v ./...

test-cover: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -race -coverprofile=coverage.out -covermode=atomic ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

bench: ## Run benchmarks
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

# Code quality targets
lint: ## Run golangci-lint
	@echo "Running golangci-lint..."
	@golangci-lint run

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@goimports -w .

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

security: ## Run security scan with gosec
	@echo "Running security scan..."
	@gosec ./...

staticcheck: ## Run staticcheck
	@echo "Running staticcheck..."
	@staticcheck ./...

# Dependency management
deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download

deps-update: ## Update dependencies
	@echo "Updating dependencies..."
	@go get -u ./...
	@go mod tidy

check-deps: ## Check for known vulnerabilities
	@echo "Checking for known vulnerabilities..."
	@go list -json -deps ./... | nancy sleuth

# Release targets
tag: ## Create a new tag (usage: make tag VERSION=v1.0.0)
ifndef VERSION
	@echo "Usage: make tag VERSION=v1.0.0"
	@exit 1
endif
	@echo "Creating tag $(VERSION)..."
	@git tag -a $(VERSION) -m "Release $(VERSION)"
	@git push origin $(VERSION)

release-check: ## Check if ready for release
	@echo "Checking if ready for release..."
	@make test-cover
	@make lint
	@make vet
	@make security
	@echo "✅ All checks passed - ready for release!"

# Utility targets
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@go clean ./...

install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/sonatypecommunity/nancy@latest

demo: examples ## Run demo applications
	@echo "Running basic payment demo..."
	@./bin/basic_payment

all: fmt vet lint test build examples ## Run all checks and build everything

# CI/CD targets (used by GitHub Actions)
ci-test: deps vet lint test-race ## Run all CI tests

ci-build: build examples ## Build all CI artifacts

.DEFAULT_GOAL := help