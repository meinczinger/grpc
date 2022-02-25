python -m grpc_tools.protoc -Igreet/greetpb --python_out=greet/greetpb/python  \
--grpc_python_out=greet/greetpb/python greet/greetpb/greet.proto