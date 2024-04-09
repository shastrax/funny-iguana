// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: iguana/v1/iguana.proto

package iguanav1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/shastrax/funny-iguana/gen/iguana/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

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
	// IguanaServiceSelectNoteProcedure is the fully-qualified name of the IguanaService's SelectNote
	// RPC.
	IguanaServiceSelectNoteProcedure = "/iguana.v1.IguanaService/SelectNote"
	// IguanaServiceSubmitNoteProcedure is the fully-qualified name of the IguanaService's SubmitNote
	// RPC.
	IguanaServiceSubmitNoteProcedure = "/iguana.v1.IguanaService/SubmitNote"
	// IguanaServiceVisitorEventProcedure is the fully-qualified name of the IguanaService's
	// VisitorEvent RPC.
	IguanaServiceVisitorEventProcedure = "/iguana.v1.IguanaService/VisitorEvent"
)

// IguanaServiceClient is a client for the iguana.v1.IguanaService service.
type IguanaServiceClient interface {
	CognitoEvent(context.Context, *connect_go.Request[v1.CognitoEventRequest]) (*connect_go.Response[v1.CognitoEventResponse], error)
	Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error)
	SelectNote(context.Context, *connect_go.Request[v1.SelectNoteRequest]) (*connect_go.Response[v1.SelectNoteResponse], error)
	SubmitNote(context.Context, *connect_go.Request[v1.SubmitNoteRequest]) (*connect_go.Response[v1.SubmitNoteResponse], error)
	VisitorEvent(context.Context, *connect_go.Request[v1.VisitorEventRequest]) (*connect_go.Response[v1.VisitorEventResponse], error)
}

// NewIguanaServiceClient constructs a client for the iguana.v1.IguanaService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIguanaServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) IguanaServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &iguanaServiceClient{
		cognitoEvent: connect_go.NewClient[v1.CognitoEventRequest, v1.CognitoEventResponse](
			httpClient,
			baseURL+IguanaServiceCognitoEventProcedure,
			opts...,
		),
		ping: connect_go.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+IguanaServicePingProcedure,
			opts...,
		),
		selectNote: connect_go.NewClient[v1.SelectNoteRequest, v1.SelectNoteResponse](
			httpClient,
			baseURL+IguanaServiceSelectNoteProcedure,
			opts...,
		),
		submitNote: connect_go.NewClient[v1.SubmitNoteRequest, v1.SubmitNoteResponse](
			httpClient,
			baseURL+IguanaServiceSubmitNoteProcedure,
			opts...,
		),
		visitorEvent: connect_go.NewClient[v1.VisitorEventRequest, v1.VisitorEventResponse](
			httpClient,
			baseURL+IguanaServiceVisitorEventProcedure,
			opts...,
		),
	}
}

// iguanaServiceClient implements IguanaServiceClient.
type iguanaServiceClient struct {
	cognitoEvent *connect_go.Client[v1.CognitoEventRequest, v1.CognitoEventResponse]
	ping         *connect_go.Client[v1.PingRequest, v1.PingResponse]
	selectNote   *connect_go.Client[v1.SelectNoteRequest, v1.SelectNoteResponse]
	submitNote   *connect_go.Client[v1.SubmitNoteRequest, v1.SubmitNoteResponse]
	visitorEvent *connect_go.Client[v1.VisitorEventRequest, v1.VisitorEventResponse]
}

// CognitoEvent calls iguana.v1.IguanaService.CognitoEvent.
func (c *iguanaServiceClient) CognitoEvent(ctx context.Context, req *connect_go.Request[v1.CognitoEventRequest]) (*connect_go.Response[v1.CognitoEventResponse], error) {
	return c.cognitoEvent.CallUnary(ctx, req)
}

// Ping calls iguana.v1.IguanaService.Ping.
func (c *iguanaServiceClient) Ping(ctx context.Context, req *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// SelectNote calls iguana.v1.IguanaService.SelectNote.
func (c *iguanaServiceClient) SelectNote(ctx context.Context, req *connect_go.Request[v1.SelectNoteRequest]) (*connect_go.Response[v1.SelectNoteResponse], error) {
	return c.selectNote.CallUnary(ctx, req)
}

// SubmitNote calls iguana.v1.IguanaService.SubmitNote.
func (c *iguanaServiceClient) SubmitNote(ctx context.Context, req *connect_go.Request[v1.SubmitNoteRequest]) (*connect_go.Response[v1.SubmitNoteResponse], error) {
	return c.submitNote.CallUnary(ctx, req)
}

// VisitorEvent calls iguana.v1.IguanaService.VisitorEvent.
func (c *iguanaServiceClient) VisitorEvent(ctx context.Context, req *connect_go.Request[v1.VisitorEventRequest]) (*connect_go.Response[v1.VisitorEventResponse], error) {
	return c.visitorEvent.CallUnary(ctx, req)
}

// IguanaServiceHandler is an implementation of the iguana.v1.IguanaService service.
type IguanaServiceHandler interface {
	CognitoEvent(context.Context, *connect_go.Request[v1.CognitoEventRequest]) (*connect_go.Response[v1.CognitoEventResponse], error)
	Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error)
	SelectNote(context.Context, *connect_go.Request[v1.SelectNoteRequest]) (*connect_go.Response[v1.SelectNoteResponse], error)
	SubmitNote(context.Context, *connect_go.Request[v1.SubmitNoteRequest]) (*connect_go.Response[v1.SubmitNoteResponse], error)
	VisitorEvent(context.Context, *connect_go.Request[v1.VisitorEventRequest]) (*connect_go.Response[v1.VisitorEventResponse], error)
}

// NewIguanaServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIguanaServiceHandler(svc IguanaServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	iguanaServiceCognitoEventHandler := connect_go.NewUnaryHandler(
		IguanaServiceCognitoEventProcedure,
		svc.CognitoEvent,
		opts...,
	)
	iguanaServicePingHandler := connect_go.NewUnaryHandler(
		IguanaServicePingProcedure,
		svc.Ping,
		opts...,
	)
	iguanaServiceSelectNoteHandler := connect_go.NewUnaryHandler(
		IguanaServiceSelectNoteProcedure,
		svc.SelectNote,
		opts...,
	)
	iguanaServiceSubmitNoteHandler := connect_go.NewUnaryHandler(
		IguanaServiceSubmitNoteProcedure,
		svc.SubmitNote,
		opts...,
	)
	iguanaServiceVisitorEventHandler := connect_go.NewUnaryHandler(
		IguanaServiceVisitorEventProcedure,
		svc.VisitorEvent,
		opts...,
	)
	return "/iguana.v1.IguanaService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IguanaServiceCognitoEventProcedure:
			iguanaServiceCognitoEventHandler.ServeHTTP(w, r)
		case IguanaServicePingProcedure:
			iguanaServicePingHandler.ServeHTTP(w, r)
		case IguanaServiceSelectNoteProcedure:
			iguanaServiceSelectNoteHandler.ServeHTTP(w, r)
		case IguanaServiceSubmitNoteProcedure:
			iguanaServiceSubmitNoteHandler.ServeHTTP(w, r)
		case IguanaServiceVisitorEventProcedure:
			iguanaServiceVisitorEventHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIguanaServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIguanaServiceHandler struct{}

func (UnimplementedIguanaServiceHandler) CognitoEvent(context.Context, *connect_go.Request[v1.CognitoEventRequest]) (*connect_go.Response[v1.CognitoEventResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("iguana.v1.IguanaService.CognitoEvent is not implemented"))
}

func (UnimplementedIguanaServiceHandler) Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("iguana.v1.IguanaService.Ping is not implemented"))
}

func (UnimplementedIguanaServiceHandler) SelectNote(context.Context, *connect_go.Request[v1.SelectNoteRequest]) (*connect_go.Response[v1.SelectNoteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("iguana.v1.IguanaService.SelectNote is not implemented"))
}

func (UnimplementedIguanaServiceHandler) SubmitNote(context.Context, *connect_go.Request[v1.SubmitNoteRequest]) (*connect_go.Response[v1.SubmitNoteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("iguana.v1.IguanaService.SubmitNote is not implemented"))
}

func (UnimplementedIguanaServiceHandler) VisitorEvent(context.Context, *connect_go.Request[v1.VisitorEventRequest]) (*connect_go.Response[v1.VisitorEventResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("iguana.v1.IguanaService.VisitorEvent is not implemented"))
}