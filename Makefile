# Makefile for Go project

# Variables
GOFMT=gofmt
GOVET=go vet

# Default target
all: fmt vet

# Run go fmt
fmt:
	@echo "Running go fmt..."
	@$(GOFMT) -l -w .

# Run go vet
vet:
	@echo "Running go vet..."
	@$(GOVET) ./...

# Clean (placeholder, add cleaning tasks if needed)
clean:
	@echo "Cleaning..."

.PHONY: all fmt vet clean
