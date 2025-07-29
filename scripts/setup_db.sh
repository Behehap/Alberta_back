#!/bin/bash

echo "--- Starting Database Setup ---"

# 1. Stop and remove old database volumes
echo "1. Stopping and removing old database volumes..."
docker-compose down -v --remove-orphans

# 2. Build Docker images (if any custom images are defined, e.g., for 'api' service)
echo "2. Building Docker images (including database image)..."
docker-compose build

# 3. Start fresh database and other services
echo "3. Starting fresh database and other services..."
docker-compose up -d

# 4. Wait for the PostgreSQL database to be ready
echo "4. Waiting for the PostgreSQL database to be ready..."
# Loop until pg_isready reports success (or timeout)
until docker-compose exec -T db pg_isready -U admin -d social; do
  echo "Database is unavailable - sleeping"
  sleep 1
done
echo "Database is up and running!"

# 5. Run database migrations
echo "5. Running database migrations..."
make migrate-up # This command runs migrations on your host, connecting to the Docker DB

# 6. Seeding the database
echo "6. Seeding the database..."
# Pipe the seed.sql content directly into the psql command inside the container
cat scripts/seed.sql | docker-compose exec -T db psql -U admin -d social
echo "Database seeded successfully."