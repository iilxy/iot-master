// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: plugin/protocol.proto

package protocol

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok    bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_plugin_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Open struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // options
}

func (x *Open) Reset() {
	*x = Open{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Open) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Open) ProtoMessage() {}

func (x *Open) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Open.ProtoReflect.Descriptor instead.
func (*Open) Descriptor() ([]byte, []int) {
	return file_plugin_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *Open) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Close struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Close) Reset() {
	*x = Close{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Close) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Close) ProtoMessage() {}

func (x *Close) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Close.ProtoReflect.Descriptor instead.
func (*Close) Descriptor() ([]byte, []int) {
	return file_plugin_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *Close) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Protocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Protocol) Reset() {
	*x = Protocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Protocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Protocol) ProtoMessage() {}

func (x *Protocol) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_protocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Protocol.ProtoReflect.Descriptor instead.
func (*Protocol) Descriptor() ([]byte, []int) {
	return file_plugin_protocol_proto_rawDescGZIP(), []int{3}
}

func (x *Protocol) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Protocol) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Read struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Length  uint32 `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *Read) Reset() {
	*x = Read{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_protocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Read) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Read) ProtoMessage() {}

func (x *Read) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_protocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Read.ProtoReflect.Descriptor instead.
func (*Read) Descriptor() ([]byte, []int) {
	return file_plugin_protocol_proto_rawDescGZIP(), []int{4}
}

func (x *Read) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Read) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok    bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_protocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_protocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_plugin_protocol_proto_rawDescGZIP(), []int{5}
}

func (x *Result) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *Result) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Result) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Write struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Write) Reset() {
	*x = Write{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_protocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Write) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Write) ProtoMessage() {}

func (x *Write) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_protocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Write.ProtoReflect.Descriptor instead.
func (*Write) Descriptor() ([]byte, []int) {
	return file_plugin_protocol_proto_rawDescGZIP(), []int{6}
}

func (x *Write) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Write) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_plugin_protocol_proto protoreflect.FileDescriptor

var file_plugin_protocol_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x16, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x38, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x04, 0x52,
	0x65, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x42, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x05, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0xfa, 0x01, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x34, 0x0a,
	0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x1a, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2e, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x1a, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2a, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a,
	0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a,
	0x11, 0x2e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_protocol_proto_rawDescOnce sync.Once
	file_plugin_protocol_proto_rawDescData = file_plugin_protocol_proto_rawDesc
)

func file_plugin_protocol_proto_rawDescGZIP() []byte {
	file_plugin_protocol_proto_rawDescOnce.Do(func() {
		file_plugin_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_protocol_proto_rawDescData)
	})
	return file_plugin_protocol_proto_rawDescData
}

var file_plugin_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_plugin_protocol_proto_goTypes = []interface{}{
	(*Response)(nil), // 0: protocol.Response
	(*Open)(nil),     // 1: protocol.Open
	(*Close)(nil),    // 2: protocol.Close
	(*Protocol)(nil), // 3: protocol.Protocol
	(*Read)(nil),     // 4: protocol.Read
	(*Result)(nil),   // 5: protocol.Result
	(*Write)(nil),    // 6: protocol.Write
}
var file_plugin_protocol_proto_depIdxs = []int32{
	3, // 0: protocol.protocol.register:input_type -> protocol.Protocol
	1, // 1: protocol.protocol.open:input_type -> protocol.Open
	2, // 2: protocol.protocol.close:input_type -> protocol.Close
	4, // 3: protocol.protocol.read:input_type -> protocol.Read
	6, // 4: protocol.protocol.write:input_type -> protocol.Write
	0, // 5: protocol.protocol.register:output_type -> protocol.Response
	0, // 6: protocol.protocol.open:output_type -> protocol.Response
	0, // 7: protocol.protocol.close:output_type -> protocol.Response
	5, // 8: protocol.protocol.read:output_type -> protocol.Result
	0, // 9: protocol.protocol.write:output_type -> protocol.Response
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_plugin_protocol_proto_init() }
func file_plugin_protocol_proto_init() {
	if File_plugin_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_plugin_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Open); i {
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
		file_plugin_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Close); i {
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
		file_plugin_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Protocol); i {
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
		file_plugin_protocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Read); i {
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
		file_plugin_protocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_plugin_protocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Write); i {
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
			RawDescriptor: file_plugin_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_protocol_proto_goTypes,
		DependencyIndexes: file_plugin_protocol_proto_depIdxs,
		MessageInfos:      file_plugin_protocol_proto_msgTypes,
	}.Build()
	File_plugin_protocol_proto = out.File
	file_plugin_protocol_proto_rawDesc = nil
	file_plugin_protocol_proto_goTypes = nil
	file_plugin_protocol_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProtocolClient is the client API for Protocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProtocolClient interface {
	Register(ctx context.Context, in *Protocol, opts ...grpc.CallOption) (*Response, error)
	Open(ctx context.Context, in *Open, opts ...grpc.CallOption) (*Response, error)
	Close(ctx context.Context, in *Close, opts ...grpc.CallOption) (*Response, error)
	Read(ctx context.Context, in *Read, opts ...grpc.CallOption) (*Result, error)
	Write(ctx context.Context, in *Write, opts ...grpc.CallOption) (*Response, error)
}

type protocolClient struct {
	cc grpc.ClientConnInterface
}

func NewProtocolClient(cc grpc.ClientConnInterface) ProtocolClient {
	return &protocolClient{cc}
}

func (c *protocolClient) Register(ctx context.Context, in *Protocol, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protocol.protocol/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) Open(ctx context.Context, in *Open, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protocol.protocol/open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) Close(ctx context.Context, in *Close, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protocol.protocol/close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) Read(ctx context.Context, in *Read, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/protocol.protocol/read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) Write(ctx context.Context, in *Write, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protocol.protocol/write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtocolServer is the server API for Protocol service.
type ProtocolServer interface {
	Register(context.Context, *Protocol) (*Response, error)
	Open(context.Context, *Open) (*Response, error)
	Close(context.Context, *Close) (*Response, error)
	Read(context.Context, *Read) (*Result, error)
	Write(context.Context, *Write) (*Response, error)
}

// UnimplementedProtocolServer can be embedded to have forward compatible implementations.
type UnimplementedProtocolServer struct {
}

func (*UnimplementedProtocolServer) Register(context.Context, *Protocol) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedProtocolServer) Open(context.Context, *Open) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (*UnimplementedProtocolServer) Close(context.Context, *Close) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (*UnimplementedProtocolServer) Read(context.Context, *Read) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedProtocolServer) Write(context.Context, *Write) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}

func RegisterProtocolServer(s *grpc.Server, srv ProtocolServer) {
	s.RegisterService(&_Protocol_serviceDesc, srv)
}

func _Protocol_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Protocol)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.protocol/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Register(ctx, req.(*Protocol))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Open)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.protocol/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Open(ctx, req.(*Open))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Close)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.protocol/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Close(ctx, req.(*Close))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Read)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.protocol/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Read(ctx, req.(*Read))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Write)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.protocol/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Write(ctx, req.(*Write))
	}
	return interceptor(ctx, in, info, handler)
}

var _Protocol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.protocol",
	HandlerType: (*ProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _Protocol_Register_Handler,
		},
		{
			MethodName: "open",
			Handler:    _Protocol_Open_Handler,
		},
		{
			MethodName: "close",
			Handler:    _Protocol_Close_Handler,
		},
		{
			MethodName: "read",
			Handler:    _Protocol_Read_Handler,
		},
		{
			MethodName: "write",
			Handler:    _Protocol_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin/protocol.proto",
}
