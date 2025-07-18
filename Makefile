# Include environment variables, if your .env.make defines DB_MIGRATOR_ADDR
# For commands run via docker-compose exec, we will use the internal Docker network address (db:5432)
# so DB_MIGRATOR_ADDR might not be directly used for those.
include .env.make
export

MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: migrate-create
migration-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@echo "Running migrations up inside Docker..."
	@docker-compose exec -T db migrate -path=/migrations -database 'postgres://admin:adminpassword@db:5432/social?sslmode=disable' up

.PHONY: migrate-down
migrate-down:
	@echo "Running migrations down inside Docker..."
	@docker-compose exec -T db migrate -path=/migrations -database 'postgres://admin:adminpassword@db:5432/social?sslmode=disable' down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: seed
seed:
	@echo "Seeding the database inside Docker..."
	@docker-compose exec -T db sh -c "psql -U admin -d social < /migrations/scripts/seed.sql"
	@echo "Database seeded successfully."

.PHONY: db-reset
db-reset:
	@echo "Destroying old database volume..."
	@docker-compose down -v --remove-orphans
	@echo "Building new custom database image..."
	@docker-compose build
	@echo "Starting fresh database..."
	@docker-compose up -d
	@echo "Waiting for database to be ready..."
	@docker-compose exec -T db sh -c "until pg_isready -U admin -d social; do echo 'Waiting...'; sleep 1; done"
	@echo "Running all migrations..."
	@make migrate-up
	@echo "Database reset and migrated successfully."

.PHONY: db-seed-fresh
db-seed-fresh: db-reset seed



migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" up

migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" down

 migrate create -seq -ext sql -dir ./cmd/migrate/migrations 
