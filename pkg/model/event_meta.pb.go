// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: event_meta.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventMetaBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventBase     *EventBase `protobuf:"bytes,1,opt,name=eventBase,proto3" json:"eventBase,omitempty"`
	MetaEventType string     `protobuf:"bytes,2,opt,name=meta_event_type,json=metaEventType,proto3" json:"meta_event_type,omitempty"`
}

func (x *EventMetaBase) Reset() {
	*x = EventMetaBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMetaBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetaBase) ProtoMessage() {}

func (x *EventMetaBase) ProtoReflect() protoreflect.Message {
	mi := &file_event_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetaBase.ProtoReflect.Descriptor instead.
func (*EventMetaBase) Descriptor() ([]byte, []int) {
	return file_event_meta_proto_rawDescGZIP(), []int{0}
}

func (x *EventMetaBase) GetEventBase() *EventBase {
	if x != nil {
		return x.EventBase
	}
	return nil
}

func (x *EventMetaBase) GetMetaEventType() string {
	if x != nil {
		return x.MetaEventType
	}
	return ""
}

type EventMetaLifecycle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventBase *EventBase `protobuf:"bytes,1,opt,name=eventBase,proto3" json:"eventBase,omitempty"`
	//
	//enable、disable、connect	事件子类型，分别表示 OneBot 启用、停用、WebSocket 连接成功
	//
	//注意，目前生命周期元事件中，只有 HTTP POST 的情况下可以收到 enable 和 disable
	//
	//只有正向 WebSocket 和反向 WebSocket 可以收到 connect。
	SubType string `protobuf:"bytes,2,opt,name=sub_type,json=subType,proto3" json:"sub_type,omitempty"`
}

func (x *EventMetaLifecycle) Reset() {
	*x = EventMetaLifecycle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMetaLifecycle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetaLifecycle) ProtoMessage() {}

func (x *EventMetaLifecycle) ProtoReflect() protoreflect.Message {
	mi := &file_event_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetaLifecycle.ProtoReflect.Descriptor instead.
func (*EventMetaLifecycle) Descriptor() ([]byte, []int) {
	return file_event_meta_proto_rawDescGZIP(), []int{1}
}

func (x *EventMetaLifecycle) GetEventBase() *EventBase {
	if x != nil {
		return x.EventBase
	}
	return nil
}

func (x *EventMetaLifecycle) GetSubType() string {
	if x != nil {
		return x.SubType
	}
	return ""
}

type EventMetaHeartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventBase *EventBase `protobuf:"bytes,1,opt,name=eventBase,proto3" json:"eventBase,omitempty"`
	Status    *anypb.Any `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	//到下次心跳的间隔，单位毫秒
	Interval int64 `protobuf:"varint,3,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *EventMetaHeartbeat) Reset() {
	*x = EventMetaHeartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_meta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMetaHeartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetaHeartbeat) ProtoMessage() {}

func (x *EventMetaHeartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_event_meta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetaHeartbeat.ProtoReflect.Descriptor instead.
func (*EventMetaHeartbeat) Descriptor() ([]byte, []int) {
	return file_event_meta_proto_rawDescGZIP(), []int{2}
}

func (x *EventMetaHeartbeat) GetEventBase() *EventBase {
	if x != nil {
		return x.EventBase
	}
	return nil
}

func (x *EventMetaHeartbeat) GetStatus() *anypb.Any {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *EventMetaHeartbeat) GetInterval() int64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

var File_event_meta_proto protoreflect.FileDescriptor

var file_event_meta_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x67, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x61,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61,
	0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x74,
	0x61, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5f, 0x0a, 0x12, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x12, 0x2e, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x42, 0x61, 0x73, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x12,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_meta_proto_rawDescOnce sync.Once
	file_event_meta_proto_rawDescData = file_event_meta_proto_rawDesc
)

func file_event_meta_proto_rawDescGZIP() []byte {
	file_event_meta_proto_rawDescOnce.Do(func() {
		file_event_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_meta_proto_rawDescData)
	})
	return file_event_meta_proto_rawDescData
}

var file_event_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_event_meta_proto_goTypes = []interface{}{
	(*EventMetaBase)(nil),      // 0: model.EventMetaBase
	(*EventMetaLifecycle)(nil), // 1: model.EventMetaLifecycle
	(*EventMetaHeartbeat)(nil), // 2: model.EventMetaHeartbeat
	(*EventBase)(nil),          // 3: model.EventBase
	(*anypb.Any)(nil),          // 4: google.protobuf.Any
}
var file_event_meta_proto_depIdxs = []int32{
	3, // 0: model.EventMetaBase.eventBase:type_name -> model.EventBase
	3, // 1: model.EventMetaLifecycle.eventBase:type_name -> model.EventBase
	3, // 2: model.EventMetaHeartbeat.eventBase:type_name -> model.EventBase
	4, // 3: model.EventMetaHeartbeat.status:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_event_meta_proto_init() }
func file_event_meta_proto_init() {
	if File_event_meta_proto != nil {
		return
	}
	file_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMetaBase); i {
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
		file_event_meta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMetaLifecycle); i {
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
		file_event_meta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMetaHeartbeat); i {
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
			RawDescriptor: file_event_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_meta_proto_goTypes,
		DependencyIndexes: file_event_meta_proto_depIdxs,
		MessageInfos:      file_event_meta_proto_msgTypes,
	}.Build()
	File_event_meta_proto = out.File
	file_event_meta_proto_rawDesc = nil
	file_event_meta_proto_goTypes = nil
	file_event_meta_proto_depIdxs = nil
}
