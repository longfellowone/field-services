# Field Services Golang Server

Go gRPC Server for Purchasing and Mobile clients.

Purchasing client (ReactJS): [field-services-purchasing](https://github.com/longfellowone/field-services-purchasing) 
![](http://g.recordit.co/bK2FPTJWjM.gif | width=400)

Mobile client (Flutter): [field-services-mobile](https://github.com/longfellowone/field-services-mobile)
![](http://g.recordit.co/21luo1Y6Hs.gif | width=400) 


## Table of Contents
* [Getting Started](#getting-started)
* [Running Go server without docker](#running-go-server-without-docker)
* [Potential Errors](#potential-errors)
* [TODO](#todo)

## Getting Started

Use the following command in the root directory to start the docker containers.

```sh
docker-compose up -d
```

## Running Go server without docker

1. Stop all running containers with `docker-compose down`
2. Open the .env file and change the `SERVICE` variable from `grpcsvc` to `localhost` on linux or `host.docker.internal` if using windows
3. Comment out the entire `grpcsvc` in docker-compose.yml
4. Restart docker containers with `docker-compose up -d`
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