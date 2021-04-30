// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: client.proto

package gonetlockprotobuf

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

type ClientAttachRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *ClientAttachRequest) Reset() {
	*x = ClientAttachRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAttachRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAttachRequest) ProtoMessage() {}

func (x *ClientAttachRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAttachRequest.ProtoReflect.Descriptor instead.
func (*ClientAttachRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *ClientAttachRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type ClientAttachResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *ClientAttachResponse) Reset() {
	*x = ClientAttachResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAttachResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAttachResponse) ProtoMessage() {}

func (x *ClientAttachResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAttachResponse.ProtoReflect.Descriptor instead.
func (*ClientAttachResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

func (x *ClientAttachResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type ClientDetachRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *ClientDetachRequest) Reset() {
	*x = ClientDetachRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientDetachRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientDetachRequest) ProtoMessage() {}

func (x *ClientDetachRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientDetachRequest.ProtoReflect.Descriptor instead.
func (*ClientDetachRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{2}
}

func (x *ClientDetachRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type ClientDetachResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *ClientDetachResponse) Reset() {
	*x = ClientDetachResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientDetachResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientDetachResponse) ProtoMessage() {}

func (x *ClientDetachResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientDetachResponse.ProtoReflect.Descriptor instead.
func (*ClientDetachResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{3}
}

func (x *ClientDetachResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type ClientCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientCreateRequest) Reset() {
	*x = ClientCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientCreateRequest) ProtoMessage() {}

func (x *ClientCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientCreateRequest.ProtoReflect.Descriptor instead.
func (*ClientCreateRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{4}
}

type ClientCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *ClientCreateResponse) Reset() {
	*x = ClientCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientCreateResponse) ProtoMessage() {}

func (x *ClientCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientCreateResponse.ProtoReflect.Descriptor instead.
func (*ClientCreateResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{5}
}

func (x *ClientCreateResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type ClientGetDefaultIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientGetDefaultIdRequest) Reset() {
	*x = ClientGetDefaultIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientGetDefaultIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientGetDefaultIdRequest) ProtoMessage() {}

func (x *ClientGetDefaultIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientGetDefaultIdRequest.ProtoReflect.Descriptor instead.
func (*ClientGetDefaultIdRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{6}
}

type ClientGetDefaultIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *ClientGetDefaultIdResponse) Reset() {
	*x = ClientGetDefaultIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientGetDefaultIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientGetDefaultIdResponse) ProtoMessage() {}

func (x *ClientGetDefaultIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientGetDefaultIdResponse.ProtoReflect.Descriptor instead.
func (*ClientGetDefaultIdResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{7}
}

func (x *ClientGetDefaultIdResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type ClientGetAttachedIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientGetAttachedIdsRequest) Reset() {
	*x = ClientGetAttachedIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientGetAttachedIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientGetAttachedIdsRequest) ProtoMessage() {}

func (x *ClientGetAttachedIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientGetAttachedIdsRequest.ProtoReflect.Descriptor instead.
func (*ClientGetAttachedIdsRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{8}
}

type ClientGetAttachedIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIds []string `protobuf:"bytes,1,rep,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
}

func (x *ClientGetAttachedIdsResponse) Reset() {
	*x = ClientGetAttachedIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientGetAttachedIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientGetAttachedIdsResponse) ProtoMessage() {}

func (x *ClientGetAttachedIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientGetAttachedIdsResponse.ProtoReflect.Descriptor instead.
func (*ClientGetAttachedIdsResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{9}
}

func (x *ClientGetAttachedIdsResponse) GetClientIds() []string {
	if x != nil {
		return x.ClientIds
	}
	return nil
}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x67, 0x6f, 0x6e, 0x65, 0x74, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x32, 0x0a, 0x13, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x33, 0x0a,
	0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x32, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x33, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x1a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x1d, 0x0a, 0x1b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d,
	0x0a, 0x1c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6e, 0x65,
	0x74, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x67, 0x6f, 0x6e, 0x65, 0x74, 0x6c, 0x6f, 0x63, 0x6b, 0x2f,
	0x67, 0x6f, 0x6e, 0x65, 0x74, 0x6c, 0x6f, 0x63, 0x6b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_client_proto_goTypes = []interface{}{
	(*ClientAttachRequest)(nil),          // 0: gonetlock.ClientAttachRequest
	(*ClientAttachResponse)(nil),         // 1: gonetlock.ClientAttachResponse
	(*ClientDetachRequest)(nil),          // 2: gonetlock.ClientDetachRequest
	(*ClientDetachResponse)(nil),         // 3: gonetlock.ClientDetachResponse
	(*ClientCreateRequest)(nil),          // 4: gonetlock.ClientCreateRequest
	(*ClientCreateResponse)(nil),         // 5: gonetlock.ClientCreateResponse
	(*ClientGetDefaultIdRequest)(nil),    // 6: gonetlock.ClientGetDefaultIdRequest
	(*ClientGetDefaultIdResponse)(nil),   // 7: gonetlock.ClientGetDefaultIdResponse
	(*ClientGetAttachedIdsRequest)(nil),  // 8: gonetlock.ClientGetAttachedIdsRequest
	(*ClientGetAttachedIdsResponse)(nil), // 9: gonetlock.ClientGetAttachedIdsResponse
}
var file_client_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAttachRequest); i {
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
		file_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAttachResponse); i {
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
		file_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientDetachRequest); i {
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
		file_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientDetachResponse); i {
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
		file_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientCreateRequest); i {
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
		file_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientCreateResponse); i {
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
		file_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientGetDefaultIdRequest); i {
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
		file_client_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientGetDefaultIdResponse); i {
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
		file_client_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientGetAttachedIdsRequest); i {
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
		file_client_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientGetAttachedIdsResponse); i {
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
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}