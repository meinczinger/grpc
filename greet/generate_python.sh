python -m grpc_tools.protoc -Igreetpb --python_out=greetpb/python  \
--grpc_python_out=greetpb/python greetpb/greet.proto