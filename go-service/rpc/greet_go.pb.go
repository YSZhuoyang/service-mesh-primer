// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: greet_go.proto

package rpc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_go_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_go_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_greet_go_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_go_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_greet_go_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_greet_go_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type LiveDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LiveDataRequest) Reset() {
	*x = LiveDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_go_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveDataRequest) ProtoMessage() {}

func (x *LiveDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_go_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveDataRequest.ProtoReflect.Descriptor instead.
func (*LiveDataRequest) Descriptor() ([]byte, []int) {
	return file_greet_go_proto_rawDescGZIP(), []int{2}
}

type LiveDataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data int32 `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LiveDataReply) Reset() {
	*x = LiveDataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_go_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveDataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveDataReply) ProtoMessage() {}

func (x *LiveDataReply) ProtoReflect() protoreflect.Message {
	mi := &file_greet_go_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveDataReply.ProtoReflect.Descriptor instead.
func (*LiveDataReply) Descriptor() ([]byte, []int) {
	return file_greet_go_proto_rawDescGZIP(), []int{3}
}

func (x *LiveDataReply) GetData() int32 {
	if x != nil {
		return x.Data
	}
	return 0
}

var File_greet_go_proto protoreflect.FileDescriptor

var file_greet_go_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0c, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x1e, 0x0a, 0x0a,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x11, 0x0a, 0x0f,
	0x4c, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x23, 0x0a, 0x0d, 0x4c, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0xab, 0x01, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x12, 0x55, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e, 0x67,
	0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0f, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2d, 0x67,
	0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x49, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4c, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x67, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greet_go_proto_rawDescOnce sync.Once
	file_greet_go_proto_rawDescData = file_greet_go_proto_rawDesc
)

func file_greet_go_proto_rawDescGZIP() []byte {
	file_greet_go_proto_rawDescOnce.Do(func() {
		file_greet_go_proto_rawDescData = protoimpl.X.CompressGZIP(file_greet_go_proto_rawDescData)
	})
	return file_greet_go_proto_rawDescData
}

var file_greet_go_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_greet_go_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),    // 0: go_service.HelloRequest
	(*HelloReply)(nil),      // 1: go_service.HelloReply
	(*LiveDataRequest)(nil), // 2: go_service.LiveDataRequest
	(*LiveDataReply)(nil),   // 3: go_service.LiveDataReply
}
var file_greet_go_proto_depIdxs = []int32{
	0, // 0: go_service.Greeter.SayHello:input_type -> go_service.HelloRequest
	2, // 1: go_service.Greeter.GetLiveData:input_type -> go_service.LiveDataRequest
	1, // 2: go_service.Greeter.SayHello:output_type -> go_service.HelloReply
	3, // 3: go_service.Greeter.GetLiveData:output_type -> go_service.LiveDataReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greet_go_proto_init() }
func file_greet_go_proto_init() {
	if File_greet_go_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greet_go_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_greet_go_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_greet_go_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveDataRequest); i {
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
		file_greet_go_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveDataReply); i {
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
			RawDescriptor: file_greet_go_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greet_go_proto_goTypes,
		DependencyIndexes: file_greet_go_proto_depIdxs,
		MessageInfos:      file_greet_go_proto_msgTypes,
	}.Build()
	File_greet_go_proto = out.File
	file_greet_go_proto_rawDesc = nil
	file_greet_go_proto_goTypes = nil
	file_greet_go_proto_depIdxs = nil
}
