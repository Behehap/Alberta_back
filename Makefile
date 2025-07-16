# This command now runs migrations INSIDE the docker container.
# It uses a simpler DSN that is more reliable inside Docker.
# Create new migration files with sequential 6-digit prefix
.PHONY: migration-create
migration-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make migration-create NAME=<migration_name>"; \
		exit 1; \
	fi; \
	mkdir -p cmd/migrate/migrations; \
	next_seq=$$(ls cmd/migrate/migrations | grep -E '^[0-9]+' | cut -d '_' -f 1 | sort -n | tail -n 1 | sed 's/^0*//'); \
	next_seq=$${next_seq:-0}; \
	next_seq=$$((next_seq + 1)); \
	next_seq=$$(printf "%06d" $$next_seq); \
	touch "cmd/migrate/migrations/$${next_seq}_$(NAME).up.sql"; \
	touch "cmd/migrate/migrations/$${next_seq}_$(NAME).down.sql"; \
	echo "Created cmd/migrate/migrations/$${next_seq}_$(NAME).up.sql"; \
	echo "Created cmd/migrate/migrations/$${next_seq}_$(NAME).down.sql";


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
	@sleep 5 # Changed from timeout /t 5 >nul to 'sleep 5' for cross-platform compatibility
	@echo "Running all migrations..."
	@make migrate-up
	@echo "Database reset and migrated successfully."

# This is the only command you need to run from now on to start fresh.
.PHONY: db-seed-fresh
db-seed-fresh: db-reset seed