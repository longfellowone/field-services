#! /bin/bash
echo "Starting docker containers..."
cd docker
docker-compose up -d
echo "Starting supplysvc..."
cd ../cmd/supplysvc/
GOOS=linux CGO_ENABLED=0 go build -o fieldsvc .
./supplysvc