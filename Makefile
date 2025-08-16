.PHONY: help build test clean examples install

# Default target
help:
	@echo "GuruUI - Terminal UI Library for Go"
	@echo ""
	@echo "Available targets:"
	@echo "  build     - Build the library"
	@echo "  test      - Run tests"
	@echo "  clean     - Clean build artifacts"
	@echo "  examples  - Build example applications"
	@echo "  install   - Install the library"
	@echo "  demo      - Run the main demo"
	@echo "  manga     - Run the manga CLI example"
	@echo "  npm       - Run the npm-style example"
	@echo "  tasks     - Run the task manager app"

# Build the library
build:
	go build -o guruui ./...

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	go clean
	rm -f guruui

# Build example applications
examples:
	cd examples && go build -o demo main.go
	cd examples/manga_cli && go build -o manga-cli main.go
	cd examples/npm_style && go build -o npm-demo main.go
	cd examples/task_manager && go build -o task-manager main.go

# Install the library
install:
	go install ./...

# Run the main demo
demo: examples
	./examples/demo

# Run the manga CLI example
manga: examples
	./examples/manga_cli/manga-cli

# Run the npm-style example
npm: examples
	./examples/npm_style/npm-demo

# Run the task manager app
tasks: examples
	./examples/task_manager/task-manager

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run

# Check dependencies
deps:
	go mod tidy
	go mod verify

# Generate documentation
docs:
	@echo "Documentation is in the docs/ directory"
	@echo "API reference: docs/API.md"

# Run all examples
all-examples: demo manga npm tasks
