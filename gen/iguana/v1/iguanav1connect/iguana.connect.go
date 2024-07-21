// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: iguana/v1/iguana.proto

package iguanav1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/shastrax/funny-iguana/gen/iguana/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// IguanaServiceName is the fully-qualified name of the IguanaService service.
	IguanaServiceName = "iguana.v1.IguanaService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IguanaServiceCognitoEventProcedure is the fully-qualified name of the IguanaService's
	// CognitoEvent RPC.
	IguanaServiceCognitoEventProcedure = "/iguana.v1.IguanaService/CognitoEvent"
	// IguanaServicePingProcedure is the fully-qualified name of the IguanaService's Ping RPC.
	IguanaServicePingProcedure = "/iguana.v1.IguanaService/Ping"
	// IguanaServiceVisitorEventProcedure is the fully-qualified name of the IguanaService's
	// VisitorEvent RPC.
	IguanaServiceVisitorEventProcedure = "/iguana.v1.IguanaService/VisitorEvent"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	iguanaServiceServiceDescriptor            = v1.File_iguana_v1_iguana_proto.Services().ByName("IguanaService")
	iguanaServiceCognitoEventMethodDescriptor = iguanaServiceServiceDescriptor.Methods().ByName("CognitoEvent")
	iguanaServicePingMethodDescriptor         = iguanaServiceServiceDescriptor.Methods().ByName("Ping")
	iguanaServiceVisitorEventMethodDescriptor = iguanaServiceServiceDescriptor.Methods().ByName("VisitorEvent")
)

// IguanaServiceClient is a client for the iguana.v1.IguanaService service.
type IguanaServiceClient interface {
	CognitoEvent(context.Context, *connect.Request[v1.CognitoEventRequest]) (*connect.Response[v1.CognitoEventResponse], error)
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
	// rpc RandomNote(RandomNoteRequest) returns (RandomNoteResponse) {}
	VisitorEvent(context.Context, *connect.Request[v1.VisitorEventRequest]) (*connect.Response[v1.VisitorEventResponse], error)
}

// NewIguanaServiceClient constructs a client for the iguana.v1.IguanaService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIguanaServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) IguanaServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &iguanaServiceClient{
		cognitoEvent: connect.NewClient[v1.CognitoEventRequest, v1.CognitoEventResponse](
			httpClient,
			baseURL+IguanaServiceCognitoEventProcedure,
			connect.WithSchema(iguanaServiceCognitoEventMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ping: connect.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+IguanaServicePingProcedure,
			connect.WithSchema(iguanaServicePingMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		visitorEvent: connect.NewClient[v1.VisitorEventRequest, v1.VisitorEventResponse](
			httpClient,
			baseURL+IguanaServiceVisitorEventProcedure,
			connect.WithSchema(iguanaServiceVisitorEventMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// iguanaServiceClient implements IguanaServiceClient.
type iguanaServiceClient struct {
	cognitoEvent *connect.Client[v1.CognitoEventRequest, v1.CognitoEventResponse]
	ping         *connect.Client[v1.PingRequest, v1.PingResponse]
	visitorEvent *connect.Client[v1.VisitorEventRequest, v1.VisitorEventResponse]
}

// CognitoEvent calls iguana.v1.IguanaService.CognitoEvent.
func (c *iguanaServiceClient) CognitoEvent(ctx context.Context, req *connect.Request[v1.CognitoEventRequest]) (*connect.Response[v1.CognitoEventResponse], error) {
	return c.cognitoEvent.CallUnary(ctx, req)
}

// Ping calls iguana.v1.IguanaService.Ping.
func (c *iguanaServiceClient) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// VisitorEvent calls iguana.v1.IguanaService.VisitorEvent.
func (c *iguanaServiceClient) VisitorEvent(ctx context.Context, req *connect.Request[v1.VisitorEventRequest]) (*connect.Response[v1.VisitorEventResponse], error) {
	return c.visitorEvent.CallUnary(ctx, req)
}

// IguanaServiceHandler is an implementation of the iguana.v1.IguanaService service.
type IguanaServiceHandler interface {
	CognitoEvent(context.Context, *connect.Request[v1.CognitoEventRequest]) (*connect.Response[v1.CognitoEventResponse], error)
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
	// rpc RandomNote(RandomNoteRequest) returns (RandomNoteResponse) {}
	VisitorEvent(context.Context, *connect.Request[v1.VisitorEventRequest]) (*connect.Response[v1.VisitorEventResponse], error)
}

// NewIguanaServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIguanaServiceHandler(svc IguanaServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	iguanaServiceCognitoEventHandler := connect.NewUnaryHandler(
		IguanaServiceCognitoEventProcedure,
		svc.CognitoEvent,
		connect.WithSchema(iguanaServiceCognitoEventMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	iguanaServicePingHandler := connect.NewUnaryHandler(
		IguanaServicePingProcedure,
		svc.Ping,
		connect.WithSchema(iguanaServicePingMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	iguanaServiceVisitorEventHandler := connect.NewUnaryHandler(
		IguanaServiceVisitorEventProcedure,
		svc.VisitorEvent,
		connect.WithSchema(iguanaServiceVisitorEventMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/iguana.v1.IguanaService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IguanaServiceCognitoEventProcedure:
			iguanaServiceCognitoEventHandler.ServeHTTP(w, r)
		case IguanaServicePingProcedure:
			iguanaServicePingHandler.ServeHTTP(w, r)
		case IguanaServiceVisitorEventProcedure:
			iguanaServiceVisitorEventHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIguanaServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIguanaServiceHandler struct{}

func (UnimplementedIguanaServiceHandler) CognitoEvent(context.Context, *connect.Request[v1.CognitoEventRequest]) (*connect.Response[v1.CognitoEventResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("iguana.v1.IguanaService.CognitoEvent is not implemented"))
}

func (UnimplementedIguanaServiceHandler) Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("iguana.v1.IguanaService.Ping is not implemented"))
}

func (UnimplementedIguanaServiceHandler) VisitorEvent(context.Context, *connect.Request[v1.VisitorEventRequest]) (*connect.Response[v1.VisitorEventResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("iguana.v1.IguanaService.VisitorEvent is not implemented"))
}
