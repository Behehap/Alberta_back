# This command now runs migrations INSIDE the docker container.
# It uses a simpler DSN that is more reliable inside Docker.
.PHONY: migrate-up
migrate-up:
	@echo "Running migrations up..."
	@docker-compose exec -T db migrate -path /migrations -database 'postgres://admin:adminpassword@db:5432/social?sslmode=disable' -verbose up

# This command also runs inside the container.
.PHONY: migrate-down
migrate-down:
	@echo "Running migrations down..."
	@docker-compose exec -T db migrate -path /migrations -database 'postgres://admin:adminpassword@db:5432/social?sslmode=disable' -verbose down

.PHONY: seed
seed:
	@echo "Seeding the database..."
	@docker-compose exec -T db psql -U admin -d social < scripts/seed.sql
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
	@timeout /t 5 >nul
	@echo "Running all migrations..."
	@make migrate-up
	@echo "Database reset and migrated successfully."

# This is the only command you need to run from now on to start fresh.
.PHONY: db-seed-fresh
db-seed-fresh: db-reset seed

