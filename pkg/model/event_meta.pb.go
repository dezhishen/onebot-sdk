// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: event_meta.proto

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

type EventMetaBaseGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventBase     *EventBaseGRPC `protobuf:"bytes,1,opt,name=eventBase,proto3" json:"eventBase,omitempty"`
	MetaEventType string         `protobuf:"bytes,2,opt,name=meta_event_type,json=metaEventType,proto3" json:"meta_event_type,omitempty"`
}

func (x *EventMetaBaseGRPC) Reset() {
	*x = EventMetaBaseGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMetaBaseGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetaBaseGRPC) ProtoMessage() {}

func (x *EventMetaBaseGRPC) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EventMetaBaseGRPC.ProtoReflect.Descriptor instead.
func (*EventMetaBaseGRPC) Descriptor() ([]byte, []int) {
	return file_event_meta_proto_rawDescGZIP(), []int{0}
}

func (x *EventMetaBaseGRPC) GetEventBase() *EventBaseGRPC {
	if x != nil {
		return x.EventBase
	}
	return nil
}

func (x *EventMetaBaseGRPC) GetMetaEventType() string {
	if x != nil {
		return x.MetaEventType
	}
	return ""
}

type EventMetaLifecycleGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventMetaBase *EventMetaBaseGRPC `protobuf:"bytes,1,opt,name=eventMetaBase,proto3" json:"eventMetaBase,omitempty"`
	//
	//enable、disable、connect	事件子类型，分别表示 OneBot 启用、停用、WebSocket 连接成功
	//
	//注意，目前生命周期元事件中，只有 HTTP POST 的情况下可以收到 enable 和 disable
	//
	//只有正向 WebSocket 和反向 WebSocket 可以收到 connect。
	SubType string `protobuf:"bytes,2,opt,name=sub_type,json=subType,proto3" json:"sub_type,omitempty"`
}

func (x *EventMetaLifecycleGRPC) Reset() {
	*x = EventMetaLifecycleGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMetaLifecycleGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetaLifecycleGRPC) ProtoMessage() {}

func (x *EventMetaLifecycleGRPC) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EventMetaLifecycleGRPC.ProtoReflect.Descriptor instead.
func (*EventMetaLifecycleGRPC) Descriptor() ([]byte, []int) {
	return file_event_meta_proto_rawDescGZIP(), []int{1}
}

func (x *EventMetaLifecycleGRPC) GetEventMetaBase() *EventMetaBaseGRPC {
	if x != nil {
		return x.EventMetaBase
	}
	return nil
}

func (x *EventMetaLifecycleGRPC) GetSubType() string {
	if x != nil {
		return x.SubType
	}
	return ""
}

type EventMetaHeartbeatStatusStatGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketReceived  uint64 `protobuf:"varint,1,opt,name=packet_received,json=packetReceived,proto3" json:"packet_received,omitempty"`
	PacketSent      uint64 `protobuf:"varint,2,opt,name=packet_sent,json=packetSent,proto3" json:"packet_sent,omitempty"`
	PacketLost      uint64 `protobuf:"varint,3,opt,name=packet_lost,json=packetLost,proto3" json:"packet_lost,omitempty"`
	MessageReceived uint64 `protobuf:"varint,4,opt,name=message_received,json=messageReceived,proto3" json:"message_received,omitempty"`
	MessageSent     uint64 `protobuf:"varint,5,opt,name=message_sent,json=messageSent,proto3" json:"message_sent,omitempty"`
	DisconnectTimes uint64 `protobuf:"varint,6,opt,name=disconnect_times,json=disconnectTimes,proto3" json:"disconnect_times,omitempty"`
	LostTimes       uint64 `protobuf:"varint,7,opt,name=lost_times,json=lostTimes,proto3" json:"lost_times,omitempty"`
	LastMessageTime uint64 `protobuf:"varint,8,opt,name=last_message_time,json=lastMessageTime,proto3" json:"last_message_time,omitempty"`
}

func (x *EventMetaHeartbeatStatusStatGRPC) Reset() {
	*x = EventMetaHeartbeatStatusStatGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_meta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMetaHeartbeatStatusStatGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetaHeartbeatStatusStatGRPC) ProtoMessage() {}

func (x *EventMetaHeartbeatStatusStatGRPC) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EventMetaHeartbeatStatusStatGRPC.ProtoReflect.Descriptor instead.
func (*EventMetaHeartbeatStatusStatGRPC) Descriptor() ([]byte, []int) {
	return file_event_meta_proto_rawDescGZIP(), []int{2}
}

func (x *EventMetaHeartbeatStatusStatGRPC) GetPacketReceived() uint64 {
	if x != nil {
		return x.PacketReceived
	}
	return 0
}

func (x *EventMetaHeartbeatStatusStatGRPC) GetPacketSent() uint64 {
	if x != nil {
		return x.PacketSent
	}
	return 0
}

func (x *EventMetaHeartbeatStatusStatGRPC) GetPacketLost() uint64 {
	if x != nil {
		return x.PacketLost
	}
	return 0
}

func (x *EventMetaHeartbeatStatusStatGRPC) GetMessageReceived() uint64 {
	if x != nil {
		return x.MessageReceived
	}
	return 0
}

func (x *EventMetaHeartbeatStatusStatGRPC) GetMessageSent() uint64 {
	if x != nil {
		return x.MessageSent
	}
	return 0
}

func (x *EventMetaHeartbeatStatusStatGRPC) GetDisconnectTimes() uint64 {
	if x != nil {
		return x.DisconnectTimes
	}
	return 0
}

func (x *EventMetaHeartbeatStatusStatGRPC) GetLostTimes() uint64 {
	if x != nil {
		return x.LostTimes
	}
	return 0
}

func (x *EventMetaHeartbeatStatusStatGRPC) GetLastMessageTime() uint64 {
	if x != nil {
		return x.LastMessageTime
	}
	return 0
}

type EventMetaHeartbeatStatusGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppEnabled     bool                              `protobuf:"varint,1,opt,name=app_enabled,json=appEnabled,proto3" json:"app_enabled,omitempty"`
	AppGood        bool                              `protobuf:"varint,2,opt,name=app_good,json=appGood,proto3" json:"app_good,omitempty"`
	AppInitialized bool                              `protobuf:"varint,3,opt,name=app_initialized,json=appInitialized,proto3" json:"app_initialized,omitempty"`
	Good           bool                              `protobuf:"varint,4,opt,name=good,proto3" json:"good,omitempty"`
	Online         bool                              `protobuf:"varint,5,opt,name=online,proto3" json:"online,omitempty"`
	PluginsGood    bool                              `protobuf:"varint,6,opt,name=plugins_good,json=pluginsGood,proto3" json:"plugins_good,omitempty"`
	Stat           *EventMetaHeartbeatStatusStatGRPC `protobuf:"bytes,7,opt,name=stat,proto3" json:"stat,omitempty"`
}

func (x *EventMetaHeartbeatStatusGRPC) Reset() {
	*x = EventMetaHeartbeatStatusGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_meta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMetaHeartbeatStatusGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetaHeartbeatStatusGRPC) ProtoMessage() {}

func (x *EventMetaHeartbeatStatusGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_event_meta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetaHeartbeatStatusGRPC.ProtoReflect.Descriptor instead.
func (*EventMetaHeartbeatStatusGRPC) Descriptor() ([]byte, []int) {
	return file_event_meta_proto_rawDescGZIP(), []int{3}
}

func (x *EventMetaHeartbeatStatusGRPC) GetAppEnabled() bool {
	if x != nil {
		return x.AppEnabled
	}
	return false
}

func (x *EventMetaHeartbeatStatusGRPC) GetAppGood() bool {
	if x != nil {
		return x.AppGood
	}
	return false
}

func (x *EventMetaHeartbeatStatusGRPC) GetAppInitialized() bool {
	if x != nil {
		return x.AppInitialized
	}
	return false
}

func (x *EventMetaHeartbeatStatusGRPC) GetGood() bool {
	if x != nil {
		return x.Good
	}
	return false
}

func (x *EventMetaHeartbeatStatusGRPC) GetOnline() bool {
	if x != nil {
		return x.Online
	}
	return false
}

func (x *EventMetaHeartbeatStatusGRPC) GetPluginsGood() bool {
	if x != nil {
		return x.PluginsGood
	}
	return false
}

func (x *EventMetaHeartbeatStatusGRPC) GetStat() *EventMetaHeartbeatStatusStatGRPC {
	if x != nil {
		return x.Stat
	}
	return nil
}

type EventMetaHeartbeatGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventMetaBase *EventMetaBaseGRPC            `protobuf:"bytes,1,opt,name=eventMetaBase,proto3" json:"eventMetaBase,omitempty"`
	Status        *EventMetaHeartbeatStatusGRPC `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	//到下次心跳的间隔，单位毫秒
	Interval uint64 `protobuf:"varint,3,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *EventMetaHeartbeatGRPC) Reset() {
	*x = EventMetaHeartbeatGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_meta_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMetaHeartbeatGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetaHeartbeatGRPC) ProtoMessage() {}

func (x *EventMetaHeartbeatGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_event_meta_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetaHeartbeatGRPC.ProtoReflect.Descriptor instead.
func (*EventMetaHeartbeatGRPC) Descriptor() ([]byte, []int) {
	return file_event_meta_proto_rawDescGZIP(), []int{4}
}

func (x *EventMetaHeartbeatGRPC) GetEventMetaBase() *EventMetaBaseGRPC {
	if x != nil {
		return x.EventMetaBase
	}
	return nil
}

func (x *EventMetaHeartbeatGRPC) GetStatus() *EventMetaHeartbeatStatusGRPC {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *EventMetaHeartbeatGRPC) GetInterval() uint64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

var File_event_meta_proto protoreflect.FileDescriptor

var file_event_meta_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x47, 0x52, 0x50, 0x43, 0x12, 0x32, 0x0a, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65,
	0x47, 0x52, 0x50, 0x43, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x73, 0x0a, 0x16, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x52, 0x50,
	0x43, 0x12, 0x3e, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x61,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x47, 0x52,
	0x50, 0x43, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x61, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd1, 0x02, 0x0a,
	0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x47, 0x52, 0x50,
	0x43, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6c, 0x6f, 0x73, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0f, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x8f, 0x02, 0x0a, 0x1c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x47, 0x52, 0x50,
	0x43, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x67, 0x6f, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x6f, 0x6f, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x67, 0x6f, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x5f, 0x67, 0x6f,
	0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x3b, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x73, 0x74,
	0x61, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x16, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x47, 0x52, 0x50, 0x43, 0x12, 0x3e, 0x0a,
	0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x47, 0x52, 0x50, 0x43, 0x52, 0x0d,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x47, 0x52,
	0x50, 0x43, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x7a, 0x68, 0x69, 0x73, 0x68, 0x65, 0x6e, 0x2f, 0x6f,
	0x6e, 0x65, 0x62, 0x6f, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_event_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_event_meta_proto_goTypes = []interface{}{
	(*EventMetaBaseGRPC)(nil),                // 0: model.EventMetaBaseGRPC
	(*EventMetaLifecycleGRPC)(nil),           // 1: model.EventMetaLifecycleGRPC
	(*EventMetaHeartbeatStatusStatGRPC)(nil), // 2: model.EventMetaHeartbeatStatusStatGRPC
	(*EventMetaHeartbeatStatusGRPC)(nil),     // 3: model.EventMetaHeartbeatStatusGRPC
	(*EventMetaHeartbeatGRPC)(nil),           // 4: model.EventMetaHeartbeatGRPC
	(*EventBaseGRPC)(nil),                    // 5: model.EventBaseGRPC
}
var file_event_meta_proto_depIdxs = []int32{
	5, // 0: model.EventMetaBaseGRPC.eventBase:type_name -> model.EventBaseGRPC
	0, // 1: model.EventMetaLifecycleGRPC.eventMetaBase:type_name -> model.EventMetaBaseGRPC
	2, // 2: model.EventMetaHeartbeatStatusGRPC.stat:type_name -> model.EventMetaHeartbeatStatusStatGRPC
	0, // 3: model.EventMetaHeartbeatGRPC.eventMetaBase:type_name -> model.EventMetaBaseGRPC
	3, // 4: model.EventMetaHeartbeatGRPC.status:type_name -> model.EventMetaHeartbeatStatusGRPC
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_event_meta_proto_init() }
func file_event_meta_proto_init() {
	if File_event_meta_proto != nil {
		return
	}
	file_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMetaBaseGRPC); i {
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
			switch v := v.(*EventMetaLifecycleGRPC); i {
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
			switch v := v.(*EventMetaHeartbeatStatusStatGRPC); i {
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
		file_event_meta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMetaHeartbeatStatusGRPC); i {
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
		file_event_meta_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMetaHeartbeatGRPC); i {
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
			NumMessages:   5,
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
