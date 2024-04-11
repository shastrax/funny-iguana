import logging

import grpc
import iguana.v1.iguana_pb2
import iguana.v1.iguana_pb2_grpc

def execute():
    request = iguana.v1.iguana_pb2.PingRequest(source = "pyclient")
    # print(request)

    with grpc.insecure_channel("localhost:8080") as channel:
        stub = iguana.v1.iguana_pb2_grpc.IguanaServiceStub(channel)
        response = stub.Ping(request)
        print("response: " + response.status)

if __name__ == "__main__":
    logging.basicConfig()
    execute()
