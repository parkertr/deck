.PHONY: test lint build clean help

# Default target
all: test lint

# Run tests
test:
	go test ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run linter (requires golangci-lint to be installed)
lint:
	golangci-lint run

# Build the example
build:
	go build -o bin/example ./example

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Format code
fmt:
	go fmt ./...

# Run go mod tidy
tidy:
	go mod tidy

# Install dependencies
deps:
	go mod download

# Install golangci-lint as a go tool
install-lint:
	@echo "Installing golangci-lint as a go tool..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.1.6

# Run all checks
check: fmt tidy test lint

# Show help
help:
	@echo "Available targets:"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  lint           - Run linter"
	@echo "  build          - Build example program"
	@echo "  clean          - Clean build artifacts"
	@echo "  fmt            - Format code"
	@echo "  tidy           - Run go mod tidy"
	@echo "  deps           - Download dependencies"
	@echo "  install-lint   - Install golangci-lint"
	@echo "  check          - Run fmt, tidy, test, and lint"
	@echo "  help           - Show this help"
