BINARY_NAME := $(shell git config --get remote.origin.url | awk -F/ '{print $$5}' | awk -F. '{print $$1}')
BINARY_VERSION := $(shell git describe --tags)
BINARY_BUILD_DATE := $(shell date +%d.%m.%Y)
WIN_BINARY_NAME := $(BINARY_NAME).exe
BUILD_FOLDER := .build

PRINTF_FORMAT := "\033[35m%-18s\033[33m %s\033[0m\n"

ABS_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
ifeq ($(shell go env GOHOSTOS), windows)
	ABS_PATH = $(PWD)
endif

.PHONY: all build windows linux vendor gen-webapi clean

all: build

build: windows linux ## Default: build for windows and linux

windows: vendor $(BUILD_FOLDER)/windows/logconfig.json ## Build artifacts for windows
	@printf $(PRINTF_FORMAT) BINARY_NAME: $(WIN_BINARY_NAME)
	@printf $(PRINTF_FORMAT) BINARY_BUILD_DATE: $(BINARY_BUILD_DATE)
	mkdir -p $(BUILD_FOLDER)/windows
	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc  go build -ldflags "-s -w -X $(BINARY_NAME).Version=$(BINARY_VERSION) -X $(BINARY_NAME).BuildDate=$(BINARY_BUILD_DATE)" -o $(BUILD_FOLDER)/windows/$(WIN_BINARY_NAME) .

linux: vendor $(BUILD_FOLDER)/linux/logconfig.json ## Build artifacts for linux
	@printf $(PRINTF_FORMAT) BINARY_NAME: $(BINARY_NAME)
	@printf $(PRINTF_FORMAT) BINARY_BUILD_DATE: $(BINARY_BUILD_DATE)
	mkdir -p $(BUILD_FOLDER)/linux
	env GOOS=linux GOARCH=amd64  go build -ldflags "-s -w -X $(BINARY_NAME).Version=$(BINARY_VERSION) -X $(BINARY_NAME).BuildDate=$(BINARY_BUILD_DATE)" -o $(BUILD_FOLDER)/linux/$(BINARY_NAME) .

vendor: ## Get dependencies according to go.sum
	env GO111MODULE=auto go mod tidy
	env GO111MODULE=auto go mod vendor

clean: ## Remove vendor and artifacts
	rm -rf vendor
	rm -rf $(BUILD_FOLDER)
