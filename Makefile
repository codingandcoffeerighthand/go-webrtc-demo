VERSION := $(shell cat VERSION)
COMMIT := $(shell git rev-parse HEAD)
GOOSE_DRIVER := postgres
GOOSE_DBSTRING := "postgres://user:password@localhost:5432/db?sslmode=disable"
GOOSE_MIGRATION_DIR := ./db/migrations

build:
	go build \
	-o build/app \
	-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT)" \
	cmd/*.go
.PHONY: build
	

db-status:
	goose ${GOOSE_DRIVER} ${GOOSE_DBSTRING} status