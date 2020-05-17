// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: test.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type In struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *In) Reset() {
	*x = In{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *In) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*In) ProtoMessage() {}

func (x *In) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use In.ProtoReflect.Descriptor instead.
func (*In) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *In) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Out struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Out) Reset() {
	*x = Out{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Out) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Out) ProtoMessage() {}

func (x *Out) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Out.ProtoReflect.Descriptor instead.
func (*Out) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *Out) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x02,
	0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x19, 0x0a, 0x03, 0x4f, 0x75, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x32, 0x21, 0x0a, 0x08, 0x50, 0x75, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x15, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x03, 0x2e, 0x49, 0x6e, 0x1a, 0x04, 0x2e, 0x4f,
	0x75, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_goTypes = []interface{}{
	(*In)(nil),  // 0: In
	(*Out)(nil), // 1: Out
}
var file_test_proto_depIdxs = []int32{
	0, // 0: PushData.Call:input_type -> In
	1, // 1: PushData.Call:output_type -> Out
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*In); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Out); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PushDataClient is the client API for PushData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PushDataClient interface {
	Call(ctx context.Context, in *In, opts ...grpc.CallOption) (PushData_CallClient, error)
}

type pushDataClient struct {
	cc grpc.ClientConnInterface
}

func NewPushDataClient(cc grpc.ClientConnInterface) PushDataClient {
	return &pushDataClient{cc}
}

func (c *pushDataClient) Call(ctx context.Context, in *In, opts ...grpc.CallOption) (PushData_CallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PushData_serviceDesc.Streams[0], "/PushData/Call", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushDataCallClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PushData_CallClient interface {
	Recv() (*Out, error)
	grpc.ClientStream
}

type pushDataCallClient struct {
	grpc.ClientStream
}

func (x *pushDataCallClient) Recv() (*Out, error) {
	m := new(Out)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PushDataServer is the server API for PushData service.
type PushDataServer interface {
	Call(*In, PushData_CallServer) error
}

// UnimplementedPushDataServer can be embedded to have forward compatible implementations.
type UnimplementedPushDataServer struct {
}

func (*UnimplementedPushDataServer) Call(*In, PushData_CallServer) error {
	return status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterPushDataServer(s *grpc.Server, srv PushDataServer) {
	s.RegisterService(&_PushData_serviceDesc, srv)
}

func _PushData_Call_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(In)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushDataServer).Call(m, &pushDataCallServer{stream})
}

type PushData_CallServer interface {
	Send(*Out) error
	grpc.ServerStream
}

type pushDataCallServer struct {
	grpc.ServerStream
}

func (x *pushDataCallServer) Send(m *Out) error {
	return x.ServerStream.SendMsg(m)
}

var _PushData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PushData",
	HandlerType: (*PushDataServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Call",
			Handler:       _PushData_Call_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "test.proto",
}
