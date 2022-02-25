import grpc
import logging
from greet.greetpb.python.greet_pb2 import Greeting
import greet_pb2, greet_pb2_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = greet_pb2_grpc.GreetServiceStub(channel)
        response = stub.Greet(
            greet_pb2.GreetRequest(
                greeting=Greeting(first_name="Tibor", last_name="Meinczinger")
            )
        )
    print("Greeter client received: " + response.result)


if __name__ == "__main__":
    logging.basicConfig()
    run()
