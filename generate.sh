#!/bin/bash

# protoc --go-grpc_out=. greet/greetpb/greet.proto 
protoc --proto_path=./greet/greetpb --go_out=./greet/greetpb/go --go_opt=paths=source_relative \
    --go-grpc_out=./greet/greetpb/go --go-grpc_opt=paths=source_relative \
    greet.proto 