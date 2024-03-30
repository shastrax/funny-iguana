// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: iguana.proto

package proto

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

	Source    string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	TimeStamp int64  `protobuf:"varint,2,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iguana_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iguana_proto_msgTypes[0]
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
	return file_iguana_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *PingRequest) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
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
		mi := &file_iguana_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iguana_proto_msgTypes[1]
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
	return file_iguana_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type SelectNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SelectNoteRequest) Reset() {
	*x = SelectNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iguana_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectNoteRequest) ProtoMessage() {}

func (x *SelectNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iguana_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectNoteRequest.ProtoReflect.Descriptor instead.
func (*SelectNoteRequest) Descriptor() ([]byte, []int) {
	return file_iguana_proto_rawDescGZIP(), []int{2}
}

func (x *SelectNoteRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *SelectNoteRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SelectNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SelectNoteResponse) Reset() {
	*x = SelectNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iguana_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectNoteResponse) ProtoMessage() {}

func (x *SelectNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iguana_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectNoteResponse.ProtoReflect.Descriptor instead.
func (*SelectNoteResponse) Descriptor() ([]byte, []int) {
	return file_iguana_proto_rawDescGZIP(), []int{3}
}

func (x *SelectNoteResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *SelectNoteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SubmitNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // requestor
}

func (x *SubmitNoteRequest) Reset() {
	*x = SubmitNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iguana_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitNoteRequest) ProtoMessage() {}

func (x *SubmitNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iguana_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitNoteRequest.ProtoReflect.Descriptor instead.
func (*SubmitNoteRequest) Descriptor() ([]byte, []int) {
	return file_iguana_proto_rawDescGZIP(), []int{4}
}

func (x *SubmitNoteRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type SubmitNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // requestor
}

func (x *SubmitNoteResponse) Reset() {
	*x = SubmitNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iguana_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitNoteResponse) ProtoMessage() {}

func (x *SubmitNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iguana_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitNoteResponse.ProtoReflect.Descriptor instead.
func (*SubmitNoteResponse) Descriptor() ([]byte, []int) {
	return file_iguana_proto_rawDescGZIP(), []int{5}
}

func (x *SubmitNoteResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type SubmitVisitorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	TimeStamp int64  `protobuf:"varint,2,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Url       string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *SubmitVisitorRequest) Reset() {
	*x = SubmitVisitorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iguana_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitVisitorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitVisitorRequest) ProtoMessage() {}

func (x *SubmitVisitorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iguana_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitVisitorRequest.ProtoReflect.Descriptor instead.
func (*SubmitVisitorRequest) Descriptor() ([]byte, []int) {
	return file_iguana_proto_rawDescGZIP(), []int{6}
}

func (x *SubmitVisitorRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SubmitVisitorRequest) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

func (x *SubmitVisitorRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type SubmitVisitorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SubmitVisitorResponse) Reset() {
	*x = SubmitVisitorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iguana_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitVisitorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitVisitorResponse) ProtoMessage() {}

func (x *SubmitVisitorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iguana_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitVisitorResponse.ProtoReflect.Descriptor instead.
func (*SubmitVisitorResponse) Descriptor() ([]byte, []int) {
	return file_iguana_proto_rawDescGZIP(), []int{7}
}

func (x *SubmitVisitorResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_iguana_proto protoreflect.FileDescriptor

var file_iguana_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x26, 0x0a, 0x0c, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x4a, 0x0a, 0x11, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x4b, 0x0a, 0x12, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x11,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x31,
	0x0a, 0x12, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x61, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x56, 0x69, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x2f, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x56, 0x69,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x93, 0x02, 0x0a, 0x06, 0x49, 0x67, 0x75, 0x61, 0x6e, 0x61,
	0x12, 0x31, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x74,
	0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x56, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x56, 0x69, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x56, 0x69, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x73, 0x74, 0x72,
	0x61, 0x78, 0x2f, 0x66, 0x75, 0x6e, 0x6e, 0x79, 0x2d, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iguana_proto_rawDescOnce sync.Once
	file_iguana_proto_rawDescData = file_iguana_proto_rawDesc
)

func file_iguana_proto_rawDescGZIP() []byte {
	file_iguana_proto_rawDescOnce.Do(func() {
		file_iguana_proto_rawDescData = protoimpl.X.CompressGZIP(file_iguana_proto_rawDescData)
	})
	return file_iguana_proto_rawDescData
}

var file_iguana_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_iguana_proto_goTypes = []interface{}{
	(*PingRequest)(nil),           // 0: proto.PingRequest
	(*PingResponse)(nil),          // 1: proto.PingResponse
	(*SelectNoteRequest)(nil),     // 2: proto.SelectNoteRequest
	(*SelectNoteResponse)(nil),    // 3: proto.SelectNoteResponse
	(*SubmitNoteRequest)(nil),     // 4: proto.SubmitNoteRequest
	(*SubmitNoteResponse)(nil),    // 5: proto.SubmitNoteResponse
	(*SubmitVisitorRequest)(nil),  // 6: proto.SubmitVisitorRequest
	(*SubmitVisitorResponse)(nil), // 7: proto.SubmitVisitorResponse
}
var file_iguana_proto_depIdxs = []int32{
	0, // 0: proto.Iguana.Ping:input_type -> proto.PingRequest
	2, // 1: proto.Iguana.SelectNote:input_type -> proto.SelectNoteRequest
	4, // 2: proto.Iguana.SubmitNote:input_type -> proto.SubmitNoteRequest
	6, // 3: proto.Iguana.SubmitVisitor:input_type -> proto.SubmitVisitorRequest
	1, // 4: proto.Iguana.Ping:output_type -> proto.PingResponse
	3, // 5: proto.Iguana.SelectNote:output_type -> proto.SelectNoteResponse
	5, // 6: proto.Iguana.SubmitNote:output_type -> proto.SubmitNoteResponse
	7, // 7: proto.Iguana.SubmitVisitor:output_type -> proto.SubmitVisitorResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_iguana_proto_init() }
func file_iguana_proto_init() {
	if File_iguana_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iguana_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_iguana_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_iguana_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectNoteRequest); i {
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
		file_iguana_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectNoteResponse); i {
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
		file_iguana_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitNoteRequest); i {
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
		file_iguana_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitNoteResponse); i {
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
		file_iguana_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitVisitorRequest); i {
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
		file_iguana_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitVisitorResponse); i {
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
			RawDescriptor: file_iguana_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iguana_proto_goTypes,
		DependencyIndexes: file_iguana_proto_depIdxs,
		MessageInfos:      file_iguana_proto_msgTypes,
	}.Build()
	File_iguana_proto = out.File
	file_iguana_proto_rawDesc = nil
	file_iguana_proto_goTypes = nil
	file_iguana_proto_depIdxs = nil
}
