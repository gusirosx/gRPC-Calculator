# gRPC-Tutorial
Based on this [tutorial](https://www.youtube.com/watch?v=Y92WWaZJl24)

## How to run this example

1. run the grpc server

```sh
$ go run server/main.go
```
or
```sh
$ make run_server
```
2. run the client

```sh
$ go run client/main.go
```
or
```sh
$ make run_client
```

## How to create proto files

1. use the makefile

```sh
$ make generate