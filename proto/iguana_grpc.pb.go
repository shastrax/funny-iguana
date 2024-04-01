// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: iguana.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Iguana_CognitoEvent_FullMethodName = "/proto.Iguana/CognitoEvent"
	Iguana_Ping_FullMethodName         = "/proto.Iguana/Ping"
	Iguana_SelectNote_FullMethodName   = "/proto.Iguana/SelectNote"
	Iguana_SubmitNote_FullMethodName   = "/proto.Iguana/SubmitNote"
	Iguana_VisitorEvent_FullMethodName = "/proto.Iguana/VisitorEvent"
)

// IguanaClient is the client API for Iguana service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IguanaClient interface {
	CognitoEvent(ctx context.Context, in *CognitoEventRequest, opts ...grpc.CallOption) (*CognitoEventResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	SelectNote(ctx context.Context, in *SelectNoteRequest, opts ...grpc.CallOption) (*SelectNoteResponse, error)
	SubmitNote(ctx context.Context, in *SubmitNoteRequest, opts ...grpc.CallOption) (*SubmitNoteResponse, error)
	VisitorEvent(ctx context.Context, in *VisitorEventRequest, opts ...grpc.CallOption) (*VisitorEventResponse, error)
}

type iguanaClient struct {
	cc grpc.ClientConnInterface
}

func NewIguanaClient(cc grpc.ClientConnInterface) IguanaClient {
	return &iguanaClient{cc}
}

func (c *iguanaClient) CognitoEvent(ctx context.Context, in *CognitoEventRequest, opts ...grpc.CallOption) (*CognitoEventResponse, error) {
	out := new(CognitoEventResponse)
	err := c.cc.Invoke(ctx, Iguana_CognitoEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iguanaClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, Iguana_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iguanaClient) SelectNote(ctx context.Context, in *SelectNoteRequest, opts ...grpc.CallOption) (*SelectNoteResponse, error) {
	out := new(SelectNoteResponse)
	err := c.cc.Invoke(ctx, Iguana_SelectNote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iguanaClient) SubmitNote(ctx context.Context, in *SubmitNoteRequest, opts ...grpc.CallOption) (*SubmitNoteResponse, error) {
	out := new(SubmitNoteResponse)
	err := c.cc.Invoke(ctx, Iguana_SubmitNote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iguanaClient) VisitorEvent(ctx context.Context, in *VisitorEventRequest, opts ...grpc.CallOption) (*VisitorEventResponse, error) {
	out := new(VisitorEventResponse)
	err := c.cc.Invoke(ctx, Iguana_VisitorEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IguanaServer is the server API for Iguana service.
// All implementations must embed UnimplementedIguanaServer
// for forward compatibility
type IguanaServer interface {
	CognitoEvent(context.Context, *CognitoEventRequest) (*CognitoEventResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	SelectNote(context.Context, *SelectNoteRequest) (*SelectNoteResponse, error)
	SubmitNote(context.Context, *SubmitNoteRequest) (*SubmitNoteResponse, error)
	VisitorEvent(context.Context, *VisitorEventRequest) (*VisitorEventResponse, error)
	mustEmbedUnimplementedIguanaServer()
}

// UnimplementedIguanaServer must be embedded to have forward compatible implementations.
type UnimplementedIguanaServer struct {
}

func (UnimplementedIguanaServer) CognitoEvent(context.Context, *CognitoEventRequest) (*CognitoEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CognitoEvent not implemented")
}
func (UnimplementedIguanaServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedIguanaServer) SelectNote(context.Context, *SelectNoteRequest) (*SelectNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectNote not implemented")
}
func (UnimplementedIguanaServer) SubmitNote(context.Context, *SubmitNoteRequest) (*SubmitNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitNote not implemented")
}
func (UnimplementedIguanaServer) VisitorEvent(context.Context, *VisitorEventRequest) (*VisitorEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VisitorEvent not implemented")
}
func (UnimplementedIguanaServer) mustEmbedUnimplementedIguanaServer() {}

// UnsafeIguanaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IguanaServer will
// result in compilation errors.
type UnsafeIguanaServer interface {
	mustEmbedUnimplementedIguanaServer()
}

func RegisterIguanaServer(s grpc.ServiceRegistrar, srv IguanaServer) {
	s.RegisterService(&Iguana_ServiceDesc, srv)
}

func _Iguana_CognitoEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CognitoEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IguanaServer).CognitoEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Iguana_CognitoEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IguanaServer).CognitoEvent(ctx, req.(*CognitoEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iguana_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IguanaServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Iguana_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IguanaServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iguana_SelectNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IguanaServer).SelectNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Iguana_SelectNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IguanaServer).SelectNote(ctx, req.(*SelectNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iguana_SubmitNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IguanaServer).SubmitNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Iguana_SubmitNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IguanaServer).SubmitNote(ctx, req.(*SubmitNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iguana_VisitorEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VisitorEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IguanaServer).VisitorEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Iguana_VisitorEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IguanaServer).VisitorEvent(ctx, req.(*VisitorEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Iguana_ServiceDesc is the grpc.ServiceDesc for Iguana service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Iguana_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Iguana",
	HandlerType: (*IguanaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CognitoEvent",
			Handler:    _Iguana_CognitoEvent_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Iguana_Ping_Handler,
		},
		{
			MethodName: "SelectNote",
			Handler:    _Iguana_SelectNote_Handler,
		},
		{
			MethodName: "SubmitNote",
			Handler:    _Iguana_SubmitNote_Handler,
		},
		{
			MethodName: "VisitorEvent",
			Handler:    _Iguana_VisitorEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iguana.proto",
}
