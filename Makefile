VERSION := $(shell cat VERSION)
COMMIT := $(shell git rev-parse HEAD)

build:
	go build \
	-o build/app \
	-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT)" \
	cmd/*.go
.PHONY: build