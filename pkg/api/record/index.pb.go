// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: index.proto

package record

import (
	context "context"
	model "github.com/dezhishen/onebot-sdk/pkg/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File      string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	OutFormat string `protobuf:"bytes,2,opt,name=out_format,json=outFormat,proto3" json:"out_format,omitempty"`
}

func (x *GetRecordRequest) Reset() {
	*x = GetRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordRequest) ProtoMessage() {}

func (x *GetRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordRequest.ProtoReflect.Descriptor instead.
func (*GetRecordRequest) Descriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{0}
}

func (x *GetRecordRequest) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *GetRecordRequest) GetOutFormat() string {
	if x != nil {
		return x.OutFormat
	}
	return ""
}

var File_index_proto protoreflect.FileDescriptor

var file_index_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x75, 0x74, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x32, 0xa5, 0x01, 0x0a, 0x1a, 0x4f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x18, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x47, 0x52,
	0x50, 0x43, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x73, 0x4f, 0x66, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x47, 0x52, 0x50, 0x43, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x7a, 0x68, 0x69, 0x73,
	0x68, 0x65, 0x6e, 0x2f, 0x6f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x3b, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_index_proto_rawDescOnce sync.Once
	file_index_proto_rawDescData = file_index_proto_rawDesc
)

func file_index_proto_rawDescGZIP() []byte {
	file_index_proto_rawDescOnce.Do(func() {
		file_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_index_proto_rawDescData)
	})
	return file_index_proto_rawDescData
}

var file_index_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_index_proto_goTypes = []interface{}{
	(*GetRecordRequest)(nil),          // 0: record.GetRecordRequest
	(*emptypb.Empty)(nil),             // 1: google.protobuf.Empty
	(*model.RecordResultGRPC)(nil),    // 2: model.RecordResultGRPC
	(*model.BoolYesOfResultGRPC)(nil), // 3: model.BoolYesOfResultGRPC
}
var file_index_proto_depIdxs = []int32{
	0, // 0: record.OnebotApiRecordGRPCService.GetRecord:input_type -> record.GetRecordRequest
	1, // 1: record.OnebotApiRecordGRPCService.CanSendRecord:input_type -> google.protobuf.Empty
	2, // 2: record.OnebotApiRecordGRPCService.GetRecord:output_type -> model.RecordResultGRPC
	3, // 3: record.OnebotApiRecordGRPCService.CanSendRecord:output_type -> model.BoolYesOfResultGRPC
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_index_proto_init() }
func file_index_proto_init() {
	if File_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordRequest); i {
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
			RawDescriptor: file_index_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_index_proto_goTypes,
		DependencyIndexes: file_index_proto_depIdxs,
		MessageInfos:      file_index_proto_msgTypes,
	}.Build()
	File_index_proto = out.File
	file_index_proto_rawDesc = nil
	file_index_proto_goTypes = nil
	file_index_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OnebotApiRecordGRPCServiceClient is the client API for OnebotApiRecordGRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnebotApiRecordGRPCServiceClient interface {
	// 获取语音
	// get_record
	GetRecord(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*model.RecordResultGRPC, error)
	// 检查是否可以发送语音
	// can_send_record
	CanSendRecord(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*model.BoolYesOfResultGRPC, error)
}

type onebotApiRecordGRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOnebotApiRecordGRPCServiceClient(cc grpc.ClientConnInterface) OnebotApiRecordGRPCServiceClient {
	return &onebotApiRecordGRPCServiceClient{cc}
}

func (c *onebotApiRecordGRPCServiceClient) GetRecord(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*model.RecordResultGRPC, error) {
	out := new(model.RecordResultGRPC)
	err := c.cc.Invoke(ctx, "/record.OnebotApiRecordGRPCService/GetRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotApiRecordGRPCServiceClient) CanSendRecord(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*model.BoolYesOfResultGRPC, error) {
	out := new(model.BoolYesOfResultGRPC)
	err := c.cc.Invoke(ctx, "/record.OnebotApiRecordGRPCService/CanSendRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnebotApiRecordGRPCServiceServer is the server API for OnebotApiRecordGRPCService service.
type OnebotApiRecordGRPCServiceServer interface {
	// 获取语音
	// get_record
	GetRecord(context.Context, *GetRecordRequest) (*model.RecordResultGRPC, error)
	// 检查是否可以发送语音
	// can_send_record
	CanSendRecord(context.Context, *emptypb.Empty) (*model.BoolYesOfResultGRPC, error)
}

// UnimplementedOnebotApiRecordGRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOnebotApiRecordGRPCServiceServer struct {
}

func (*UnimplementedOnebotApiRecordGRPCServiceServer) GetRecord(context.Context, *GetRecordRequest) (*model.RecordResultGRPC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecord not implemented")
}
func (*UnimplementedOnebotApiRecordGRPCServiceServer) CanSendRecord(context.Context, *emptypb.Empty) (*model.BoolYesOfResultGRPC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanSendRecord not implemented")
}

func RegisterOnebotApiRecordGRPCServiceServer(s *grpc.Server, srv OnebotApiRecordGRPCServiceServer) {
	s.RegisterService(&_OnebotApiRecordGRPCService_serviceDesc, srv)
}

func _OnebotApiRecordGRPCService_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotApiRecordGRPCServiceServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/record.OnebotApiRecordGRPCService/GetRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotApiRecordGRPCServiceServer).GetRecord(ctx, req.(*GetRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotApiRecordGRPCService_CanSendRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotApiRecordGRPCServiceServer).CanSendRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/record.OnebotApiRecordGRPCService/CanSendRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotApiRecordGRPCServiceServer).CanSendRecord(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnebotApiRecordGRPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "record.OnebotApiRecordGRPCService",
	HandlerType: (*OnebotApiRecordGRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecord",
			Handler:    _OnebotApiRecordGRPCService_GetRecord_Handler,
		},
		{
			MethodName: "CanSendRecord",
			Handler:    _OnebotApiRecordGRPCService_CanSendRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index.proto",
}
