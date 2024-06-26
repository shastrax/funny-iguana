# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from iguana.v1 import iguana_pb2 as iguana_dot_v1_dot_iguana__pb2


class IguanaServiceStub(object):
    """import "google/type/datetime.proto";


    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Ping = channel.unary_unary(
                '/iguana.v1.IguanaService/Ping',
                request_serializer=iguana_dot_v1_dot_iguana__pb2.PingRequest.SerializeToString,
                response_deserializer=iguana_dot_v1_dot_iguana__pb2.PingResponse.FromString,
                )
        self.VisitorEvent = channel.unary_unary(
                '/iguana.v1.IguanaService/VisitorEvent',
                request_serializer=iguana_dot_v1_dot_iguana__pb2.VisitorEventRequest.SerializeToString,
                response_deserializer=iguana_dot_v1_dot_iguana__pb2.VisitorEventResponse.FromString,
                )


class IguanaServiceServicer(object):
    """import "google/type/datetime.proto";


    """

    def Ping(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def VisitorEvent(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_IguanaServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Ping': grpc.unary_unary_rpc_method_handler(
                    servicer.Ping,
                    request_deserializer=iguana_dot_v1_dot_iguana__pb2.PingRequest.FromString,
                    response_serializer=iguana_dot_v1_dot_iguana__pb2.PingResponse.SerializeToString,
            ),
            'VisitorEvent': grpc.unary_unary_rpc_method_handler(
                    servicer.VisitorEvent,
                    request_deserializer=iguana_dot_v1_dot_iguana__pb2.VisitorEventRequest.FromString,
                    response_serializer=iguana_dot_v1_dot_iguana__pb2.VisitorEventResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'iguana.v1.IguanaService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class IguanaService(object):
    """import "google/type/datetime.proto";


    """

    @staticmethod
    def Ping(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/iguana.v1.IguanaService/Ping',
            iguana_dot_v1_dot_iguana__pb2.PingRequest.SerializeToString,
            iguana_dot_v1_dot_iguana__pb2.PingResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def VisitorEvent(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/iguana.v1.IguanaService/VisitorEvent',
            iguana_dot_v1_dot_iguana__pb2.VisitorEventRequest.SerializeToString,
            iguana_dot_v1_dot_iguana__pb2.VisitorEventResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
