// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: account.proto

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

type AccountGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *AccountGRPC) Reset() {
	*x = AccountGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountGRPC) ProtoMessage() {}

func (x *AccountGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountGRPC.ProtoReflect.Descriptor instead.
func (*AccountGRPC) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *AccountGRPC) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AccountGRPC) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type AccountResultGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    *AccountGRPC `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Retcode int64        `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Status  string       `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Msg     string       `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	Wording string       `protobuf:"bytes,5,opt,name=wording,proto3" json:"wording,omitempty"`
}

func (x *AccountResultGRPC) Reset() {
	*x = AccountResultGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountResultGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountResultGRPC) ProtoMessage() {}

func (x *AccountResultGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountResultGRPC.ProtoReflect.Descriptor instead.
func (*AccountResultGRPC) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountResultGRPC) GetData() *AccountGRPC {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AccountResultGRPC) GetRetcode() int64 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *AccountResultGRPC) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AccountResultGRPC) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AccountResultGRPC) GetWording() string {
	if x != nil {
		return x.Wording
	}
	return ""
}

type QQProfileGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname     string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Company      string `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	Email        string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Collage      string `protobuf:"bytes,4,opt,name=collage,proto3" json:"collage,omitempty"`
	PersonalNote string `protobuf:"bytes,5,opt,name=personal_note,json=personalNote,proto3" json:"personal_note,omitempty"`
}

func (x *QQProfileGRPC) Reset() {
	*x = QQProfileGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QQProfileGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QQProfileGRPC) ProtoMessage() {}

func (x *QQProfileGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QQProfileGRPC.ProtoReflect.Descriptor instead.
func (*QQProfileGRPC) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
}

func (x *QQProfileGRPC) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *QQProfileGRPC) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *QQProfileGRPC) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *QQProfileGRPC) GetCollage() string {
	if x != nil {
		return x.Collage
	}
	return ""
}

func (x *QQProfileGRPC) GetPersonalNote() string {
	if x != nil {
		return x.PersonalNote
	}
	return ""
}

type QQProfileResultGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    *QQProfileGRPC `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Retcode int64          `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Status  string         `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Msg     string         `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	Wording string         `protobuf:"bytes,5,opt,name=wording,proto3" json:"wording,omitempty"`
}

func (x *QQProfileResultGRPC) Reset() {
	*x = QQProfileResultGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QQProfileResultGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QQProfileResultGRPC) ProtoMessage() {}

func (x *QQProfileResultGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QQProfileResultGRPC.ProtoReflect.Descriptor instead.
func (*QQProfileResultGRPC) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{3}
}

func (x *QQProfileResultGRPC) GetData() *QQProfileGRPC {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QQProfileResultGRPC) GetRetcode() int64 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *QQProfileResultGRPC) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *QQProfileResultGRPC) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *QQProfileResultGRPC) GetWording() string {
	if x != nil {
		return x.Wording
	}
	return ""
}

type VariantGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelShow string `protobuf:"bytes,1,opt,name=model_show,json=modelShow,proto3" json:"model_show,omitempty"`
	NeedPay   bool   `protobuf:"varint,2,opt,name=need_pay,json=needPay,proto3" json:"need_pay,omitempty"`
}

func (x *VariantGRPC) Reset() {
	*x = VariantGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariantGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariantGRPC) ProtoMessage() {}

func (x *VariantGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariantGRPC.ProtoReflect.Descriptor instead.
func (*VariantGRPC) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{4}
}

func (x *VariantGRPC) GetModelShow() string {
	if x != nil {
		return x.ModelShow
	}
	return ""
}

func (x *VariantGRPC) GetNeedPay() bool {
	if x != nil {
		return x.NeedPay
	}
	return false
}

type ModelShowGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variants []*VariantGRPC `protobuf:"bytes,1,rep,name=variants,proto3" json:"variants,omitempty"`
}

func (x *ModelShowGRPC) Reset() {
	*x = ModelShowGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelShowGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelShowGRPC) ProtoMessage() {}

func (x *ModelShowGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelShowGRPC.ProtoReflect.Descriptor instead.
func (*ModelShowGRPC) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{5}
}

func (x *ModelShowGRPC) GetVariants() []*VariantGRPC {
	if x != nil {
		return x.Variants
	}
	return nil
}

type ModelShowResultGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    *ModelShowGRPC `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Retcode int64          `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Status  string         `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Msg     string         `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	Wording string         `protobuf:"bytes,5,opt,name=wording,proto3" json:"wording,omitempty"`
}

func (x *ModelShowResultGRPC) Reset() {
	*x = ModelShowResultGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelShowResultGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelShowResultGRPC) ProtoMessage() {}

func (x *ModelShowResultGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelShowResultGRPC.ProtoReflect.Descriptor instead.
func (*ModelShowResultGRPC) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{6}
}

func (x *ModelShowResultGRPC) GetData() *ModelShowGRPC {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ModelShowResultGRPC) GetRetcode() int64 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ModelShowResultGRPC) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ModelShowResultGRPC) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ModelShowResultGRPC) GetWording() string {
	if x != nil {
		return x.Wording
	}
	return ""
}

type DeviceGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId      int64  `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	DeviceName string `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	DeviceKind string `protobuf:"bytes,3,opt,name=device_kind,json=deviceKind,proto3" json:"device_kind,omitempty"`
}

func (x *DeviceGRPC) Reset() {
	*x = DeviceGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceGRPC) ProtoMessage() {}

func (x *DeviceGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceGRPC.ProtoReflect.Descriptor instead.
func (*DeviceGRPC) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{7}
}

func (x *DeviceGRPC) GetAppId() int64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *DeviceGRPC) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *DeviceGRPC) GetDeviceKind() string {
	if x != nil {
		return x.DeviceKind
	}
	return ""
}

type OnlineClientsGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*DeviceGRPC `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *OnlineClientsGRPC) Reset() {
	*x = OnlineClientsGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineClientsGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineClientsGRPC) ProtoMessage() {}

func (x *OnlineClientsGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineClientsGRPC.ProtoReflect.Descriptor instead.
func (*OnlineClientsGRPC) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{8}
}

func (x *OnlineClientsGRPC) GetClients() []*DeviceGRPC {
	if x != nil {
		return x.Clients
	}
	return nil
}

type OnlineClientsResultGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    *OnlineClientsGRPC `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Retcode int64              `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Status  string             `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Msg     string             `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	Wording string             `protobuf:"bytes,5,opt,name=wording,proto3" json:"wording,omitempty"`
}

func (x *OnlineClientsResultGRPC) Reset() {
	*x = OnlineClientsResultGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineClientsResultGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineClientsResultGRPC) ProtoMessage() {}

func (x *OnlineClientsResultGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineClientsResultGRPC.ProtoReflect.Descriptor instead.
func (*OnlineClientsResultGRPC) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{9}
}

func (x *OnlineClientsResultGRPC) GetData() *OnlineClientsGRPC {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *OnlineClientsResultGRPC) GetRetcode() int64 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *OnlineClientsResultGRPC) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OnlineClientsResultGRPC) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *OnlineClientsResultGRPC) GetWording() string {
	if x != nil {
		return x.Wording
	}
	return ""
}

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x42, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x47, 0x52, 0x50, 0x43, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x11, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x47, 0x52, 0x50, 0x43,
	0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x52,
	0x50, 0x43, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x9a, 0x01, 0x0a, 0x0d, 0x51, 0x51, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x47, 0x52, 0x50, 0x43, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e,
	0x6f, 0x74, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x13, 0x51, 0x51, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x47, 0x52, 0x50, 0x43, 0x12, 0x28, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x51, 0x51, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x52, 0x50, 0x43, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x22, 0x47, 0x0a, 0x0b, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x47, 0x52,
	0x50, 0x43, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x68, 0x6f,
	0x77, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x65, 0x65, 0x64, 0x50, 0x61, 0x79, 0x22, 0x3f, 0x0a, 0x0d,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x68, 0x6f, 0x77, 0x47, 0x52, 0x50, 0x43, 0x12, 0x2e, 0x0a,
	0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x47,
	0x52, 0x50, 0x43, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x9d, 0x01,
	0x0a, 0x13, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x47, 0x52, 0x50, 0x43, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x53, 0x68, 0x6f, 0x77, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x65, 0x0a,
	0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x52, 0x50, 0x43, 0x12, 0x15, 0x0a, 0x06, 0x61,
	0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4b, 0x69, 0x6e, 0x64, 0x22, 0x40, 0x0a, 0x11, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x47, 0x52, 0x50, 0x43, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x52, 0x50, 0x43, 0x52, 0x07, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x17, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x47, 0x52,
	0x50, 0x43, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x7a,
	0x68, 0x69, 0x73, 0x68, 0x65, 0x6e, 0x2f, 0x6f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3b, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData = file_account_proto_rawDesc
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_proto_rawDescData)
	})
	return file_account_proto_rawDescData
}

var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_account_proto_goTypes = []interface{}{
	(*AccountGRPC)(nil),             // 0: model.AccountGRPC
	(*AccountResultGRPC)(nil),       // 1: model.AccountResultGRPC
	(*QQProfileGRPC)(nil),           // 2: model.QQProfileGRPC
	(*QQProfileResultGRPC)(nil),     // 3: model.QQProfileResultGRPC
	(*VariantGRPC)(nil),             // 4: model.VariantGRPC
	(*ModelShowGRPC)(nil),           // 5: model.ModelShowGRPC
	(*ModelShowResultGRPC)(nil),     // 6: model.ModelShowResultGRPC
	(*DeviceGRPC)(nil),              // 7: model.DeviceGRPC
	(*OnlineClientsGRPC)(nil),       // 8: model.OnlineClientsGRPC
	(*OnlineClientsResultGRPC)(nil), // 9: model.OnlineClientsResultGRPC
}
var file_account_proto_depIdxs = []int32{
	0, // 0: model.AccountResultGRPC.data:type_name -> model.AccountGRPC
	2, // 1: model.QQProfileResultGRPC.data:type_name -> model.QQProfileGRPC
	4, // 2: model.ModelShowGRPC.variants:type_name -> model.VariantGRPC
	5, // 3: model.ModelShowResultGRPC.data:type_name -> model.ModelShowGRPC
	7, // 4: model.OnlineClientsGRPC.clients:type_name -> model.DeviceGRPC
	8, // 5: model.OnlineClientsResultGRPC.data:type_name -> model.OnlineClientsGRPC
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountGRPC); i {
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
		file_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountResultGRPC); i {
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
		file_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QQProfileGRPC); i {
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
		file_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QQProfileResultGRPC); i {
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
		file_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariantGRPC); i {
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
		file_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelShowGRPC); i {
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
		file_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelShowResultGRPC); i {
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
		file_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceGRPC); i {
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
		file_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineClientsGRPC); i {
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
		file_account_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineClientsResultGRPC); i {
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
			RawDescriptor: file_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		MessageInfos:      file_account_proto_msgTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_rawDesc = nil
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}
