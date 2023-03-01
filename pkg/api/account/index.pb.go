// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: index.proto

package account

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

var File_index_proto protoreflect.FileDescriptor

var file_index_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xf5, 0x02, 0x0a, 0x1b, 0x4f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x41, 0x70, 0x69,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x47, 0x52, 0x50, 0x43, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x51, 0x51, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x51,
	0x51, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x68, 0x6f, 0x77,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x47, 0x52, 0x50, 0x43, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c,
	0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x14, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x68, 0x6f, 0x77, 0x47, 0x52,
	0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x47, 0x52, 0x50, 0x43, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x7a, 0x68, 0x69, 0x73, 0x68,
	0x65, 0x6e, 0x2f, 0x6f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_index_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),                 // 0: google.protobuf.Empty
	(*model.QQProfileGRPC)(nil),           // 1: model.QQProfileGRPC
	(*model.ModelShowGRPC)(nil),           // 2: model.ModelShowGRPC
	(*model.AccountResultGRPC)(nil),       // 3: model.AccountResultGRPC
	(*model.ModelShowResultGRPC)(nil),     // 4: model.ModelShowResultGRPC
	(*model.OnlineClientsResultGRPC)(nil), // 5: model.OnlineClientsResultGRPC
}
var file_index_proto_depIdxs = []int32{
	0, // 0: account.OnebotApiAccountGRPCService.GetLoginInfo:input_type -> google.protobuf.Empty
	1, // 1: account.OnebotApiAccountGRPCService.SetQQProfile:input_type -> model.QQProfileGRPC
	0, // 2: account.OnebotApiAccountGRPCService.GetModelShow:input_type -> google.protobuf.Empty
	2, // 3: account.OnebotApiAccountGRPCService.SetModelShow:input_type -> model.ModelShowGRPC
	0, // 4: account.OnebotApiAccountGRPCService.GetOnlineClients:input_type -> google.protobuf.Empty
	3, // 5: account.OnebotApiAccountGRPCService.GetLoginInfo:output_type -> model.AccountResultGRPC
	0, // 6: account.OnebotApiAccountGRPCService.SetQQProfile:output_type -> google.protobuf.Empty
	4, // 7: account.OnebotApiAccountGRPCService.GetModelShow:output_type -> model.ModelShowResultGRPC
	0, // 8: account.OnebotApiAccountGRPCService.SetModelShow:output_type -> google.protobuf.Empty
	5, // 9: account.OnebotApiAccountGRPCService.GetOnlineClients:output_type -> model.OnlineClientsResultGRPC
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_index_proto_init() }
func file_index_proto_init() {
	if File_index_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_index_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_index_proto_goTypes,
		DependencyIndexes: file_index_proto_depIdxs,
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

// OnebotApiAccountGRPCServiceClient is the client API for OnebotApiAccountGRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnebotApiAccountGRPCServiceClient interface {
	GetLoginInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*model.AccountResultGRPC, error)
	SetQQProfile(ctx context.Context, in *model.QQProfileGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetModelShow(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*model.ModelShowResultGRPC, error)
	SetModelShow(ctx context.Context, in *model.ModelShowGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetOnlineClients(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*model.OnlineClientsResultGRPC, error)
}

type onebotApiAccountGRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOnebotApiAccountGRPCServiceClient(cc grpc.ClientConnInterface) OnebotApiAccountGRPCServiceClient {
	return &onebotApiAccountGRPCServiceClient{cc}
}

func (c *onebotApiAccountGRPCServiceClient) GetLoginInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*model.AccountResultGRPC, error) {
	out := new(model.AccountResultGRPC)
	err := c.cc.Invoke(ctx, "/account.OnebotApiAccountGRPCService/GetLoginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotApiAccountGRPCServiceClient) SetQQProfile(ctx context.Context, in *model.QQProfileGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/account.OnebotApiAccountGRPCService/SetQQProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotApiAccountGRPCServiceClient) GetModelShow(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*model.ModelShowResultGRPC, error) {
	out := new(model.ModelShowResultGRPC)
	err := c.cc.Invoke(ctx, "/account.OnebotApiAccountGRPCService/GetModelShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotApiAccountGRPCServiceClient) SetModelShow(ctx context.Context, in *model.ModelShowGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/account.OnebotApiAccountGRPCService/SetModelShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotApiAccountGRPCServiceClient) GetOnlineClients(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*model.OnlineClientsResultGRPC, error) {
	out := new(model.OnlineClientsResultGRPC)
	err := c.cc.Invoke(ctx, "/account.OnebotApiAccountGRPCService/GetOnlineClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnebotApiAccountGRPCServiceServer is the server API for OnebotApiAccountGRPCService service.
type OnebotApiAccountGRPCServiceServer interface {
	GetLoginInfo(context.Context, *emptypb.Empty) (*model.AccountResultGRPC, error)
	SetQQProfile(context.Context, *model.QQProfileGRPC) (*emptypb.Empty, error)
	GetModelShow(context.Context, *emptypb.Empty) (*model.ModelShowResultGRPC, error)
	SetModelShow(context.Context, *model.ModelShowGRPC) (*emptypb.Empty, error)
	GetOnlineClients(context.Context, *emptypb.Empty) (*model.OnlineClientsResultGRPC, error)
}

// UnimplementedOnebotApiAccountGRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOnebotApiAccountGRPCServiceServer struct {
}

func (*UnimplementedOnebotApiAccountGRPCServiceServer) GetLoginInfo(context.Context, *emptypb.Empty) (*model.AccountResultGRPC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginInfo not implemented")
}
func (*UnimplementedOnebotApiAccountGRPCServiceServer) SetQQProfile(context.Context, *model.QQProfileGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetQQProfile not implemented")
}
func (*UnimplementedOnebotApiAccountGRPCServiceServer) GetModelShow(context.Context, *emptypb.Empty) (*model.ModelShowResultGRPC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelShow not implemented")
}
func (*UnimplementedOnebotApiAccountGRPCServiceServer) SetModelShow(context.Context, *model.ModelShowGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetModelShow not implemented")
}
func (*UnimplementedOnebotApiAccountGRPCServiceServer) GetOnlineClients(context.Context, *emptypb.Empty) (*model.OnlineClientsResultGRPC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnlineClients not implemented")
}

func RegisterOnebotApiAccountGRPCServiceServer(s *grpc.Server, srv OnebotApiAccountGRPCServiceServer) {
	s.RegisterService(&_OnebotApiAccountGRPCService_serviceDesc, srv)
}

func _OnebotApiAccountGRPCService_GetLoginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotApiAccountGRPCServiceServer).GetLoginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.OnebotApiAccountGRPCService/GetLoginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotApiAccountGRPCServiceServer).GetLoginInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotApiAccountGRPCService_SetQQProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.QQProfileGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotApiAccountGRPCServiceServer).SetQQProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.OnebotApiAccountGRPCService/SetQQProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotApiAccountGRPCServiceServer).SetQQProfile(ctx, req.(*model.QQProfileGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotApiAccountGRPCService_GetModelShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotApiAccountGRPCServiceServer).GetModelShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.OnebotApiAccountGRPCService/GetModelShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotApiAccountGRPCServiceServer).GetModelShow(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotApiAccountGRPCService_SetModelShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ModelShowGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotApiAccountGRPCServiceServer).SetModelShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.OnebotApiAccountGRPCService/SetModelShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotApiAccountGRPCServiceServer).SetModelShow(ctx, req.(*model.ModelShowGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotApiAccountGRPCService_GetOnlineClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotApiAccountGRPCServiceServer).GetOnlineClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.OnebotApiAccountGRPCService/GetOnlineClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotApiAccountGRPCServiceServer).GetOnlineClients(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnebotApiAccountGRPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.OnebotApiAccountGRPCService",
	HandlerType: (*OnebotApiAccountGRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLoginInfo",
			Handler:    _OnebotApiAccountGRPCService_GetLoginInfo_Handler,
		},
		{
			MethodName: "SetQQProfile",
			Handler:    _OnebotApiAccountGRPCService_SetQQProfile_Handler,
		},
		{
			MethodName: "GetModelShow",
			Handler:    _OnebotApiAccountGRPCService_GetModelShow_Handler,
		},
		{
			MethodName: "SetModelShow",
			Handler:    _OnebotApiAccountGRPCService_SetModelShow_Handler,
		},
		{
			MethodName: "GetOnlineClients",
			Handler:    _OnebotApiAccountGRPCService_GetOnlineClients_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index.proto",
}
