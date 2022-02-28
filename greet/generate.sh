#!/bin/bash

# protoc --go-grpc_out=. greet/greetpb/greet.proto 
protoc --proto_path=./greetpb --go_out=./greetpb/go --go_opt=paths=source_relative \
    --go-grpc_out=./greetpb/go --go-grpc_opt=paths=source_relative \
    greet.proto 