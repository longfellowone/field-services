#! /bin/bash
echo "Starting docker containers..."
cd docker
docker-compose up -d
echo "Starting fielsvc..."
cd ../cmd/fieldsvc/
GOOS=linux CGO_ENABLED=0 go build -o fieldsvc .
./fieldsvc