#! /bin/bash
echo "Starting docker containers..."
cd docker
docker-compose up -d
echo "Starting grpcsvc..."
cd ../cmd/grpcsvc/
GOOS=linux CGO_ENABLED=0 go build -o grpcsvc .
./grpcsvc