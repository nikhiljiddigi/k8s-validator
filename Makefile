# Makefile for k8s-validate CLI

BINARY_NAME := k8s-validate
BIN_DIR     := bin

.PHONY: all build run test lint vendor clean help

all: build

# Clean the CLI binary
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

# Compile the CLI binary from project root (main.go)
build: clean
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(BINARY_NAME) .
	@chmod +x $(BIN_DIR)/$(BINARY_NAME)
	@echo "Built $(BIN_DIR)/$(BINARY_NAME)"

# Run the CLI against sample manifests
run:
	@echo "Running CLI against manifests/"
	@./$(BIN_DIR)/$(BINARY_NAME) validate -f manifests/ --exemptions .exemptions.yaml --output table

# Run unit tests
test:
	@go test ./... -v

# Lint (requires golangci-lint)
lint:
	@golangci-lint run

# Vendor dependencies
vendor:
	@go mod vendor

# Clean up
clean:
	@rm -rf $(BIN_DIR)

# Help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile \
	  | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'
