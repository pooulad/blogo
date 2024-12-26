#!/bin/bash

export COMPOSE_PROJECT_NAME=blogo

echo "Building and starting containers..."
docker-compose up --build -d

echo "Waiting for services to start..."
sleep 5

echo "Displaying logs for the API and DB services..."
docker-compose logs -f
