# Sequence generator service

Implement a small HTTP service in Golang that generates a unique sequence number for every request it receives.
Requirements:

* The service should listen on a configurable port and respond to HTTP GET requests.
* The service should return a unique sequence number in the response body for each request it receives.
* The service should ensure that the sequence number generated is unique across all requests.
* The service should handle multiple requests concurrently and should be thread-safe.

## Instructions

`go run ./...`

`go test ./...`

To make a request hit:
`http://localhost:8080/sequence`

To mimic client hit:
`http://localhost:8080/cmd/client`

No API docs as so simple. Would add swagger and/or grpc if this was a real service.

## Assumptions

* IDs must be unique.
* IDs are numerical values only.
* IDs fit into 64-bits.
* IDs are ordered sequentially by date (So we can't use UUIDs).
* Ability to generate over 10,000 unique IDs per second. (Stretch goal)

## Two approaches

### Centralised counter

All requests go through one counter mechanism.

Pros:

* Simple

Cons:

* Single point of failure
* Constrained by having to go through lock
* Locks can cause tricky bugs and are hard to test. In this case it's very simple but still a consideration.

Note:
Atomic implementation was around 2x faster than mutex implementation.

### Timestamp based id

The basic idea is to use a timestamp plus some distinguishing feature to generate a sequential unique id.
This would be unique on a single instance but also could be scaled to multiple instances by using node id or similar.

I'm using nanoseconds for the timestamp, which takes 51 bits. This leaves 13 bits for other distinguishing features.

**Guaranteed unique sequential IDs**

Use node id, machine id and a counter along with timestamp. On a single machine this is guaranteed to be unique however we also benefit from being able to scale this service to multiple machines. In our implementation we are using **server side counter with lock**, however you could also use a client side counter and ip to distinguish requests.

Pros:

* Able to scale to multiple instances

Cons:

* Still using locks, could probably get rid of these with client side counter.

### Load testing

Done with `ali` for lightweight load testing.

`ali --duration=3s --rate=1000 http://localhost:8080/sequence`

Results can be seen in latencies folder.
