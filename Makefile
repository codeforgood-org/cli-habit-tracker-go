# Makefile for Habit Tracker CLI

# Variables
BINARY_NAME=habit
MAIN_PATH=./cmd/habit
GO=go
GOFLAGS=-v
LDFLAGS=-s -w
INSTALL_PATH=/usr/local/bin

# Colors for output
COLOR_RESET=\033[0m
COLOR_BOLD=\033[1m
COLOR_GREEN=\033[32m
COLOR_YELLOW=\033[33m

.PHONY: all build test coverage clean install uninstall run help lint fmt vet

# Default target
all: clean build test

## help: Display this help message
help:
	@echo "$(COLOR_BOLD)Habit Tracker CLI - Makefile Commands$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_GREEN)Build Commands:$(COLOR_RESET)"
	@echo "  make build          - Build the binary"
	@echo "  make install        - Install the binary to $(INSTALL_PATH)"
	@echo "  make uninstall      - Remove the installed binary"
	@echo "  make clean          - Remove build artifacts"
	@echo ""
	@echo "$(COLOR_GREEN)Development Commands:$(COLOR_RESET)"
	@echo "  make test           - Run all tests"
	@echo "  make coverage       - Run tests with coverage report"
	@echo "  make lint           - Run linters (requires golangci-lint)"
	@echo "  make fmt            - Format code"
	@echo "  make vet            - Run go vet"
	@echo "  make run            - Run the application"
	@echo ""
	@echo "$(COLOR_GREEN)Other:$(COLOR_RESET)"
	@echo "  make all            - Clean, build, and test"
	@echo "  make help           - Display this help message"

## build: Build the binary
build:
	@echo "$(COLOR_BOLD)Building $(BINARY_NAME)...$(COLOR_RESET)"
	$(GO) build $(GOFLAGS) -ldflags="$(LDFLAGS)" -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "$(COLOR_GREEN)✓ Build complete: $(BINARY_NAME)$(COLOR_RESET)"

## test: Run all tests
test:
	@echo "$(COLOR_BOLD)Running tests...$(COLOR_RESET)"
	$(GO) test -v ./...
	@echo "$(COLOR_GREEN)✓ Tests complete$(COLOR_RESET)"

## coverage: Run tests with coverage
coverage:
	@echo "$(COLOR_BOLD)Running tests with coverage...$(COLOR_RESET)"
	$(GO) test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "$(COLOR_GREEN)✓ Coverage report generated: coverage.html$(COLOR_RESET)"

## lint: Run linters
lint:
	@echo "$(COLOR_BOLD)Running linters...$(COLOR_RESET)"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
		echo "$(COLOR_GREEN)✓ Linting complete$(COLOR_RESET)"; \
	else \
		echo "$(COLOR_YELLOW)⚠ golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest$(COLOR_RESET)"; \
	fi

## fmt: Format code
fmt:
	@echo "$(COLOR_BOLD)Formatting code...$(COLOR_RESET)"
	$(GO) fmt ./...
	@echo "$(COLOR_GREEN)✓ Formatting complete$(COLOR_RESET)"

## vet: Run go vet
vet:
	@echo "$(COLOR_BOLD)Running go vet...$(COLOR_RESET)"
	$(GO) vet ./...
	@echo "$(COLOR_GREEN)✓ Vet complete$(COLOR_RESET)"

## clean: Remove build artifacts
clean:
	@echo "$(COLOR_BOLD)Cleaning build artifacts...$(COLOR_RESET)"
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out coverage.html
	@echo "$(COLOR_GREEN)✓ Clean complete$(COLOR_RESET)"

## install: Install the binary
install: build
	@echo "$(COLOR_BOLD)Installing $(BINARY_NAME) to $(INSTALL_PATH)...$(COLOR_RESET)"
	@sudo cp $(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "$(COLOR_GREEN)✓ Installed successfully$(COLOR_RESET)"

## uninstall: Remove the installed binary
uninstall:
	@echo "$(COLOR_BOLD)Uninstalling $(BINARY_NAME)...$(COLOR_RESET)"
	@sudo rm -f $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "$(COLOR_GREEN)✓ Uninstalled successfully$(COLOR_RESET)"

## run: Run the application
run: build
	@./$(BINARY_NAME)

## mod: Download dependencies
mod:
	@echo "$(COLOR_BOLD)Downloading dependencies...$(COLOR_RESET)"
	$(GO) mod download
	$(GO) mod tidy
	@echo "$(COLOR_GREEN)✓ Dependencies updated$(COLOR_RESET)"
