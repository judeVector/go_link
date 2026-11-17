# Load environment variables from .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Default database URL (override with .env or make DATABASE_URL=...)
DATABASE_URL ?= postgresql://user:pass@localhost:5432/go_link?sslmode=disable
MIGRATIONS_PATH ?= ./cmd/migrate/migrations

.PHONY: migrate-up migrate-down migrate-force migrate-version migrate-create

## Apply all pending migrations
migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

## Rollback all migrations
migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down

## Rollback N migrations (usage: make migrate-down-n N=1)
migrate-down-n:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down $(N)

## Force set migration version (usage: make migrate-force VERSION=1)
migrate-force:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" force $(VERSION)

## Show current migration version
migrate-version:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version

## Create a new migration (usage: make migrate-create NAME=create_posts_table)
migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(NAME)
