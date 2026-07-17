.PHONY: dev api web db sqlc

dev: db api web

db:
	docker compose up -d db

api:
	cd apps/api && go run ./cmd/api

web:
	cd apps/web && bun run dev

sqlc:
	cd apps/api && sqlc generate

build-api:
	cd apps/api && go build -o bin/api ./cmd/api

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-build:
	docker compose build

lint-api:
	cd apps/api && golangci-lint run

test-api:
	cd apps/api && go test ./...
