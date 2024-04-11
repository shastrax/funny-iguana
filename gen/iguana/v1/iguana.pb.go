// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: iguana/v1/iguana.proto

//option go_package = "github.com/shastrax/funny-iguana/proto";

package iguanav1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iguana_v1_iguana_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iguana_v1_iguana_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_iguana_v1_iguana_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iguana_v1_iguana_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iguana_v1_iguana_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_iguana_v1_iguana_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_iguana_v1_iguana_proto protoreflect.FileDescriptor

var file_iguana_v1_iguana_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x67, 0x75, 0x61,
	0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x61,
	0x2e, 0x76, 0x31, 0x22, 0x25, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x26, 0x0a, 0x0c, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x32, 0x4a, 0x0a, 0x0d, 0x49, 0x67, 0x75, 0x61, 0x6e, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x69, 0x67,
	0x75, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x9a,
	0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31,
	0x42, 0x0b, 0x49, 0x67, 0x75, 0x61, 0x6e, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x73,
	0x74, 0x72, 0x61, 0x78, 0x2f, 0x66, 0x75, 0x6e, 0x6e, 0x79, 0x2d, 0x69, 0x67, 0x75, 0x61, 0x6e,
	0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x61, 0x2f, 0x76, 0x31, 0x3b,
	0x69, 0x67, 0x75, 0x61, 0x6e, 0x61, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02,
	0x09, 0x49, 0x67, 0x75, 0x61, 0x6e, 0x61, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x49, 0x67, 0x75,
	0x61, 0x6e, 0x61, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x49, 0x67, 0x75, 0x61, 0x6e, 0x61, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0a, 0x49, 0x67, 0x75, 0x61, 0x6e, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_iguana_v1_iguana_proto_rawDescOnce sync.Once
	file_iguana_v1_iguana_proto_rawDescData = file_iguana_v1_iguana_proto_rawDesc
)

func file_iguana_v1_iguana_proto_rawDescGZIP() []byte {
	file_iguana_v1_iguana_proto_rawDescOnce.Do(func() {
		file_iguana_v1_iguana_proto_rawDescData = protoimpl.X.CompressGZIP(file_iguana_v1_iguana_proto_rawDescData)
	})
	return file_iguana_v1_iguana_proto_rawDescData
}

var file_iguana_v1_iguana_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_iguana_v1_iguana_proto_goTypes = []interface{}{
	(*PingRequest)(nil),  // 0: iguana.v1.PingRequest
	(*PingResponse)(nil), // 1: iguana.v1.PingResponse
}
var file_iguana_v1_iguana_proto_depIdxs = []int32{
	0, // 0: iguana.v1.IguanaService.Ping:input_type -> iguana.v1.PingRequest
	1, // 1: iguana.v1.IguanaService.Ping:output_type -> iguana.v1.PingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_iguana_v1_iguana_proto_init() }
func file_iguana_v1_iguana_proto_init() {
	if File_iguana_v1_iguana_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iguana_v1_iguana_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iguana_v1_iguana_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_iguana_v1_iguana_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iguana_v1_iguana_proto_goTypes,
		DependencyIndexes: file_iguana_v1_iguana_proto_depIdxs,
		MessageInfos:      file_iguana_v1_iguana_proto_msgTypes,
	}.Build()
	File_iguana_v1_iguana_proto = out.File
	file_iguana_v1_iguana_proto_rawDesc = nil
	file_iguana_v1_iguana_proto_goTypes = nil
	file_iguana_v1_iguana_proto_depIdxs = nil
}