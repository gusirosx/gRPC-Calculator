# gRPC-Calculator

This project creates a API to get two params via REST protocol and pass to a microservice using gRPC protocol and return a given response to the client using REST protocol, achieving communication between micro-services in Golang.

## How to run this example

1. run grpc server

```sh
$ make run_server
```

2. run gin client

```sh
$ make run_client
```

# Input

1. use browser to test the application using the following link

```sh
http://localhost:8080/add/125/125
```
and
```sh
http://localhost:8080/Multiply/125/125
```

# Output
```
id: 61150315 
name: "Gustavo Rodrigues" 
username: "gusirosx" 
```

## How to create proto files

1. use the makefile

```sh
$ make generate
```

# Links
Project based on this [tutorial](https://www.youtube.com/watch?v=Y92WWaZJl24)
