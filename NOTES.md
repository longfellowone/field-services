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