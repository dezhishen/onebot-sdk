// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: event_meta.proto

package meta

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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_event_meta_proto protoreflect.FileDescriptor

var file_event_meta_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x16, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb6, 0x01,
	0x0a, 0x1a, 0x4f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x10,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x12, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x52, 0x50, 0x43, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x10, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x1d, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x7a, 0x68, 0x69, 0x73, 0x68, 0x65, 0x6e, 0x2f, 0x6f,
	0x6e, 0x65, 0x62, 0x6f, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x3b, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_event_meta_proto_goTypes = []interface{}{
	(*model.EventMetaLifecycleGRPC)(nil), // 0: model.EventMetaLifecycleGRPC
	(*model.EventMetaHeartbeatGRPC)(nil), // 1: model.EventMetaHeartbeatGRPC
	(*emptypb.Empty)(nil),                // 2: google.protobuf.Empty
}
var file_event_meta_proto_depIdxs = []int32{
	0, // 0: meta.OnebotEventMetaGRPCService.HandlerLifecycle:input_type -> model.EventMetaLifecycleGRPC
	1, // 1: meta.OnebotEventMetaGRPCService.HandlerHeartBeat:input_type -> model.EventMetaHeartbeatGRPC
	2, // 2: meta.OnebotEventMetaGRPCService.HandlerLifecycle:output_type -> google.protobuf.Empty
	2, // 3: meta.OnebotEventMetaGRPCService.HandlerHeartBeat:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_event_meta_proto_init() }
func file_event_meta_proto_init() {
	if File_event_meta_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_meta_proto_goTypes,
		DependencyIndexes: file_event_meta_proto_depIdxs,
	}.Build()
	File_event_meta_proto = out.File
	file_event_meta_proto_rawDesc = nil
	file_event_meta_proto_goTypes = nil
	file_event_meta_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OnebotEventMetaGRPCServiceClient is the client API for OnebotEventMetaGRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnebotEventMetaGRPCServiceClient interface {
	// ListenMetaLifecycle(handler func(data model.EventMetaLifecycle) error)
	HandlerLifecycle(ctx context.Context, in *model.EventMetaLifecycleGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListenMetaHeartBeat(handler func(data model.EventMetaHeartbeat) error)
	HandlerHeartBeat(ctx context.Context, in *model.EventMetaHeartbeatGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type onebotEventMetaGRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOnebotEventMetaGRPCServiceClient(cc grpc.ClientConnInterface) OnebotEventMetaGRPCServiceClient {
	return &onebotEventMetaGRPCServiceClient{cc}
}

func (c *onebotEventMetaGRPCServiceClient) HandlerLifecycle(ctx context.Context, in *model.EventMetaLifecycleGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/meta.OnebotEventMetaGRPCService/HandlerLifecycle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotEventMetaGRPCServiceClient) HandlerHeartBeat(ctx context.Context, in *model.EventMetaHeartbeatGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/meta.OnebotEventMetaGRPCService/HandlerHeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnebotEventMetaGRPCServiceServer is the server API for OnebotEventMetaGRPCService service.
type OnebotEventMetaGRPCServiceServer interface {
	// ListenMetaLifecycle(handler func(data model.EventMetaLifecycle) error)
	HandlerLifecycle(context.Context, *model.EventMetaLifecycleGRPC) (*emptypb.Empty, error)
	// ListenMetaHeartBeat(handler func(data model.EventMetaHeartbeat) error)
	HandlerHeartBeat(context.Context, *model.EventMetaHeartbeatGRPC) (*emptypb.Empty, error)
}

// UnimplementedOnebotEventMetaGRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOnebotEventMetaGRPCServiceServer struct {
}

func (*UnimplementedOnebotEventMetaGRPCServiceServer) HandlerLifecycle(context.Context, *model.EventMetaLifecycleGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandlerLifecycle not implemented")
}
func (*UnimplementedOnebotEventMetaGRPCServiceServer) HandlerHeartBeat(context.Context, *model.EventMetaHeartbeatGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandlerHeartBeat not implemented")
}

func RegisterOnebotEventMetaGRPCServiceServer(s *grpc.Server, srv OnebotEventMetaGRPCServiceServer) {
	s.RegisterService(&_OnebotEventMetaGRPCService_serviceDesc, srv)
}

func _OnebotEventMetaGRPCService_HandlerLifecycle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventMetaLifecycleGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventMetaGRPCServiceServer).HandlerLifecycle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta.OnebotEventMetaGRPCService/HandlerLifecycle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventMetaGRPCServiceServer).HandlerLifecycle(ctx, req.(*model.EventMetaLifecycleGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotEventMetaGRPCService_HandlerHeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventMetaHeartbeatGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventMetaGRPCServiceServer).HandlerHeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta.OnebotEventMetaGRPCService/HandlerHeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventMetaGRPCServiceServer).HandlerHeartBeat(ctx, req.(*model.EventMetaHeartbeatGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnebotEventMetaGRPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "meta.OnebotEventMetaGRPCService",
	HandlerType: (*OnebotEventMetaGRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandlerLifecycle",
			Handler:    _OnebotEventMetaGRPCService_HandlerLifecycle_Handler,
		},
		{
			MethodName: "HandlerHeartBeat",
			Handler:    _OnebotEventMetaGRPCService_HandlerHeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event_meta.proto",
}
