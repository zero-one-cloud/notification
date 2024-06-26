// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: notification/v1/notification.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SendSmsVerifyCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 如果需要带区号则自己修改校验条件
	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *SendSmsVerifyCodeReq) Reset() {
	*x = SendSmsVerifyCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_v1_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsVerifyCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsVerifyCodeReq) ProtoMessage() {}

func (x *SendSmsVerifyCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsVerifyCodeReq.ProtoReflect.Descriptor instead.
func (*SendSmsVerifyCodeReq) Descriptor() ([]byte, []int) {
	return file_notification_v1_notification_proto_rawDescGZIP(), []int{0}
}

func (x *SendSmsVerifyCodeReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type SendSmsVerifyCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmsId int64 `protobuf:"varint,1,opt,name=sms_id,json=smsId,proto3" json:"sms_id,omitempty"`
	// 状态 0待发送 1已发送 2发送失败
	Status int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Code   string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *SendSmsVerifyCodeReply) Reset() {
	*x = SendSmsVerifyCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_v1_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsVerifyCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsVerifyCodeReply) ProtoMessage() {}

func (x *SendSmsVerifyCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsVerifyCodeReply.ProtoReflect.Descriptor instead.
func (*SendSmsVerifyCodeReply) Descriptor() ([]byte, []int) {
	return file_notification_v1_notification_proto_rawDescGZIP(), []int{1}
}

func (x *SendSmsVerifyCodeReply) GetSmsId() int64 {
	if x != nil {
		return x.SmsId
	}
	return 0
}

func (x *SendSmsVerifyCodeReply) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SendSmsVerifyCodeReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type SendSmsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendSmsReq) Reset() {
	*x = SendSmsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_v1_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsReq) ProtoMessage() {}

func (x *SendSmsReq) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsReq.ProtoReflect.Descriptor instead.
func (*SendSmsReq) Descriptor() ([]byte, []int) {
	return file_notification_v1_notification_proto_rawDescGZIP(), []int{2}
}

type SendSmsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendSmsReply) Reset() {
	*x = SendSmsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_v1_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsReply) ProtoMessage() {}

func (x *SendSmsReply) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsReply.ProtoReflect.Descriptor instead.
func (*SendSmsReply) Descriptor() ([]byte, []int) {
	return file_notification_v1_notification_proto_rawDescGZIP(), []int{3}
}

type SendEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// 主题
	Subject string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	// 内容
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendEmailReq) Reset() {
	*x = SendEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_v1_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailReq) ProtoMessage() {}

func (x *SendEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailReq.ProtoReflect.Descriptor instead.
func (*SendEmailReq) Descriptor() ([]byte, []int) {
	return file_notification_v1_notification_proto_rawDescGZIP(), []int{4}
}

func (x *SendEmailReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SendEmailReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SendEmailReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendEmailReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendEmailReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SendEmailReply) Reset() {
	*x = SendEmailReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_v1_notification_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailReply) ProtoMessage() {}

func (x *SendEmailReply) ProtoReflect() protoreflect.Message {
	mi := &file_notification_v1_notification_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailReply.ProtoReflect.Descriptor instead.
func (*SendEmailReply) Descriptor() ([]byte, []int) {
	return file_notification_v1_notification_proto_rawDescGZIP(), []int{5}
}

func (x *SendEmailReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_notification_v1_notification_proto protoreflect.FileDescriptor

var file_notification_v1_notification_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x36, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01,
	0x0b, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x5b, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64,
	0x53, 0x6d, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6d, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x6d, 0x73, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x9b, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x02, 0x18, 0x40, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x21,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xed, 0x02,
	0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d,
	0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x6d, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10,
	0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x63, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x12, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x73, 0x6d, 0x73,
	0x2f, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x66, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0b, 0x3a, 0x01, 0x2a, 0x22, 0x06, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x18, 0x5a,
	0x16, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_v1_notification_proto_rawDescOnce sync.Once
	file_notification_v1_notification_proto_rawDescData = file_notification_v1_notification_proto_rawDesc
)

func file_notification_v1_notification_proto_rawDescGZIP() []byte {
	file_notification_v1_notification_proto_rawDescOnce.Do(func() {
		file_notification_v1_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_v1_notification_proto_rawDescData)
	})
	return file_notification_v1_notification_proto_rawDescData
}

var file_notification_v1_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_notification_v1_notification_proto_goTypes = []interface{}{
	(*SendSmsVerifyCodeReq)(nil),   // 0: api.notification.v1.SendSmsVerifyCodeReq
	(*SendSmsVerifyCodeReply)(nil), // 1: api.notification.v1.SendSmsVerifyCodeReply
	(*SendSmsReq)(nil),             // 2: api.notification.v1.SendSmsReq
	(*SendSmsReply)(nil),           // 3: api.notification.v1.SendSmsReply
	(*SendEmailReq)(nil),           // 4: api.notification.v1.SendEmailReq
	(*SendEmailReply)(nil),         // 5: api.notification.v1.SendEmailReply
}
var file_notification_v1_notification_proto_depIdxs = []int32{
	0, // 0: api.notification.v1.NotificationService.SendSmsVerifyCode:input_type -> api.notification.v1.SendSmsVerifyCodeReq
	2, // 1: api.notification.v1.NotificationService.SendSms:input_type -> api.notification.v1.SendSmsReq
	4, // 2: api.notification.v1.NotificationService.SendEmail:input_type -> api.notification.v1.SendEmailReq
	1, // 3: api.notification.v1.NotificationService.SendSmsVerifyCode:output_type -> api.notification.v1.SendSmsVerifyCodeReply
	3, // 4: api.notification.v1.NotificationService.SendSms:output_type -> api.notification.v1.SendSmsReply
	5, // 5: api.notification.v1.NotificationService.SendEmail:output_type -> api.notification.v1.SendEmailReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notification_v1_notification_proto_init() }
func file_notification_v1_notification_proto_init() {
	if File_notification_v1_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_v1_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsVerifyCodeReq); i {
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
		file_notification_v1_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsVerifyCodeReply); i {
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
		file_notification_v1_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsReq); i {
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
		file_notification_v1_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsReply); i {
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
		file_notification_v1_notification_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailReq); i {
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
		file_notification_v1_notification_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailReply); i {
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
			RawDescriptor: file_notification_v1_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notification_v1_notification_proto_goTypes,
		DependencyIndexes: file_notification_v1_notification_proto_depIdxs,
		MessageInfos:      file_notification_v1_notification_proto_msgTypes,
	}.Build()
	File_notification_v1_notification_proto = out.File
	file_notification_v1_notification_proto_rawDesc = nil
	file_notification_v1_notification_proto_goTypes = nil
	file_notification_v1_notification_proto_depIdxs = nil
}
