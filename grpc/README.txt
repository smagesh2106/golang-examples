go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


----

syntax = "proto3";

package hello;

option go_package = "example.com/hello/hellopb";

service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}


---------

protoc --go_out=. --go-grpc_out=. hello.proto
