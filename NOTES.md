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
[OID]: oid1 - [PID]: pid1 - [STATUS]: ->New [POs]: 
	3 name3 Back Ordered(6:28PM) req:84 rec:73 po:
[OID]: oid2 - [PID]: pid1 - [STATUS]: ->New ->Sent ->Complete [POs]: po1 | po3 | 
	1 name1 Filled(6:28PM) req:12 rec:12 po:po9
	2 name2 Filled(6:28PM) req:23 rec:23 po:
