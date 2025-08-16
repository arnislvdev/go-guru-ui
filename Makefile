.PHONY: help build clean test install lint format

# Default target
help:
	@echo "GuruUI Development Commands:"
	@echo ""
	@echo "  build     - Make the program"
	@echo "  install   - Get parts and make the program"
	@echo "  test      - Check if everything works"
	@echo "  lint      - Check code quality"
	@echo "  format    - Make code look nice"
	@echo "  clean     - Remove old files"
	@echo "  run       - Run the program"
	@echo "  docker    - Make a Docker container"
	@echo ""

# Build the binary
build:
	@echo "Making GuruUI..."
	go build -o guruui cmd/guruui/main.go
	@echo "Done! Program: ./guruui"

# Install dependencies and build
install:
	@echo "Getting the parts..."
	go mod tidy
	go mod download
	@echo "Making GuruUI..."
	go build -o guruui cmd/guruui/main.go
	@echo "Done! Program: ./guruui"

# Run tests
test:
	@echo "Checking if everything works..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Checking everything and making a report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Report ready: coverage.html"

# Run linter
lint:
	@echo "Checking code quality..."
	golangci-lint run

# Format code
format:
	@echo "Making code look nice..."
	go fmt ./...
	go vet ./...

# Clean build artifacts
clean:
	@echo "Removing old files..."
	rm -f guruui
	rm -f coverage.out
	rm -f coverage.html
	@echo "Done!"

# Run locally
run: build
	@echo "Running GuruUI..."
	./guruui --help

# Build Docker image
docker:
	@echo "Making Docker container..."
	docker build -t guruui:latest .
	@echo "Docker container ready: guruui:latest"

# Development setup
dev-setup:
	@echo "Setting up tools for development..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	@echo "Development tools ready!"

# Quick development cycle
dev: format test build
	@echo "Development cycle done!"

# Release build
release: clean
	@echo "Making programs for different systems..."
	GOOS=linux GOARCH=amd64 go build -o guruui-linux-amd64 cmd/guruui/main.go
	GOOS=darwin GOARCH=amd64 go build -o guruui-darwin-amd64 cmd/guruui/main.go
	GOOS=windows GOARCH=amd64 go build -o guruui-windows-amd64.exe cmd/guruui/main.go
	@echo "Programs ready for different systems!"
