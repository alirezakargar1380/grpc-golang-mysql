// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: user_log.proto

package user_log

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res *Log_Res `protobuf:"bytes,2,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_user_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_user_log_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetRes() *Log_Res {
	if x != nil {
		return x.Res
	}
	return nil
}

type Log_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Brand string `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *Log_Res) Reset() {
	*x = Log_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log_Res) ProtoMessage() {}

func (x *Log_Res) ProtoReflect() protoreflect.Message {
	mi := &file_user_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log_Res.ProtoReflect.Descriptor instead.
func (*Log_Res) Descriptor() ([]byte, []int) {
	return file_user_log_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Log_Res) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Log_Res) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

var File_user_log_proto protoreflect.FileDescriptor

var file_user_log_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x6c, 0x6f, 0x67, 0x22, 0x56, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x1e, 0x0a, 0x03,
	0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x67, 0x2e,
	0x4c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x52, 0x03, 0x72, 0x65, 0x73, 0x1a, 0x2f, 0x0a, 0x03,
	0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x32, 0x31, 0x0a,
	0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0a,
	0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x08, 0x2e, 0x6c, 0x6f, 0x67,
	0x2e, 0x4c, 0x6f, 0x67, 0x1a, 0x08, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x22, 0x00,
	0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_log_proto_rawDescOnce sync.Once
	file_user_log_proto_rawDescData = file_user_log_proto_rawDesc
)

func file_user_log_proto_rawDescGZIP() []byte {
	file_user_log_proto_rawDescOnce.Do(func() {
		file_user_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_log_proto_rawDescData)
	})
	return file_user_log_proto_rawDescData
}

var file_user_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_log_proto_goTypes = []interface{}{
	(*Log)(nil),     // 0: log.Log
	(*Log_Res)(nil), // 1: log.Log.Res
}
var file_user_log_proto_depIdxs = []int32{
	1, // 0: log.Log.res:type_name -> log.Log.Res
	0, // 1: log.ChatService.AddMessage:input_type -> log.Log
	0, // 2: log.ChatService.AddMessage:output_type -> log.Log
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_log_proto_init() }
func file_user_log_proto_init() {
	if File_user_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_user_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log_Res); i {
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
			RawDescriptor: file_user_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_log_proto_goTypes,
		DependencyIndexes: file_user_log_proto_depIdxs,
		MessageInfos:      file_user_log_proto_msgTypes,
	}.Build()
	File_user_log_proto = out.File
	file_user_log_proto_rawDesc = nil
	file_user_log_proto_goTypes = nil
	file_user_log_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	AddMessage(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Log, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) AddMessage(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Log, error) {
	out := new(Log)
	err := c.cc.Invoke(ctx, "/log.ChatService/AddMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	AddMessage(context.Context, *Log) (*Log, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) AddMessage(context.Context, *Log) (*Log, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMessage not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_AddMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/log.ChatService/AddMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddMessage(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "log.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMessage",
			Handler:    _ChatService_AddMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_log.proto",
}
