.PHONY: test test-coverage test-verbose test-integration test-all lint fmt build clean help

# Default target
help:
	@echo "Available targets:"
	@echo "  make test              - Run unit tests (with mocks)"
	@echo "  make test-integration  - Run integration tests (real API calls)"
	@echo "  make test-all          - Run both unit and integration tests"
	@echo "  make test-coverage     - Run tests with coverage report"
	@echo "  make test-verbose      - Run tests with verbose output"
	@echo "  make fmt               - Format code"
	@echo "  make lint              - Run linters"
	@echo "  make build             - Build the project"
	@echo "  make clean             - Clean build artifacts"
	@echo "  make deps              - Download dependencies"

# Run unit tests (with mocks)
test:
	go test ./...

# Run tests with coverage
test-coverage:
	go test -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run tests with verbose output
test-verbose:
	go test -v ./...

# Run integration tests (real API calls)
test-integration:
	@echo "Running integration tests (real API calls)..."
	go test -v -tags=integration ./...

# Run all tests (unit + integration)
test-all:
	@echo "Running unit tests..."
	go test ./...
	@echo ""
	@echo "Running integration tests..."
	go test -v -tags=integration ./...

# Format code
fmt:
	go fmt ./...
	gofmt -s -w .

# Run linters (requires golangci-lint)
lint:
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Install from https://golangci-lint.run/usage/install/" && exit 1)
	golangci-lint run

# Build the project
build:
	go build ./...

# Clean build artifacts
clean:
	go clean
	rm -f coverage.out coverage.html

# Download dependencies
deps:
	go mod download
	go mod tidy

