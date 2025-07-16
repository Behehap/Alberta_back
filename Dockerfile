# Start with the official PostgreSQL 16.3 image
FROM postgres:16.3

# Install curl, which we need to download the migrate tool
RUN apt-get update && apt-get install -y curl

# Download and install the migrate tool directly into our image
# This makes it available to run inside the container
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate /usr/bin/migrate

# Set the working directory for our migrations inside the container
WORKDIR /migrations
