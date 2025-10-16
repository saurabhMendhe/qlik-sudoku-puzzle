.PHONY: all build test bench clean run install fmt vet coverage help

#https://github.com/golangci/golangci-lint
GOLANGCI_LINT_VERSION := v1.61.0
GOLANGCI_LINT := bin/golangci-lint
GOLANGCI_LINT_URL := https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh
LINT_EXISTS := $(or $(and $(wildcard $(GOLANGCI_LINT)),1),0)

#https://github.com/securego/gosec
GOSEC_VERSION := v2.21.0
GOSEC := bin/gosec
GOSEC_URL := https://raw.githubusercontent.com/securego/gosec/master/install.sh
GOSEC_EXISTS := $(or $(and $(wildcard $(GOSEC)),1),0)

# Default target
all: fmt vet test build

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod verify

# Install development tools
install-tools:
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Build the application
build:
	@echo "Building qlik-sudoku-puzzle..."
	@go build -o qlik-sudoku-puzzle

# Run tests
test:
	@echo "Running tests..."
	@go test ./sudoku -v

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	@go test ./sudoku -cover
	@go test ./sudoku -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run linter
lint:
ifeq ($(LINT_EXISTS), 0)
	curl -sfL ${GOLANGCI_LINT_URL} | sh -s ${GOLANGCI_LINT_VERSION}
endif
	$(GOLANGCI_LINT)  run

## security: secure the source code
security:
ifeq ($(GOSEC_EXISTS), 0)
	curl -sfL ${GOSEC_URL} | sh -s ${GOSEC_VERSION}
endif
	$(GOSEC) -exclude-dir util -exclude-dir test/mock ./...

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	@go test ./sudoku -bench=. -benchmem

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

# Run the application with example
run: build
	@echo "Running qlik-sudoku-puzzle with example .."
	@./qlik-sudoku-puzzle -input 5,3,0,0,7,0,0,0,0,6,0,0,1,9,5,0,0,0,0,9,8,0,0,0,0,6,0,8,0,0,0,6,0,0,0,3,4,0,0,8,0,3,0,0,1,7,0,0,0,2,0,0,0,6,0,6,0,0,0,0,2,8,0,0,0,0,4,1,9,0,0,5,0,0,0,0,8,0,0,7,9

# Install to $GOPATH/bin
install:
	@echo "Installing qlik-sudoku-puzzle..."
	@go install

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f qlik-sudoku-puzzle
	@rm -f coverage.out coverage.html

# Docker commands
docker-build:
	@echo "Building Docker image..."
	@docker build -t qlik-sudoku-puzzle:latest .

docker-run:
	@echo "Running Docker container..."
	@docker run --rm qlik-sudoku-puzzle:latest

# docker-push:
# 	@echo "Pushing Docker image..."
# 	@docker tag qlik-sudoku-puzzle:latest $(DOCKER_REGISTRY)/qlik-sudoku-puzzle:latest
# 	@docker push $(DOCKER_REGISTRY)/qlik-sudoku-puzzle:latest

# Show help
help:
	@echo "Available targets:"
	@echo "  all       - Format, vet, test, and build (default)"
	@echo "  build     - Build the application"
	@echo "  test      - Run tests"
	@echo "  coverage  - Run tests with coverage report"
	@echo "  bench     - Run benchmarks"
	@echo "  fmt       - Format code"
	@echo "  vet       - Run go vet"
	@echo "  run       - Build and run with example puzzle"
	@echo "  install   - Install to GOPATH/bin"
	@echo "  clean     - Remove build artifacts"
	@echo "  help      - Show this help message"