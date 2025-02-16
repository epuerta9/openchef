DB_URL="sqlite://local.db"

.PHONY: build run dev test generate clean setup release release-tag release-build

build:
	go build -o bin/api ./cmd/api
	go build -o bin/openchef ./cmd/openchef

run:
	./bin/openchef

dev:
	air

test:
	go test ./...

generate:
	templ generate

clean:
	rm -rf bin tmp

setup:
	go mod tidy
	go install github.com/cosmtrek/air@latest

# Database commands
.PHONY: db-setup db-status db-migrate db-new db-generate

db-setup:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	curl -sSf https://atlasgo.sh | sh

db-status:
	atlas schema inspect -u ${DB_URL}

db-migrate:
	atlas schema apply \
		-u ${DB_URL} \
		--to file://internal/database/schema/schema.sql \
		--dev-url "sqlite://file?mode=memory"

db-new:
	atlas migrate new

db-generate:
	sqlc generate

# NATS commands
.PHONY: nats-dev nats-prod

nats-dev:
	unset NATS_CONFIG

nats-prod:
	export NATS_CONFIG=./config/nats.conf

VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS := -X github.com/epuerta9/openchef/internal/version.Version=$(VERSION) \
           -X github.com/epuerta9/openchef/internal/version.Commit=$(COMMIT) \
           -X github.com/epuerta9/openchef/internal/version.BuildTime=$(BUILD_TIME)

release-build:
	go build -ldflags "$(LDFLAGS)" -o bin/openchef cmd/openchef/main.go

release-tag:
	@echo "Current version: $(VERSION)"
	@read -p "New version (e.g., 0.1.0): " version; \
	git tag -a "v$$version" -m "Release v$$version"
	@echo "Tagged new version. Run 'git push --tags' to trigger release"

release: release-build release-tag
	@echo "Release prepared. To publish:"
	@echo "1. git push origin main"
	@echo "2. git push --tags"
	@echo "This will trigger the GitHub Actions release workflow" 