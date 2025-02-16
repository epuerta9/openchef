DB_URL="sqlite://local.db"

.PHONY: build run dev test generate clean setup

build:
	go build -o bin/api ./cmd/api
	go build -o bin/worker ./cmd/worker

run:
	./bin/api

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