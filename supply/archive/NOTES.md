####
protoc -I=. supply.proto \
--js_out=import_style=commonjs:. \
--grpc-web_out=import_style=commonjs,mode=grpcwebtext:.

####
go get -u
go mod tidy

####
WDIR := /go/src/github.com/user/repo
DIR := ${CURDIR}:${WDIR}

DOCKER_IMAGE := user/image

login:
	docker run -i -v $(DIR) -w $(WDIR) --entrypoint=/bin/bash -t $(DOCKER_IMAGE)

dockerbuild:
	docker build -f Dockerfile -t $(DOCKER_IMAGE) .

dockerpush:
	docker push $(DOCKER_IMAGE):latest

PHONY: dockerbuild login dockerpush

####
Use _____ID to auto generate resolver for lazy fetch

#### Hot reload
https://github.com/oxequa/realize/issues/190

#### .SH Files
chmod u+x ./

#### Google wire
https://github.com/google/go-cloud/tree/master/samples
https://github.com/terashi58/wire-example
https://github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/tree/master/ch10
https://groups.google.com/forum/#!forum/natsio

####
https://github.com/altairsix/eventsource/blob/master/pgstore/store.go


-> API receive command, publish event

-> Subscribe to order.event
-> On event, service.DoEvent(), return awk

-> repository.Load events by ID from local DB
-> Apply events
-> DoEvent()
-> repository.Save event to local DB
-> Publish event


-> Subscribe to order.event
-> On event, service.DoEvent(), return awk

-> gRPC command
-> repository.Load events by ID from local DB
-> Apply events
-> Validate then apply new event
-> repository.Save event to local DB
// -> Publish event
-> gRPC response

####
Dont return pointer item, return full struct
Item = Item.DoSomething

####
Map PID to slice of Items

####
Long switch statements 
[]func()
https://dizzy.zone/2018/07/28/Refactoring-Go-switch-statements/
https://play.golang.org/p/O0Dl1Nj9INz

#### gRPC Fix
go get github.com/golang/protobuf@8d0c54c1246661d9a51ca0ba455d22116d485eaa

#### GO MOD
go mod tidy
