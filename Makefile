# Makefile for building and running the fe application

# Application name
APP=fe

# Entry point for the application
CMD=./cmd/fe

# Output directory for the built binary
BIN=./bin

# Supported platforms
SUPPORTED := darwin linux freebsd netbsd openbsd  windows

# Detect the host OS
OS := $(shell uname | tr '[:upper:]' '[:lower:]')

# Handle special case for Windows (mingw)
ifeq ($(findstring mingw,$(OS)),mingw)
	OS := windows
endif

# Default to host OS/architecture if not set
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Add .exe extension for Windows builds
EXT := 
ifeq ($(GOOS),windows)
	EXT := .exe
endif

# Suppress command output
.SILENT:

# Make sure make runs even if files exist with the same name as targets
.PHONY: all build clean prereq

# Default target
all: build

# Check for prerequisites
prereq:
	echo "Detected platform: $(OS)"
	if echo "$(SUPPORTED)" | grep -wq "$(OS)"; then \
		echo "Platform $(OS) is supported."; \
	else \
		echo "Error: Platform $(OS) is not supported."; \
		exit 1; \
	fi

# Build the application
build: clean prereq
	echo "Building $(APP) for $(GOOS)/$(GOARCH)..."
	mkdir -p $(BIN)
	go build -o $(BIN)/$(APP)$(EXT) $(CMD)
	echo "Build complete!"

# Run the application
run:
	go run $(CMD)

# Clean up the build artifacts
clean:
	echo "Cleaning..."
	rm -rf $(BIN)
	echo "Clean complete."