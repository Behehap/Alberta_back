#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

echo "--- Starting Database Setup ---"

echo "1. Stopping and removing old database volumes..."
docker-compose down -v --remove-orphans

echo "2. Building Docker images (including database image)..."
docker-compose build

echo "3. Starting fresh database and other services..."
# The migrator service will automatically run 'migrate up' here
docker-compose up -d

echo "4. Waiting for the PostgreSQL database to be ready..."
until docker-compose exec -T db sh -c "pg_isready -U admin -d social"; do
  echo "Database is unavailable - sleeping"
  sleep 2
done
echo "Database is up and running!"

# REMOVED STEP 5: Running database migrations (up)...
# This step is no longer needed because the 'migrator' service in docker-compose.yml
# is configured to run migrations automatically on startup.
# docker-compose exec -T migrator migrate -path=/migrations -database 'postgres://admin:adminpassword@db:5432/social?sslmode=disable' up

echo "5. Seeding the database..." # Renumbered to 5
docker-compose exec -T db sh -c "psql -U admin -d social < /app/scripts/seed.sql"

echo "--- Database Setup Complete ---"