# Makefile for 6502 Emulator Project

# Go source files
SRC := $(main.go)

# Binary output directory
BIN_DIR := bin

# Binary name
BIN_NAME := emulator

# Go build command
BUILD_CMD := go build -o $(BIN_DIR)/$(BIN_NAME)

# Default target
all: build

# Build the project
build:
	@echo "Building the project..."
	@mkdir -p $(BIN_DIR)
	$(BUILD_CMD) $(SRC)

# Run the project
run: build
	@echo "Running the project..."
	@$(BIN_DIR)/$(BIN_NAME)

# Clean the build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Format the code
fmt:
	@echo "Formatting the code..."
	@go fmt ./...

# Lint the code
lint:
	@echo "Linting the code..."
	@golangci-lint run

.PHONY: all build run clean test fmt lint

rom: FORCE
	@echo "Building the ROM..."
	@vasm -Fbin -dotdir -wdc02 -o ./rom/bios.bin ./asm/HelloWorld.asm

FORCE: