// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: record.proto

package model

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

type RecordGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File      string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	OutFormat string `protobuf:"bytes,2,opt,name=out_format,json=outFormat,proto3" json:"out_format,omitempty"`
}

func (x *RecordGRPC) Reset() {
	*x = RecordGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordGRPC) ProtoMessage() {}

func (x *RecordGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordGRPC.ProtoReflect.Descriptor instead.
func (*RecordGRPC) Descriptor() ([]byte, []int) {
	return file_record_proto_rawDescGZIP(), []int{0}
}

func (x *RecordGRPC) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *RecordGRPC) GetOutFormat() string {
	if x != nil {
		return x.OutFormat
	}
	return ""
}

type RecordResultGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    *RecordGRPC `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Retcode int64       `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Status  string      `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Msg     string      `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	Wording string      `protobuf:"bytes,5,opt,name=wording,proto3" json:"wording,omitempty"`
}

func (x *RecordResultGRPC) Reset() {
	*x = RecordResultGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordResultGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordResultGRPC) ProtoMessage() {}

func (x *RecordResultGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordResultGRPC.ProtoReflect.Descriptor instead.
func (*RecordResultGRPC) Descriptor() ([]byte, []int) {
	return file_record_proto_rawDescGZIP(), []int{1}
}

func (x *RecordResultGRPC) GetData() *RecordGRPC {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RecordResultGRPC) GetRetcode() int64 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *RecordResultGRPC) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RecordResultGRPC) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RecordResultGRPC) GetWording() string {
	if x != nil {
		return x.Wording
	}
	return ""
}

var File_record_proto protoreflect.FileDescriptor

var file_record_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x3f, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x47,
	0x52, 0x50, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x5f, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x75, 0x74,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x47, 0x52, 0x50, 0x43, 0x12, 0x25, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x65, 0x7a, 0x68, 0x69, 0x73, 0x68, 0x65, 0x6e, 0x2f, 0x6f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3b, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_record_proto_rawDescOnce sync.Once
	file_record_proto_rawDescData = file_record_proto_rawDesc
)

func file_record_proto_rawDescGZIP() []byte {
	file_record_proto_rawDescOnce.Do(func() {
		file_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_record_proto_rawDescData)
	})
	return file_record_proto_rawDescData
}

var file_record_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_record_proto_goTypes = []interface{}{
	(*RecordGRPC)(nil),       // 0: model.RecordGRPC
	(*RecordResultGRPC)(nil), // 1: model.RecordResultGRPC
}
var file_record_proto_depIdxs = []int32{
	0, // 0: model.RecordResultGRPC.data:type_name -> model.RecordGRPC
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_record_proto_init() }
func file_record_proto_init() {
	if File_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordGRPC); i {
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
		file_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordResultGRPC); i {
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
			RawDescriptor: file_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_record_proto_goTypes,
		DependencyIndexes: file_record_proto_depIdxs,
		MessageInfos:      file_record_proto_msgTypes,
	}.Build()
	File_record_proto = out.File
	file_record_proto_rawDesc = nil
	file_record_proto_goTypes = nil
	file_record_proto_depIdxs = nil
}