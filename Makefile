BINARY_NAME := kubectx-manager
MODULE_NAME := github.com/eta-cybersolutions/kubectx-manager
VERSION ?= $(shell git describe --tags --always --dirty)
BUILD_DIR := ./build

.PHONY: all build clean install

all: build

build:
	@echo "🔧 Building $(BINARY_NAME)..."
	go mod tidy
	GO111MODULE=on CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(BINARY_NAME) ./main.go
	@echo "✅ Built $(BUILD_DIR)/$(BINARY_NAME)"

install:
	@echo "📦 Installing $(BINARY_NAME) to $$HOME/.local/bin..."
	mkdir -p $$HOME/.local/bin
	cp $(BUILD_DIR)/$(BINARY_NAME) $$HOME/.local/bin/$(BINARY_NAME)
	@echo "✅ Installed at $$HOME/.local/bin/$(BINARY_NAME)"

clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
