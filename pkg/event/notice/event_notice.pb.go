// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: event_notice.proto

package notice

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

var File_event_notice_proto protoreflect.FileDescriptor

var file_event_notice_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x1a, 0x18, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x88, 0x08, 0x0a, 0x1c, 0x4f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x17, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x21, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x47, 0x52,
	0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x16,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x20, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x5a, 0x0a, 0x19, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x12,
	0x23, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65,
	0x47, 0x52, 0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5a,
	0x0a, 0x19, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x12, 0x23, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x47, 0x52, 0x50, 0x43,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x14, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42,
	0x61, 0x6e, 0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x61, 0x6e, 0x47, 0x52,
	0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x15,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x41, 0x64, 0x64, 0x12, 0x1f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x41,
	0x64, 0x64, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x56, 0x0a, 0x17, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x21, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x18, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x63, 0x61, 0x6c, 0x6c, 0x12, 0x22, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x63, 0x61, 0x6c, 0x6c, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x5e, 0x0a, 0x1b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x6f, 0x6b,
	0x65, 0x12, 0x25, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x50, 0x6f, 0x6b, 0x65, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x60, 0x0a, 0x1c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x48, 0x6f, 0x6e,
	0x6f, 0x72, 0x12, 0x26, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x48, 0x6f, 0x6e, 0x6f, 0x72, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x20, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4c,
	0x75, 0x63, 0x6b, 0x79, 0x4b, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4c, 0x75, 0x63, 0x6b, 0x79, 0x4b, 0x69, 0x6e, 0x67,
	0x47, 0x52, 0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x39,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x7a,
	0x68, 0x69, 0x73, 0x68, 0x65, 0x6e, 0x2f, 0x6f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_event_notice_proto_goTypes = []interface{}{
	(*model.EventNoticeGroupUploadGRPC)(nil),          // 0: model.EventNoticeGroupUploadGRPC
	(*model.EventNoticeGroupAdminGRPC)(nil),           // 1: model.EventNoticeGroupAdminGRPC
	(*model.EventNoticeGroupDecreaseGRPC)(nil),        // 2: model.EventNoticeGroupDecreaseGRPC
	(*model.EventNoticeGroupIncreaseGRPC)(nil),        // 3: model.EventNoticeGroupIncreaseGRPC
	(*model.EventNoticeGroupBanGRPC)(nil),             // 4: model.EventNoticeGroupBanGRPC
	(*model.EventNoticeFriendAddGRPC)(nil),            // 5: model.EventNoticeFriendAddGRPC
	(*model.EventNoticeGroupRecallGRPC)(nil),          // 6: model.EventNoticeGroupRecallGRPC
	(*model.EventNoticeFriendRecallGRPC)(nil),         // 7: model.EventNoticeFriendRecallGRPC
	(*model.EventNoticeGroupNotifyPokeGRPC)(nil),      // 8: model.EventNoticeGroupNotifyPokeGRPC
	(*model.EventNoticeGroupNotifyHonorGRPC)(nil),     // 9: model.EventNoticeGroupNotifyHonorGRPC
	(*model.EventNoticeGroupNotifyLuckyKingGRPC)(nil), // 10: model.EventNoticeGroupNotifyLuckyKingGRPC
	(*emptypb.Empty)(nil),                             // 11: google.protobuf.Empty
}
var file_event_notice_proto_depIdxs = []int32{
	0,  // 0: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupUpload:input_type -> model.EventNoticeGroupUploadGRPC
	1,  // 1: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupAdmin:input_type -> model.EventNoticeGroupAdminGRPC
	2,  // 2: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupDecrease:input_type -> model.EventNoticeGroupDecreaseGRPC
	3,  // 3: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupIncrease:input_type -> model.EventNoticeGroupIncreaseGRPC
	4,  // 4: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupBan:input_type -> model.EventNoticeGroupBanGRPC
	5,  // 5: notice.OnebotEventNoticeGRPCService.HandleNoticeFriendAdd:input_type -> model.EventNoticeFriendAddGRPC
	6,  // 6: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupRecall:input_type -> model.EventNoticeGroupRecallGRPC
	7,  // 7: notice.OnebotEventNoticeGRPCService.HandleNoticeFriendRecall:input_type -> model.EventNoticeFriendRecallGRPC
	8,  // 8: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupNotifyPoke:input_type -> model.EventNoticeGroupNotifyPokeGRPC
	9,  // 9: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupNotifyHonor:input_type -> model.EventNoticeGroupNotifyHonorGRPC
	10, // 10: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupNotifyLuckyKing:input_type -> model.EventNoticeGroupNotifyLuckyKingGRPC
	11, // 11: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupUpload:output_type -> google.protobuf.Empty
	11, // 12: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupAdmin:output_type -> google.protobuf.Empty
	11, // 13: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupDecrease:output_type -> google.protobuf.Empty
	11, // 14: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupIncrease:output_type -> google.protobuf.Empty
	11, // 15: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupBan:output_type -> google.protobuf.Empty
	11, // 16: notice.OnebotEventNoticeGRPCService.HandleNoticeFriendAdd:output_type -> google.protobuf.Empty
	11, // 17: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupRecall:output_type -> google.protobuf.Empty
	11, // 18: notice.OnebotEventNoticeGRPCService.HandleNoticeFriendRecall:output_type -> google.protobuf.Empty
	11, // 19: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupNotifyPoke:output_type -> google.protobuf.Empty
	11, // 20: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupNotifyHonor:output_type -> google.protobuf.Empty
	11, // 21: notice.OnebotEventNoticeGRPCService.HandleNoticeGroupNotifyLuckyKing:output_type -> google.protobuf.Empty
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_event_notice_proto_init() }
func file_event_notice_proto_init() {
	if File_event_notice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_notice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_notice_proto_goTypes,
		DependencyIndexes: file_event_notice_proto_depIdxs,
	}.Build()
	File_event_notice_proto = out.File
	file_event_notice_proto_rawDesc = nil
	file_event_notice_proto_goTypes = nil
	file_event_notice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OnebotEventNoticeGRPCServiceClient is the client API for OnebotEventNoticeGRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnebotEventNoticeGRPCServiceClient interface {
	// ListenNoticeGroupUpload(handler func(data model.EventNoticeGroupUpload) error)
	HandleNoticeGroupUpload(ctx context.Context, in *model.EventNoticeGroupUploadGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListenNoticeGroupAdmin(handler func(data model.EventNoticeGroupAdmin) error)
	HandleNoticeGroupAdmin(ctx context.Context, in *model.EventNoticeGroupAdminGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListenNoticeGroupDecrease(handler func(data model.EventNoticeGroupDecrease) error)
	HandleNoticeGroupDecrease(ctx context.Context, in *model.EventNoticeGroupDecreaseGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListenNoticeGroupIncrease(handler func(data model.EventNoticeGroupIncrease) error)
	HandleNoticeGroupIncrease(ctx context.Context, in *model.EventNoticeGroupIncreaseGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListenNoticeGroupBan(handler func(data model.EventNoticeGroupBan) error)
	HandleNoticeGroupBan(ctx context.Context, in *model.EventNoticeGroupBanGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListenNoticeFriendAdd(handler func(data model.EventNoticeFriendAdd) error)
	HandleNoticeFriendAdd(ctx context.Context, in *model.EventNoticeFriendAddGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListenNoticeGroupRecall(handler func(data model.EventNoticeGroupRecall) error)
	HandleNoticeGroupRecall(ctx context.Context, in *model.EventNoticeGroupRecallGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListenNoticeFriendRecall(handler func(data model.EventNoticeFriendRecall) error)
	HandleNoticeFriendRecall(ctx context.Context, in *model.EventNoticeFriendRecallGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListenNoticeGroupNotifyPoke(handler func(data model.EventNoticeGroupNotifyPoke) error)
	HandleNoticeGroupNotifyPoke(ctx context.Context, in *model.EventNoticeGroupNotifyPokeGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListenNoticeGroupNotifyHonor(handler func(data model.EventNoticeGroupNotifyHonor) error)
	HandleNoticeGroupNotifyHonor(ctx context.Context, in *model.EventNoticeGroupNotifyHonorGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListenNoticeGroupNotifyLuckyKing(handler func(data model.EventNoticeGroupNotifyLuckyKing) error)
	HandleNoticeGroupNotifyLuckyKing(ctx context.Context, in *model.EventNoticeGroupNotifyLuckyKingGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type onebotEventNoticeGRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOnebotEventNoticeGRPCServiceClient(cc grpc.ClientConnInterface) OnebotEventNoticeGRPCServiceClient {
	return &onebotEventNoticeGRPCServiceClient{cc}
}

func (c *onebotEventNoticeGRPCServiceClient) HandleNoticeGroupUpload(ctx context.Context, in *model.EventNoticeGroupUploadGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotEventNoticeGRPCServiceClient) HandleNoticeGroupAdmin(ctx context.Context, in *model.EventNoticeGroupAdminGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotEventNoticeGRPCServiceClient) HandleNoticeGroupDecrease(ctx context.Context, in *model.EventNoticeGroupDecreaseGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupDecrease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotEventNoticeGRPCServiceClient) HandleNoticeGroupIncrease(ctx context.Context, in *model.EventNoticeGroupIncreaseGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupIncrease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotEventNoticeGRPCServiceClient) HandleNoticeGroupBan(ctx context.Context, in *model.EventNoticeGroupBanGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupBan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotEventNoticeGRPCServiceClient) HandleNoticeFriendAdd(ctx context.Context, in *model.EventNoticeFriendAddGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/notice.OnebotEventNoticeGRPCService/HandleNoticeFriendAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotEventNoticeGRPCServiceClient) HandleNoticeGroupRecall(ctx context.Context, in *model.EventNoticeGroupRecallGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupRecall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotEventNoticeGRPCServiceClient) HandleNoticeFriendRecall(ctx context.Context, in *model.EventNoticeFriendRecallGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/notice.OnebotEventNoticeGRPCService/HandleNoticeFriendRecall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotEventNoticeGRPCServiceClient) HandleNoticeGroupNotifyPoke(ctx context.Context, in *model.EventNoticeGroupNotifyPokeGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupNotifyPoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotEventNoticeGRPCServiceClient) HandleNoticeGroupNotifyHonor(ctx context.Context, in *model.EventNoticeGroupNotifyHonorGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupNotifyHonor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotEventNoticeGRPCServiceClient) HandleNoticeGroupNotifyLuckyKing(ctx context.Context, in *model.EventNoticeGroupNotifyLuckyKingGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupNotifyLuckyKing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnebotEventNoticeGRPCServiceServer is the server API for OnebotEventNoticeGRPCService service.
type OnebotEventNoticeGRPCServiceServer interface {
	// ListenNoticeGroupUpload(handler func(data model.EventNoticeGroupUpload) error)
	HandleNoticeGroupUpload(context.Context, *model.EventNoticeGroupUploadGRPC) (*emptypb.Empty, error)
	// ListenNoticeGroupAdmin(handler func(data model.EventNoticeGroupAdmin) error)
	HandleNoticeGroupAdmin(context.Context, *model.EventNoticeGroupAdminGRPC) (*emptypb.Empty, error)
	// ListenNoticeGroupDecrease(handler func(data model.EventNoticeGroupDecrease) error)
	HandleNoticeGroupDecrease(context.Context, *model.EventNoticeGroupDecreaseGRPC) (*emptypb.Empty, error)
	// ListenNoticeGroupIncrease(handler func(data model.EventNoticeGroupIncrease) error)
	HandleNoticeGroupIncrease(context.Context, *model.EventNoticeGroupIncreaseGRPC) (*emptypb.Empty, error)
	// ListenNoticeGroupBan(handler func(data model.EventNoticeGroupBan) error)
	HandleNoticeGroupBan(context.Context, *model.EventNoticeGroupBanGRPC) (*emptypb.Empty, error)
	// ListenNoticeFriendAdd(handler func(data model.EventNoticeFriendAdd) error)
	HandleNoticeFriendAdd(context.Context, *model.EventNoticeFriendAddGRPC) (*emptypb.Empty, error)
	// ListenNoticeGroupRecall(handler func(data model.EventNoticeGroupRecall) error)
	HandleNoticeGroupRecall(context.Context, *model.EventNoticeGroupRecallGRPC) (*emptypb.Empty, error)
	// ListenNoticeFriendRecall(handler func(data model.EventNoticeFriendRecall) error)
	HandleNoticeFriendRecall(context.Context, *model.EventNoticeFriendRecallGRPC) (*emptypb.Empty, error)
	// ListenNoticeGroupNotifyPoke(handler func(data model.EventNoticeGroupNotifyPoke) error)
	HandleNoticeGroupNotifyPoke(context.Context, *model.EventNoticeGroupNotifyPokeGRPC) (*emptypb.Empty, error)
	// ListenNoticeGroupNotifyHonor(handler func(data model.EventNoticeGroupNotifyHonor) error)
	HandleNoticeGroupNotifyHonor(context.Context, *model.EventNoticeGroupNotifyHonorGRPC) (*emptypb.Empty, error)
	// ListenNoticeGroupNotifyLuckyKing(handler func(data model.EventNoticeGroupNotifyLuckyKing) error)
	HandleNoticeGroupNotifyLuckyKing(context.Context, *model.EventNoticeGroupNotifyLuckyKingGRPC) (*emptypb.Empty, error)
}

// UnimplementedOnebotEventNoticeGRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOnebotEventNoticeGRPCServiceServer struct {
}

func (*UnimplementedOnebotEventNoticeGRPCServiceServer) HandleNoticeGroupUpload(context.Context, *model.EventNoticeGroupUploadGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleNoticeGroupUpload not implemented")
}
func (*UnimplementedOnebotEventNoticeGRPCServiceServer) HandleNoticeGroupAdmin(context.Context, *model.EventNoticeGroupAdminGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleNoticeGroupAdmin not implemented")
}
func (*UnimplementedOnebotEventNoticeGRPCServiceServer) HandleNoticeGroupDecrease(context.Context, *model.EventNoticeGroupDecreaseGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleNoticeGroupDecrease not implemented")
}
func (*UnimplementedOnebotEventNoticeGRPCServiceServer) HandleNoticeGroupIncrease(context.Context, *model.EventNoticeGroupIncreaseGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleNoticeGroupIncrease not implemented")
}
func (*UnimplementedOnebotEventNoticeGRPCServiceServer) HandleNoticeGroupBan(context.Context, *model.EventNoticeGroupBanGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleNoticeGroupBan not implemented")
}
func (*UnimplementedOnebotEventNoticeGRPCServiceServer) HandleNoticeFriendAdd(context.Context, *model.EventNoticeFriendAddGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleNoticeFriendAdd not implemented")
}
func (*UnimplementedOnebotEventNoticeGRPCServiceServer) HandleNoticeGroupRecall(context.Context, *model.EventNoticeGroupRecallGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleNoticeGroupRecall not implemented")
}
func (*UnimplementedOnebotEventNoticeGRPCServiceServer) HandleNoticeFriendRecall(context.Context, *model.EventNoticeFriendRecallGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleNoticeFriendRecall not implemented")
}
func (*UnimplementedOnebotEventNoticeGRPCServiceServer) HandleNoticeGroupNotifyPoke(context.Context, *model.EventNoticeGroupNotifyPokeGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleNoticeGroupNotifyPoke not implemented")
}
func (*UnimplementedOnebotEventNoticeGRPCServiceServer) HandleNoticeGroupNotifyHonor(context.Context, *model.EventNoticeGroupNotifyHonorGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleNoticeGroupNotifyHonor not implemented")
}
func (*UnimplementedOnebotEventNoticeGRPCServiceServer) HandleNoticeGroupNotifyLuckyKing(context.Context, *model.EventNoticeGroupNotifyLuckyKingGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleNoticeGroupNotifyLuckyKing not implemented")
}

func RegisterOnebotEventNoticeGRPCServiceServer(s *grpc.Server, srv OnebotEventNoticeGRPCServiceServer) {
	s.RegisterService(&_OnebotEventNoticeGRPCService_serviceDesc, srv)
}

func _OnebotEventNoticeGRPCService_HandleNoticeGroupUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventNoticeGroupUploadGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupUpload(ctx, req.(*model.EventNoticeGroupUploadGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotEventNoticeGRPCService_HandleNoticeGroupAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventNoticeGroupAdminGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupAdmin(ctx, req.(*model.EventNoticeGroupAdminGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotEventNoticeGRPCService_HandleNoticeGroupDecrease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventNoticeGroupDecreaseGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupDecrease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupDecrease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupDecrease(ctx, req.(*model.EventNoticeGroupDecreaseGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotEventNoticeGRPCService_HandleNoticeGroupIncrease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventNoticeGroupIncreaseGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupIncrease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupIncrease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupIncrease(ctx, req.(*model.EventNoticeGroupIncreaseGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotEventNoticeGRPCService_HandleNoticeGroupBan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventNoticeGroupBanGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupBan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupBan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupBan(ctx, req.(*model.EventNoticeGroupBanGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotEventNoticeGRPCService_HandleNoticeFriendAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventNoticeFriendAddGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeFriendAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.OnebotEventNoticeGRPCService/HandleNoticeFriendAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeFriendAdd(ctx, req.(*model.EventNoticeFriendAddGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotEventNoticeGRPCService_HandleNoticeGroupRecall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventNoticeGroupRecallGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupRecall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupRecall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupRecall(ctx, req.(*model.EventNoticeGroupRecallGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotEventNoticeGRPCService_HandleNoticeFriendRecall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventNoticeFriendRecallGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeFriendRecall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.OnebotEventNoticeGRPCService/HandleNoticeFriendRecall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeFriendRecall(ctx, req.(*model.EventNoticeFriendRecallGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotEventNoticeGRPCService_HandleNoticeGroupNotifyPoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventNoticeGroupNotifyPokeGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupNotifyPoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupNotifyPoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupNotifyPoke(ctx, req.(*model.EventNoticeGroupNotifyPokeGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotEventNoticeGRPCService_HandleNoticeGroupNotifyHonor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventNoticeGroupNotifyHonorGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupNotifyHonor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupNotifyHonor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupNotifyHonor(ctx, req.(*model.EventNoticeGroupNotifyHonorGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotEventNoticeGRPCService_HandleNoticeGroupNotifyLuckyKing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventNoticeGroupNotifyLuckyKingGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupNotifyLuckyKing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notice.OnebotEventNoticeGRPCService/HandleNoticeGroupNotifyLuckyKing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotEventNoticeGRPCServiceServer).HandleNoticeGroupNotifyLuckyKing(ctx, req.(*model.EventNoticeGroupNotifyLuckyKingGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnebotEventNoticeGRPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notice.OnebotEventNoticeGRPCService",
	HandlerType: (*OnebotEventNoticeGRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleNoticeGroupUpload",
			Handler:    _OnebotEventNoticeGRPCService_HandleNoticeGroupUpload_Handler,
		},
		{
			MethodName: "HandleNoticeGroupAdmin",
			Handler:    _OnebotEventNoticeGRPCService_HandleNoticeGroupAdmin_Handler,
		},
		{
			MethodName: "HandleNoticeGroupDecrease",
			Handler:    _OnebotEventNoticeGRPCService_HandleNoticeGroupDecrease_Handler,
		},
		{
			MethodName: "HandleNoticeGroupIncrease",
			Handler:    _OnebotEventNoticeGRPCService_HandleNoticeGroupIncrease_Handler,
		},
		{
			MethodName: "HandleNoticeGroupBan",
			Handler:    _OnebotEventNoticeGRPCService_HandleNoticeGroupBan_Handler,
		},
		{
			MethodName: "HandleNoticeFriendAdd",
			Handler:    _OnebotEventNoticeGRPCService_HandleNoticeFriendAdd_Handler,
		},
		{
			MethodName: "HandleNoticeGroupRecall",
			Handler:    _OnebotEventNoticeGRPCService_HandleNoticeGroupRecall_Handler,
		},
		{
			MethodName: "HandleNoticeFriendRecall",
			Handler:    _OnebotEventNoticeGRPCService_HandleNoticeFriendRecall_Handler,
		},
		{
			MethodName: "HandleNoticeGroupNotifyPoke",
			Handler:    _OnebotEventNoticeGRPCService_HandleNoticeGroupNotifyPoke_Handler,
		},
		{
			MethodName: "HandleNoticeGroupNotifyHonor",
			Handler:    _OnebotEventNoticeGRPCService_HandleNoticeGroupNotifyHonor_Handler,
		},
		{
			MethodName: "HandleNoticeGroupNotifyLuckyKing",
			Handler:    _OnebotEventNoticeGRPCService_HandleNoticeGroupNotifyLuckyKing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event_notice.proto",
}
