# Field Services Golang Server

Go gRPC Server for Purchasing and Mobile clients.

Purchasing client (ReactJS): [field-services-purchasing](https://github.com/longfellowone/field-services-purchasing)  
Mobile client (Flutter): [field-services-mobile](https://github.com/longfellowone/field-services-mobile)

## Table of Contents
* [Getting Started](#getting-started)
  * [Linux](#linux)
  * [Windows](#windows)
* [Running Go server without docker](#running-go-server-without-docker)
* [Potential Errors](#potential-errors)
* [TODO](#todo)

## Getting Started

### Linux

Use the following commands in the root directory. 

```sh
make up
```
To start docker containers

```sh
make down
```
To stop docker containers

```sh
make build
```
To rebuild docker containers

```sh
make clean
```
To stop docker containers and clean up

### Windows

1. Open the .env file and change the `SERVICE` variable from `grpcsvc` to `host.docker.internal`

2. Use the following commands in the root directory

```sh
docker-compose up -d
```
To start docker containers

```sh
docker-compose down
```
To stop docker containers

```sh
docker-compose build
```
To rebuild docker containers

## Running Go server without docker

1. Stop all running containers
2. Open the .env file and change the `SERVICE` variable from `grpcsvc` to  your local IP address e.g. `192.168.0.176`
3. Comment out the entire `grpcsvc` in docker-compose.yml
4. Restart docker containers
5. In the root directory run `go run cmd/grpcsvc/main.go`

## Potential Errors

It may take a few seconds for the containers to start up. If you are getting errors, first try shutting down and restarting the containers

```Error: Http response at 400 or 500 level```  
Issue connecting to Envoy, check envoy has not crashed

```Error: upstream connect error or disconnect/reset before headers. reset reason: connection failure```  
Envoy cannot connect to outgoing service, check SERVICE environmental variable in .env file is correct

```Error: no healthy upstream```  
Envoy OK, check Go service is running and has not crashed

## Todo

- Fix tests
- More...