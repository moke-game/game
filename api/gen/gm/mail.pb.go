// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: gm/mail.proto

package pb

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

type SendMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Sign string `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *SendMailRequest) Reset() {
	*x = SendMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gm_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailRequest) ProtoMessage() {}

func (x *SendMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gm_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailRequest.ProtoReflect.Descriptor instead.
func (*SendMailRequest) Descriptor() ([]byte, []int) {
	return file_gm_mail_proto_rawDescGZIP(), []int{0}
}

func (x *SendMailRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *SendMailRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

type MailSendData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelRange string   `protobuf:"bytes,1,opt,name=channel_range,json=channelRange,proto3" json:"channel_range,omitempty"`   // channel range (1:全部渠道,2:指定渠道)
	ChannelId    string   `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`            // channel id
	SendType     string   `protobuf:"bytes,3,opt,name=send_type,json=sendType,proto3" json:"send_type,omitempty"`               // send type (1:全服发送,2:玩家发送)
	PlatformId   string   `protobuf:"bytes,4,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`         // platform id
	ServerId     string   `protobuf:"bytes,5,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`               // server id
	LevelType    string   `protobuf:"bytes,6,opt,name=level_type,json=levelType,proto3" json:"level_type,omitempty"`            // level type (1:全部等级,2:限定等级)
	LevelStart   string   `protobuf:"bytes,7,opt,name=level_start,json=levelStart,proto3" json:"level_start,omitempty"`         // level start
	LevelEnd     string   `protobuf:"bytes,8,opt,name=level_end,json=levelEnd,proto3" json:"level_end,omitempty"`               // level end
	RoleId       string   `protobuf:"bytes,9,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`                     // role id
	Sender       string   `protobuf:"bytes,10,opt,name=sender,proto3" json:"sender,omitempty"`                                  // sender
	Title        string   `protobuf:"bytes,11,opt,name=title,proto3" json:"title,omitempty"`                                    // mail title
	Content      string   `protobuf:"bytes,12,opt,name=content,proto3" json:"content,omitempty"`                                // mail content (country_code:title:content;country_code:title:content;...)
	Items        []string `protobuf:"bytes,13,rep,name=items,proto3" json:"items,omitempty"`                                    // mail rewards
	StartTime    string   `protobuf:"bytes,14,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`           // mail start time
	EndTime      string   `protobuf:"bytes,15,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`                 // mail end time
	ThemeId      int32    `protobuf:"varint,16,opt,name=theme_id,json=themeId,proto3" json:"theme_id,omitempty"`                // mail theme id
	RegisterTime int64    `protobuf:"varint,17,opt,name=register_time,json=registerTime,proto3" json:"register_time,omitempty"` // register time (-1: 不限制,0:当前时间,>0:指定时间)
}

func (x *MailSendData) Reset() {
	*x = MailSendData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gm_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailSendData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailSendData) ProtoMessage() {}

func (x *MailSendData) ProtoReflect() protoreflect.Message {
	mi := &file_gm_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailSendData.ProtoReflect.Descriptor instead.
func (*MailSendData) Descriptor() ([]byte, []int) {
	return file_gm_mail_proto_rawDescGZIP(), []int{1}
}

func (x *MailSendData) GetChannelRange() string {
	if x != nil {
		return x.ChannelRange
	}
	return ""
}

func (x *MailSendData) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *MailSendData) GetSendType() string {
	if x != nil {
		return x.SendType
	}
	return ""
}

func (x *MailSendData) GetPlatformId() string {
	if x != nil {
		return x.PlatformId
	}
	return ""
}

func (x *MailSendData) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *MailSendData) GetLevelType() string {
	if x != nil {
		return x.LevelType
	}
	return ""
}

func (x *MailSendData) GetLevelStart() string {
	if x != nil {
		return x.LevelStart
	}
	return ""
}

func (x *MailSendData) GetLevelEnd() string {
	if x != nil {
		return x.LevelEnd
	}
	return ""
}

func (x *MailSendData) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *MailSendData) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MailSendData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MailSendData) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MailSendData) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MailSendData) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *MailSendData) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *MailSendData) GetThemeId() int32 {
	if x != nil {
		return x.ThemeId
	}
	return 0
}

func (x *MailSendData) GetRegisterTime() int64 {
	if x != nil {
		return x.RegisterTime
	}
	return 0
}

type MailSendResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleName string `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *MailSendResult) Reset() {
	*x = MailSendResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gm_mail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailSendResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailSendResult) ProtoMessage() {}

func (x *MailSendResult) ProtoReflect() protoreflect.Message {
	mi := &file_gm_mail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailSendResult.ProtoReflect.Descriptor instead.
func (*MailSendResult) Descriptor() ([]byte, []int) {
	return file_gm_mail_proto_rawDescGZIP(), []int{2}
}

func (x *MailSendResult) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MailSendResult) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type SendMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // send status
	Code   string          `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`     // send code
	Info   string          `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`     // send error info
	Data   *MailSendResult `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`     // response result
}

func (x *SendMailResponse) Reset() {
	*x = SendMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gm_mail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailResponse) ProtoMessage() {}

func (x *SendMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gm_mail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailResponse.ProtoReflect.Descriptor instead.
func (*SendMailResponse) Descriptor() ([]byte, []int) {
	return file_gm_mail_proto_rawDescGZIP(), []int{3}
}

func (x *SendMailResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SendMailResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SendMailResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *SendMailResponse) GetData() *MailSendResult {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_gm_mail_proto protoreflect.FileDescriptor

var file_gm_mail_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x6d, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x67, 0x6d, 0x2e, 0x70, 0x62, 0x22, 0x39, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67,
	0x6e, 0x22, 0xfb, 0x03, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x45, 0x6e, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x46, 0x0a, 0x0e, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7d, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6d, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gm_mail_proto_rawDescOnce sync.Once
	file_gm_mail_proto_rawDescData = file_gm_mail_proto_rawDesc
)

func file_gm_mail_proto_rawDescGZIP() []byte {
	file_gm_mail_proto_rawDescOnce.Do(func() {
		file_gm_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_gm_mail_proto_rawDescData)
	})
	return file_gm_mail_proto_rawDescData
}

var file_gm_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gm_mail_proto_goTypes = []any{
	(*SendMailRequest)(nil),  // 0: gm.pb.SendMailRequest
	(*MailSendData)(nil),     // 1: gm.pb.MailSendData
	(*MailSendResult)(nil),   // 2: gm.pb.MailSendResult
	(*SendMailResponse)(nil), // 3: gm.pb.SendMailResponse
}
var file_gm_mail_proto_depIdxs = []int32{
	2, // 0: gm.pb.SendMailResponse.data:type_name -> gm.pb.MailSendResult
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gm_mail_proto_init() }
func file_gm_mail_proto_init() {
	if File_gm_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gm_mail_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SendMailRequest); i {
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
		file_gm_mail_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MailSendData); i {
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
		file_gm_mail_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MailSendResult); i {
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
		file_gm_mail_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SendMailResponse); i {
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
			RawDescriptor: file_gm_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gm_mail_proto_goTypes,
		DependencyIndexes: file_gm_mail_proto_depIdxs,
		MessageInfos:      file_gm_mail_proto_msgTypes,
	}.Build()
	File_gm_mail_proto = out.File
	file_gm_mail_proto_rawDesc = nil
	file_gm_mail_proto_goTypes = nil
	file_gm_mail_proto_depIdxs = nil
}
