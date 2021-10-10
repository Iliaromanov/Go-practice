## General Notes

https://grpc.io/ 

https://developers.google.com/protocol-buffers 

gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

- gRPC APIs are language neutral (like REST)
- protocol buffers serve a similar purpose to json or xml files (serialization), but are much smaller and much more efficient

### Setup

- Install protocol buffer compiler and add it to PATH
- Install the main grpc package
> `go get -u google.golang.org/grpc`
- Install protocol buffer package for go
> `go get -u github.com/golang/protobuf/protoc-gen-go`

### Proto File:

- proto file tells Go how it should encode/decode various pieces of data
- the package defined in proto file will be the name of the Go module we work with. eg. `package proto;`
- message is used to define the structure of Requests/Responses when serialized. eg:
```Go
message Request {
  int64 a = 1;
  int64 b = 2;
}
```
  The numbers that a,b are set to denote the size of those variables within the buffer when serialized. Eg. nubmers 1-15 mean those values will take up 1 byte when the request is serialized, so a,b above will take up 1 byte each. 166-2047 mean the values will take up 2 bytes.
- define a service using the 'service' keyword. Eg:
```Go
service NameOfService {
  rpc NameOfFunc1(Request) returns (Response);
  rpc NameOfFunc2(Request) returns (Response);
  ...
}
```
  The Request and Response in the above eg. should be defined with the 'message' keyword.
- Proto files are compiled into Go files
> `protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto service.proto`