# Makefile for the maker-checker service

APP_NAME = MakerChecker
CMD_PATH = ./cmd/main.go
BUILD_DIR = ./bin

.PHONY: all build fmt run tidy vendor clean

all: tidy build

build:
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_PATH)

fmt:
	@echo "Formatting code..."
	gofmt -s -w .
	go fmt ./...

run:
	@echo "Running the application..."
	go run $(CMD_PATH)

tidy:
	@echo "Tidying up modules..."
	go mod tidy

vendor:
	@echo "Creating vendor directory..."
	go mod vendor

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)