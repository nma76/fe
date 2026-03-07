# Makefile for building and running the fe application

APP=fe
CMD=./cmd/fe
BIN=./bin

# Default to host OS/architecture if not set
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Add .exe extension for Windows builds
EXT := 
ifeq ($(GOOS),windows)
	EXT := .exe
endif

.PHONY: all build clean

# Default target
all: build

# Build the application
build:
	mkdir -p $(BIN)
	go build -o $(BIN)/$(APP)$(EXT) $(CMD)

# Run the application
run:
	go run $(CMD)

# Clean up the build artifacts
clean:
	rm -rf $(BIN)