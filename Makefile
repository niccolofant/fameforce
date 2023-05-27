include .env

DATABASE_HOST ?= localhost
DATABASE_PORT ?= 5432
DATABASE_NAME ?= postgres
DATABASE_USER ?= postgres
DATABASE_PASSWORD ?= postgres

MIGRATIONS_DIR=migrations

COLOR_INFO=\033[1;34m

new_migration:
	@if [ -z "$(name)" ]; then \
		echo "Please specify a name for the migration."; \
		exit 1; \
	fi
	@echo "${COLOR_INFO}Creating new migration: ${name}..."
	migrate create -ext sql -dir ${MIGRATIONS_DIR} -seq $(name)

migrate_up:
	@echo "${COLOR_INFO}Running migrations..."
	migrate -path ${MIGRATIONS_DIR} -database "postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable" up

migrate_down:
	@echo "${COLOR_INFO}Rolling back migrations..."
	migrate -path ${MIGRATIONS_DIR} -database "postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable" down

