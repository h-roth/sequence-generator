# Sequence generator service

Implement a small HTTP service in Golang that generates a unique sequence number for every request it receives.
Requirements:

* The service should listen on a configurable port and respond to HTTP GET requests.
* The service should return a unique sequence number in the response body for each request it receives.
* The service should ensure that the sequence number generated is unique across all requests.
* The service should handle multiple requests concurrently and should be thread-safe.

Additional features:

* Swagger docs given this is a public API.
* GRPC option
* Test coverage - unit and integration

## Considerations

TODO setup with config options for port as per anthgg
Probably use a cmd/ package
Write tests
Determine the max size sequence number we can supprot
Determine requests per second
If messages are created at the same time use other distinguishing features

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

Use the timestamp of the request to generate a unique id.
If two requests come in at the same time, use the IP address to distinguish them.

Depending on the application we could vary this distinguishing feature -  i.e passing an additional sequence number from the client.

**Guaranteed unique sequential IDs**

Use node id, machine id and a counter along with timestamp. On a single machine this is guaranteed to be unique however we also benefit from being able to scale this service to multiple machines.

### Load testing

Done with `ali` for lightweight load testing.

`ali --duration=3s --rate=1000 http://localhost:8080/`
