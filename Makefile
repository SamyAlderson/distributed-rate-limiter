# Set default Go binary to use for building
GOBIN := $(shell go env GOBIN)
GOBUILD := go build -mod=readonly
GOCLEAN := go clean -modcache
GOTEST := go test -mod=readonly -v -cover

# Project root directory
PROJECT_ROOT := $(shell dirname $(abspath $(firstword $(MAKEFILE_LIST))))

# Target binaries for distribution
TARGET_BINARIES := $(PROJECT_ROOT)/bin/distributed-rate-limiter

# Build target
build: go.sum
	$(GOBUILD) -o $(PROJECT_ROOT)/bin/distributed-rate-limiter main.go
	@echo "Built distributed-rate-limiter"

# Test target
test: go.sum
	$(GOTEST) -covermode=atomic -coverprofile=$(PROJECT_ROOT)/coverage.out ./...
	@echo "Ran tests"

# Clean target
clean:
	$(GOCLEAN)
	rm -f $(PROJECT_ROOT)/coverage.out
	rm -rf $(PROJECT_ROOT)/bin

# Install target (generate Go modules and install dependencies)
install:
	go mod init distributed-rate-limiter
	go mod tidy
	go mod vendor
	@echo "Installed dependencies"

# Release target (generate binaries and update version)
release: install
	go mod verify
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o $(PROJECT_ROOT)/bin/distributed-rate-limiter
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(PROJECT_ROOT)/bin/distributed-rate-limiter
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o $(PROJECT_ROOT)/bin/distributed-rate-limiter
	@echo "Generated binaries"

# Default target
default: build

# Run the rate limiter
run:
	$(PROJECT_ROOT)/bin/distributed-rate-limiter --config=$(PROJECT_ROOT)/config.json

# Generate Go sum file
go.sum:
	@echo "go.sum generated"

# Generate coverage report
coverage:
	gocov convert $(PROJECT_ROOT)/coverage.out > $(PROJECT_ROOT)/coverage.txt
	gocov report -format=text $(PROJECT_ROOT)/coverage.txt

# Clean up after running tests
.PHONY: clean-test
clean-test:
	rm -rf $(PROJECT_ROOT)/tmp