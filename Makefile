# Makefile for building and running the fe application

# Application name
APP=fe

# Entry point for the application
CMD=./cmd/fe

# Output directory for the built binary
BIN=./bin

# Version information, set at build time using -ldflags
VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT ?= $(shell git rev-parse --short HEAD)
BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Set ldflags for embedding version information into the binary
LDFLAGS = -X 'main.Version=$(VERSION)' -X 'main.Commit=$(COMMIT)' -X 'main.BuildDate=$(BUILD_DATE)'

# Default to the host's GOOS and GOARCH if not set
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Default build (native)
all: build

# Suppress command output
.SILENT:

# Make sure make runs even if files exist with the same name as targets
.PHONY: all build clean install uninstall release build-linux build-darwin build-windows

# Add .exe extension for Windows builds
EXT := 
ifeq ($(GOOS),windows)
	EXT := .exe
endif

build:
	echo "Building $(APP) for $(GOOS)/$(GOARCH)..."
	mkdir -p $(BIN)
	go build -ldflags "$(LDFLAGS)" -o $(BIN)/$(APP)$(EXT) $(CMD)
	echo "Build complete: $(BIN)/$(APP)$(EXT)"

install: build
	if [ -w /usr/local/bin ]; then \
		echo "Installing to /usr/local/bin..."; \
		install -m 755 $(BIN)/$(APP)$(EXT) /usr/local/bin/$(APP); \
	else \
		echo "Using sudo to install to /usr/local/bin..."; \
		sudo install -m 755 $(BIN)/$(APP)$(EXT) /usr/local/bin/$(APP); \
	fi
	echo "Installed!"

uninstall:
	if [ -w /usr/local/bin ]; then \
		echo "Uninstalling from /usr/local/bin..."; \
		rm -f /usr/local/bin/$(APP); \
	else \
		echo "Using sudo to uninstall from /usr/local/bin..."; \
		sudo rm -f /usr/local/bin/$(APP); \
	fi
	echo "Uninstalled!"

clean:
	echo "Cleaning..."
	rm -rf $(BIN)
	echo "Clean complete."

# cross-compile the application for all supported platforms
build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BIN)/$(APP)-linux-amd64 $(CMD)
	GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(BIN)/$(APP)-linux-arm64 $(CMD)
	echo "Linux builds complete!"

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BIN)/$(APP)-darwin-amd64 $(CMD)
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(BIN)/$(APP)-darwin-arm64 $(CMD)
	echo "MacOs builds complete!"

build-windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BIN)/$(APP)-windows-amd64.exe $(CMD)
	GOOS=windows GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(BIN)/$(APP)-windows-arm64.exe $(CMD)
	echo "Windows builds complete!"

release: clean build-linux build-darwin build-windows
	echo "All platform builds complete!"