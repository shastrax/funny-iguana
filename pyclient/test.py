import logging

import grpc
import iguana_pb2
import iguana_pb2_grpc

def execute():
    print("hit")
    with grpc.insecure_channel("localhost:50053") as channel:
        print("xxx")
#        stub = iguana_pb2_grpc.IguanaServiceStub(channel)
#        response = stub.Ping(iguana_pb2.Ping(source="test"))
#        print("Response: " + response.status)

if __name__ == "__main__":
    logging.basicConfig()
    execute()
