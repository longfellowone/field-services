TO DO

default:
	@echo "=============building Local API============="
	docker build -f cmd/api/Dockerfile -t api .

up: default
	@echo "=============starting api locally============="
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test -v -cover ./...

clean: down
	@echo "=============cleaning up============="
	rm -f api
	docker system prune -f
	docker volume prune -f

	#! /bin/bash
    cd docker
    echo "Stopping services ..."
    docker-compose down


    #! /bin/bash
    echo "Starting docker containers..."
    cd docker
    docker-compose up -d
    echo "Starting grpcsvc..."
    cd ../cmd/grpcsvc/
    GOOS=linux CGO_ENABLED=0 go build -o grpcsvc .
    ./grpcsvc