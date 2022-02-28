#!/bin/bash

# protoc --go-grpc_out=. greet/greetpb/greet.proto 
protoc --proto_path=./calculatorpb --go_out=./calculatorpb/go --go_opt=paths=source_relative \
    --go-grpc_out=./calculatorpb/go --go-grpc_opt=paths=source_relative \
    calculator.proto 