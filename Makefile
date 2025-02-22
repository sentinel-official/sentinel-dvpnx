.DEFAULT_GOAL := noop

.PHONY: noop
noop:
	@:

GIT_COMMIT := $(shell git log -1 --format='%H')
GIT_TAG := $(shell git describe --tags | sed 's/^v//' | rev | cut -d - -f 2- | rev)

TAGS := $(strip netgo)
LD_FLAGS := -s -w \
	-X github.com/sentinel-official/sentinel-go-sdk/version.Commit=${GIT_COMMIT} \
	-X github.com/sentinel-official/sentinel-go-sdk/version.Tag=${GIT_TAG}

.PHONY: benchmark
benchmark:
	@go test -bench -mod=readonly -v ./...

.PHONY: build
build:
	go build -ldflags="${LD_FLAGS}" -mod=readonly -tags="${TAGS}" -trimpath \
		-o ./bin/sentinel-dvpnx main.go

.PHONY: clean
clean:
	rm -rf ./bin ./vendor

.PHONY: install
install:
	go build -ldflags="${LD_FLAGS}" -mod=readonly -tags="${TAGS}" -trimpath \
		-o "${GOPATH}/bin/sentinel-dvpnx" main.go

.PHONY: build-image
build-image:
	@docker build --compress --file Dockerfile --force-rm --tag sentinel-dvpnx .

.PHONY: go-lint
go-lint:
	@golangci-lint run --fix

.PHONY: test
test:
	@go test -cover -mod=readonly -v ./...

.PHONY: tools
tools:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
