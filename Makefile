.PHONY: help build clean test install lint format

# Default target
help:
	@echo "GuruUI Development Commands:"
	@echo ""
	@echo "  build     - Build the GuruUI binary"
	@echo "  install   - Install dependencies and build"
	@echo "  test      - Run all tests"
	@echo "  lint      - Run linter checks"
	@echo "  format    - Format Go code"
	@echo "  clean     - Clean build artifacts"
	@echo "  run       - Run GuruUI locally"
	@echo "  docker    - Build Docker image"
	@echo ""

# Build the binary
build:
	@echo "Building GuruUI..."
	go build -o guruui cmd/guruui/main.go
	@echo "Build complete! Binary: ./guruui"

# Install dependencies and build
install:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download
	@echo "Building GuruUI..."
	go build -o guruui cmd/guruui/main.go
	@echo "Install complete! Binary: ./guruui"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Format code
format:
	@echo "Formatting Go code..."
	go fmt ./...
	go vet ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f guruui
	rm -f coverage.out
	rm -f coverage.html
	@echo "Clean complete!"

# Run locally
run: build
	@echo "Running GuruUI..."
	./guruui --help

# Build Docker image
docker:
	@echo "Building Docker image..."
	docker build -t guruui:latest .
	@echo "Docker image built: guruui:latest"

# Development setup
dev-setup:
	@echo "Setting up development environment..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	@echo "Development tools installed!"

# Quick development cycle
dev: format test build
	@echo "Development cycle complete!"

# Release build
release: clean
	@echo "Building release binaries..."
	GOOS=linux GOARCH=amd64 go build -o guruui-linux-amd64 cmd/guruui/main.go
	GOOS=darwin GOARCH=amd64 go build -o guruui-darwin-amd64 cmd/guruui/main.go
	GOOS=windows GOARCH=amd64 go build -o guruui-windows-amd64.exe cmd/guruui/main.go
	@echo "Release binaries built!"
