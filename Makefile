include .envrc	

MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY:migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))


.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY:migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-force
migrate-force:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) force $(version)
.PHONY: migrate-version
migrate-version:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) version

.PHONY: db-reset
db-reset:
	@echo "Dropping and recreating database..."
	@dropdb --if-exists social
	@createdb social
	@make migrate-up
