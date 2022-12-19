# go-task-application

A Task Management web application written in Golang microservices using the hexagonal architecture pattern.

This mono-repo holds all each microservice as an individual Go module. Each service holds a Dockerfile for containerization.

![](docs/diagram.png)

## Services 

`gateway-service` - [README](https://github.com/samverrall/go-task-application/blob/main/gateway-service/README.md)

A reverse proxy gateway which takes a REST HTTP API and forwards through to gRPC calls.

`task-service` - [README](https://github.com/samverrall/go-task-application/blob/main/task-service/README.md)


`user-service` - [README](https://github.com/samverrall/go-task-application/blob/main/user-service/README.md)

Handles user authentication and authorisation. Directly communicates user and session data with the gateway. Also exposes a gRPC adapter that other services can invoke if needed.

## Service Pattern 

See the [Task Service](https://github.com/samverrall/go-task-application/tree/main/task-service) for an example hexagonal (ports and adapters) microservice.

- cmd/
	- grpc/
		- main.go
- internal/
	- adapters/
		- left/
			- grpc/
		- right/
			- sqlite/
		- port/
			- domain/
			- repository/
			- service/
- pkg/
	- config/
