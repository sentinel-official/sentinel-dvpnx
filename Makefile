.DEFAULT_GOAL := help

GIT_COMMIT := $(shell git log -1 --format='%H' 2>/dev/null || echo "unknown")
GIT_TAG    := $(shell git describe --tags --always --dirty 2>/dev/null | sed 's/^v//')

TAGS      := $(strip netgo)
LD_FLAGS  := -s -w \
	-X github.com/sentinel-official/sentinel-go-sdk/version.Commit=$(GIT_COMMIT) \
	-X github.com/sentinel-official/sentinel-go-sdk/version.Tag=$(GIT_TAG)

build_flags = -ldflags="$(LD_FLAGS)" -mod=readonly -tags="$(TAGS)" -trimpath

GOBIN ?= $(shell go env GOBIN)
ifeq ($(GOBIN),)
  GOBIN := $(shell go env GOPATH)/bin
endif

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build        Build the binary (./bin/sentinel-dvpnx)"
	@echo "  install      Install the binary to \$$GOBIN"
	@echo "  clean        Remove build artifacts"
	@echo "  test         Run tests with coverage"
	@echo "  benchmark    Run benchmarks"
	@echo "  go-lint      Run golangci-lint with auto-fix"
	@echo "  build-image  Build Docker image"

.PHONY: build
build:
	go build $(build_flags) -o ./bin/sentinel-dvpnx main.go

.PHONY: install
install:
	go build $(build_flags) -o "$(GOBIN)/sentinel-dvpnx" main.go

.PHONY: clean
clean:
	$(RM) -r ./bin ./vendor

.PHONY: test
test:
	go test -cover -mod=readonly -v ./...

.PHONY: benchmark
benchmark:
	go test -bench -mod=readonly -v ./...

.PHONY: go-lint
go-lint:
	golangci-lint run --fix

.PHONY: build-image
build-image:
	docker build --compress --file Dockerfile --force-rm --tag sentinel-dvpnx .
