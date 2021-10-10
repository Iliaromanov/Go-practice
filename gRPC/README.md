## General Notes

https://grpc.io/ 


gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

- gRPC APIs are language neutral (like REST)


### Setup

- Install protocol buffer compiler and add it to PATH
- Install the main grpc package
> `go get -u google.golang.org/grpc`
- Install protocol buffer package for go
> `go get -u github.com/golang/protobuf/protoc-gen-go`

### Proto Buffs:

https://developers.google.com/protocol-buffers  