#!/bin/bash

export COMPOSE_PROJECT_NAME=blogo

echo "Stopping and removing containers..."
docker-compose down

# Optionally remove volumes (use cautiously, as this will delete data)
# echo "Removing volumes..."
# docker-compose down --volumes
