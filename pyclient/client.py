import logging

import grpc
import iguana_pb2
import iguana_pb2_grpc

def execute():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    print("Will try to greet world ...")
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = helloworld_pb2_grpc.GreeterStub(channel)
        response = stub.SayHello(helloworld_pb2.HelloRequest(name="you"))
    print("Greeter client received: " + response.message)


def execute2():
    with grpc.insecure_channel("localhost:8080") as channel:
        stub = mixer_pb2_grpc.MixerStub(channel)
        response = stub.EnqueueCommand(mixer_pb2.EnqueueRequest(command="test", client_id="client"))
        print("Response: " + response.receipt_id)

if __name__ == "__main__":
    logging.basicConfig()
    execute()