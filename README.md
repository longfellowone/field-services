## Field Services Golang Server

Go gRPC Server. See instructions below to get started with docker.

Purchasing client: [field-services-purchasing](https://github.com/longfellowone/field-services-purchasing)
- ReactJS
- gRPC-Web

Mobile client: [field-services-mobile](https://github.com/longfellowone/field-services-mobile)
- Flutter
- gRPC-Dart  

## Linux/Mac

Use the following commands in the root directory. 

#### `make up`
To start docker containers

#### `make down`
To stop docker containers

#### `make build`
To rebuild docker containers

#### `make clean`
To stop docker containers and clean up

## Windows

1. Open the .env file and change the SERVICE_NAME from 'grpcsvc' to 'host.docker.internal'
2. Use the following commands in the root directory

#### `docker-compose up -d`
To start docker containers

#### `docker-compose down`
To stop docker containers

#### `docker-compose build`
To rebuild docker containers

## Todo

- Fix tests
- More...

## Potential Errors

It make take a few seconds for the containers to start up. If you are getting errors, first try shutting down and restarting the containers

#### `Error: Http response at 400 or 500 level`
Issue connecting to Envoy, check envoy has not crashed

#### `Error: upstream connect error or disconnect/reset before headers. reset reason: connection failure`
Envoy cannot find service, check SERVICE address in envoy config is correct

#### `Error: no healthy upstream`
Envoy OK, check go service has not crashed