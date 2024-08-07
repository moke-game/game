// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: bff/chat.proto

package bffpb

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

type ChatType int32

const (
	ChatType_CHAT_WORLD  ChatType = 0 //世界聊天
	ChatType_CHAT_PLAYER ChatType = 1 //私聊
	ChatType_CHAT_TEAM   ChatType = 2 //组队聊天
)

// Enum value maps for ChatType.
var (
	ChatType_name = map[int32]string{
		0: "CHAT_WORLD",
		1: "CHAT_PLAYER",
		2: "CHAT_TEAM",
	}
	ChatType_value = map[string]int32{
		"CHAT_WORLD":  0,
		"CHAT_PLAYER": 1,
		"CHAT_TEAM":   2,
	}
)

func (x ChatType) Enum() *ChatType {
	p := new(ChatType)
	*p = x
	return p
}

func (x ChatType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatType) Descriptor() protoreflect.EnumDescriptor {
	return file_bff_chat_proto_enumTypes[0].Descriptor()
}

func (ChatType) Type() protoreflect.EnumType {
	return &file_bff_chat_proto_enumTypes[0]
}

func (x ChatType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatType.Descriptor instead.
func (ChatType) EnumDescriptor() ([]byte, []int) {
	return file_bff_chat_proto_rawDescGZIP(), []int{0}
}

// 聊天信息
type ChatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendUid    int64   `protobuf:"varint,1,opt,name=send_uid,json=sendUid,proto3" json:"send_uid,omitempty"`          //发送方id
	ReceiveUid int64   `protobuf:"varint,2,opt,name=receive_uid,json=receiveUid,proto3" json:"receive_uid,omitempty"` //接收方id（世界聊天和组队聊天不需要发）
	Content    string  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`                          //文本内容
	VoiceUrl   string  `protobuf:"bytes,4,opt,name=voice_url,json=voiceUrl,proto3" json:"voice_url,omitempty"`        //语音url
	VoiceTime  float32 `protobuf:"fixed32,5,opt,name=voice_time,json=voiceTime,proto3" json:"voice_time,omitempty"`   //语音条时长
	EmojiId    int32   `protobuf:"varint,6,opt,name=emoji_id,json=emojiId,proto3" json:"emoji_id,omitempty"`          //表情
}

func (x *ChatInfo) Reset() {
	*x = ChatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInfo) ProtoMessage() {}

func (x *ChatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bff_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInfo.ProtoReflect.Descriptor instead.
func (*ChatInfo) Descriptor() ([]byte, []int) {
	return file_bff_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatInfo) GetSendUid() int64 {
	if x != nil {
		return x.SendUid
	}
	return 0
}

func (x *ChatInfo) GetReceiveUid() int64 {
	if x != nil {
		return x.ReceiveUid
	}
	return 0
}

func (x *ChatInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatInfo) GetVoiceUrl() string {
	if x != nil {
		return x.VoiceUrl
	}
	return ""
}

func (x *ChatInfo) GetVoiceTime() float32 {
	if x != nil {
		return x.VoiceTime
	}
	return 0
}

func (x *ChatInfo) GetEmojiId() int32 {
	if x != nil {
		return x.EmojiId
	}
	return 0
}

// 聊天信息
type C2SCHATMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatType ChatType  `protobuf:"varint,1,opt,name=chat_type,json=chatType,proto3,enum=bff.ChatType" json:"chat_type,omitempty"` //频道
	ChatInfo *ChatInfo `protobuf:"bytes,2,opt,name=chat_info,json=chatInfo,proto3" json:"chat_info,omitempty"`
}

func (x *C2SCHATMessage) Reset() {
	*x = C2SCHATMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SCHATMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SCHATMessage) ProtoMessage() {}

func (x *C2SCHATMessage) ProtoReflect() protoreflect.Message {
	mi := &file_bff_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SCHATMessage.ProtoReflect.Descriptor instead.
func (*C2SCHATMessage) Descriptor() ([]byte, []int) {
	return file_bff_chat_proto_rawDescGZIP(), []int{1}
}

func (x *C2SCHATMessage) GetChatType() ChatType {
	if x != nil {
		return x.ChatType
	}
	return ChatType_CHAT_WORLD
}

func (x *C2SCHATMessage) GetChatInfo() *ChatInfo {
	if x != nil {
		return x.ChatInfo
	}
	return nil
}

type S2CCHATMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *S2CCHATMessage) Reset() {
	*x = S2CCHATMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CCHATMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CCHATMessage) ProtoMessage() {}

func (x *S2CCHATMessage) ProtoReflect() protoreflect.Message {
	mi := &file_bff_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CCHATMessage.ProtoReflect.Descriptor instead.
func (*S2CCHATMessage) Descriptor() ([]byte, []int) {
	return file_bff_chat_proto_rawDescGZIP(), []int{2}
}

// 订阅频道
type C2SSubChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatType ChatType `protobuf:"varint,1,opt,name=chat_type,json=chatType,proto3,enum=bff.ChatType" json:"chat_type,omitempty"` //频道
}

func (x *C2SSubChannel) Reset() {
	*x = C2SSubChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SSubChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SSubChannel) ProtoMessage() {}

func (x *C2SSubChannel) ProtoReflect() protoreflect.Message {
	mi := &file_bff_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SSubChannel.ProtoReflect.Descriptor instead.
func (*C2SSubChannel) Descriptor() ([]byte, []int) {
	return file_bff_chat_proto_rawDescGZIP(), []int{3}
}

func (x *C2SSubChannel) GetChatType() ChatType {
	if x != nil {
		return x.ChatType
	}
	return ChatType_CHAT_WORLD
}

type S2CSubChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *S2CSubChannel) Reset() {
	*x = S2CSubChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CSubChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CSubChannel) ProtoMessage() {}

func (x *S2CSubChannel) ProtoReflect() protoreflect.Message {
	mi := &file_bff_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CSubChannel.ProtoReflect.Descriptor instead.
func (*S2CSubChannel) Descriptor() ([]byte, []int) {
	return file_bff_chat_proto_rawDescGZIP(), []int{4}
}

// 世界聊天信息返回
type S2CCHATReceiveWorldMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatInfo   *ChatInfo         `protobuf:"bytes,1,opt,name=chat_info,json=chatInfo,proto3" json:"chat_info,omitempty"`
	SendPlayer *PlayerSimpleInfo `protobuf:"bytes,2,opt,name=send_player,json=sendPlayer,proto3" json:"send_player,omitempty"` //发送者信息
}

func (x *S2CCHATReceiveWorldMessage) Reset() {
	*x = S2CCHATReceiveWorldMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CCHATReceiveWorldMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CCHATReceiveWorldMessage) ProtoMessage() {}

func (x *S2CCHATReceiveWorldMessage) ProtoReflect() protoreflect.Message {
	mi := &file_bff_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CCHATReceiveWorldMessage.ProtoReflect.Descriptor instead.
func (*S2CCHATReceiveWorldMessage) Descriptor() ([]byte, []int) {
	return file_bff_chat_proto_rawDescGZIP(), []int{5}
}

func (x *S2CCHATReceiveWorldMessage) GetChatInfo() *ChatInfo {
	if x != nil {
		return x.ChatInfo
	}
	return nil
}

func (x *S2CCHATReceiveWorldMessage) GetSendPlayer() *PlayerSimpleInfo {
	if x != nil {
		return x.SendPlayer
	}
	return nil
}

// 组队聊天信息返回
type S2CCHATReceiveTeamMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatInfo   *ChatInfo         `protobuf:"bytes,1,opt,name=chat_info,json=chatInfo,proto3" json:"chat_info,omitempty"`
	SendPlayer *PlayerSimpleInfo `protobuf:"bytes,2,opt,name=send_player,json=sendPlayer,proto3" json:"send_player,omitempty"` //发送者信息
}

func (x *S2CCHATReceiveTeamMessage) Reset() {
	*x = S2CCHATReceiveTeamMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CCHATReceiveTeamMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CCHATReceiveTeamMessage) ProtoMessage() {}

func (x *S2CCHATReceiveTeamMessage) ProtoReflect() protoreflect.Message {
	mi := &file_bff_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CCHATReceiveTeamMessage.ProtoReflect.Descriptor instead.
func (*S2CCHATReceiveTeamMessage) Descriptor() ([]byte, []int) {
	return file_bff_chat_proto_rawDescGZIP(), []int{6}
}

func (x *S2CCHATReceiveTeamMessage) GetChatInfo() *ChatInfo {
	if x != nil {
		return x.ChatInfo
	}
	return nil
}

func (x *S2CCHATReceiveTeamMessage) GetSendPlayer() *PlayerSimpleInfo {
	if x != nil {
		return x.SendPlayer
	}
	return nil
}

// 单人聊天信息返回
type S2CCHATReceivePlayerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatInfo   *ChatInfo         `protobuf:"bytes,1,opt,name=chat_info,json=chatInfo,proto3" json:"chat_info,omitempty"`
	SendPlayer *PlayerSimpleInfo `protobuf:"bytes,2,opt,name=send_player,json=sendPlayer,proto3" json:"send_player,omitempty"` //发送者信息
}

func (x *S2CCHATReceivePlayerMessage) Reset() {
	*x = S2CCHATReceivePlayerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CCHATReceivePlayerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CCHATReceivePlayerMessage) ProtoMessage() {}

func (x *S2CCHATReceivePlayerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_bff_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CCHATReceivePlayerMessage.ProtoReflect.Descriptor instead.
func (*S2CCHATReceivePlayerMessage) Descriptor() ([]byte, []int) {
	return file_bff_chat_proto_rawDescGZIP(), []int{7}
}

func (x *S2CCHATReceivePlayerMessage) GetChatInfo() *ChatInfo {
	if x != nil {
		return x.ChatInfo
	}
	return nil
}

func (x *S2CCHATReceivePlayerMessage) GetSendPlayer() *PlayerSimpleInfo {
	if x != nil {
		return x.SendPlayer
	}
	return nil
}

// 获取拥有的表情
type C2SChatGetEmoji struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2SChatGetEmoji) Reset() {
	*x = C2SChatGetEmoji{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SChatGetEmoji) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SChatGetEmoji) ProtoMessage() {}

func (x *C2SChatGetEmoji) ProtoReflect() protoreflect.Message {
	mi := &file_bff_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SChatGetEmoji.ProtoReflect.Descriptor instead.
func (*C2SChatGetEmoji) Descriptor() ([]byte, []int) {
	return file_bff_chat_proto_rawDescGZIP(), []int{8}
}

// 返回拥有的表情
type S2CChatGetEmoji struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmojiId []int32 `protobuf:"varint,1,rep,packed,name=emoji_id,json=emojiId,proto3" json:"emoji_id,omitempty"` //已解锁的表情ID
}

func (x *S2CChatGetEmoji) Reset() {
	*x = S2CChatGetEmoji{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_chat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CChatGetEmoji) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CChatGetEmoji) ProtoMessage() {}

func (x *S2CChatGetEmoji) ProtoReflect() protoreflect.Message {
	mi := &file_bff_chat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CChatGetEmoji.ProtoReflect.Descriptor instead.
func (*S2CChatGetEmoji) Descriptor() ([]byte, []int) {
	return file_bff_chat_proto_rawDescGZIP(), []int{9}
}

func (x *S2CChatGetEmoji) GetEmojiId() []int32 {
	if x != nil {
		return x.EmojiId
	}
	return nil
}

var File_bff_chat_proto protoreflect.FileDescriptor

var file_bff_chat_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x66, 0x66, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x62, 0x66, 0x66, 0x1a, 0x11, 0x62, 0x66, 0x66, 0x2f, 0x62, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x08, 0x43, 0x68, 0x61,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x55, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x55, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6d, 0x6f, 0x6a, 0x69,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6d, 0x6f, 0x6a, 0x69,
	0x49, 0x64, 0x22, 0x68, 0x0a, 0x0e, 0x43, 0x32, 0x53, 0x43, 0x48, 0x41, 0x54, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2a, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x10, 0x0a, 0x0e,
	0x53, 0x32, 0x43, 0x43, 0x48, 0x41, 0x54, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3b,
	0x0a, 0x0d, 0x43, 0x32, 0x53, 0x53, 0x75, 0x62, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x2a, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x53,
	0x32, 0x43, 0x53, 0x75, 0x62, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x80, 0x01, 0x0a,
	0x1a, 0x53, 0x32, 0x43, 0x43, 0x48, 0x41, 0x54, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x62, 0x66, 0x66, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63,
	0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x5f,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62,
	0x66, 0x66, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22,
	0x7f, 0x0a, 0x19, 0x53, 0x32, 0x43, 0x43, 0x48, 0x41, 0x54, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x63, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x62, 0x66, 0x66, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x22, 0x81, 0x01, 0x0a, 0x1b, 0x53, 0x32, 0x43, 0x43, 0x48, 0x41, 0x54, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2a, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0b,
	0x73, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x32, 0x53, 0x43, 0x68, 0x61, 0x74, 0x47,
	0x65, 0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x22, 0x2c, 0x0a, 0x0f, 0x53, 0x32, 0x43, 0x43, 0x68,
	0x61, 0x74, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6d,
	0x6f, 0x6a, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6d,
	0x6f, 0x6a, 0x69, 0x49, 0x64, 0x2a, 0x3a, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x10,
	0x02, 0x42, 0x0c, 0x5a, 0x0a, 0x62, 0x66, 0x66, 0x2f, 0x3b, 0x62, 0x66, 0x66, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bff_chat_proto_rawDescOnce sync.Once
	file_bff_chat_proto_rawDescData = file_bff_chat_proto_rawDesc
)

func file_bff_chat_proto_rawDescGZIP() []byte {
	file_bff_chat_proto_rawDescOnce.Do(func() {
		file_bff_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_bff_chat_proto_rawDescData)
	})
	return file_bff_chat_proto_rawDescData
}

var file_bff_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bff_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_bff_chat_proto_goTypes = []any{
	(ChatType)(0),                       // 0: bff.ChatType
	(*ChatInfo)(nil),                    // 1: bff.ChatInfo
	(*C2SCHATMessage)(nil),              // 2: bff.C2SCHATMessage
	(*S2CCHATMessage)(nil),              // 3: bff.S2CCHATMessage
	(*C2SSubChannel)(nil),               // 4: bff.C2SSubChannel
	(*S2CSubChannel)(nil),               // 5: bff.S2CSubChannel
	(*S2CCHATReceiveWorldMessage)(nil),  // 6: bff.S2CCHATReceiveWorldMessage
	(*S2CCHATReceiveTeamMessage)(nil),   // 7: bff.S2CCHATReceiveTeamMessage
	(*S2CCHATReceivePlayerMessage)(nil), // 8: bff.S2CCHATReceivePlayerMessage
	(*C2SChatGetEmoji)(nil),             // 9: bff.C2SChatGetEmoji
	(*S2CChatGetEmoji)(nil),             // 10: bff.S2CChatGetEmoji
	(*PlayerSimpleInfo)(nil),            // 11: bff.PlayerSimpleInfo
}
var file_bff_chat_proto_depIdxs = []int32{
	0,  // 0: bff.C2SCHATMessage.chat_type:type_name -> bff.ChatType
	1,  // 1: bff.C2SCHATMessage.chat_info:type_name -> bff.ChatInfo
	0,  // 2: bff.C2SSubChannel.chat_type:type_name -> bff.ChatType
	1,  // 3: bff.S2CCHATReceiveWorldMessage.chat_info:type_name -> bff.ChatInfo
	11, // 4: bff.S2CCHATReceiveWorldMessage.send_player:type_name -> bff.PlayerSimpleInfo
	1,  // 5: bff.S2CCHATReceiveTeamMessage.chat_info:type_name -> bff.ChatInfo
	11, // 6: bff.S2CCHATReceiveTeamMessage.send_player:type_name -> bff.PlayerSimpleInfo
	1,  // 7: bff.S2CCHATReceivePlayerMessage.chat_info:type_name -> bff.ChatInfo
	11, // 8: bff.S2CCHATReceivePlayerMessage.send_player:type_name -> bff.PlayerSimpleInfo
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_bff_chat_proto_init() }
func file_bff_chat_proto_init() {
	if File_bff_chat_proto != nil {
		return
	}
	file_bff_bcommon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bff_chat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ChatInfo); i {
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
		file_bff_chat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*C2SCHATMessage); i {
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
		file_bff_chat_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*S2CCHATMessage); i {
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
		file_bff_chat_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*C2SSubChannel); i {
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
		file_bff_chat_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*S2CSubChannel); i {
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
		file_bff_chat_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*S2CCHATReceiveWorldMessage); i {
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
		file_bff_chat_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*S2CCHATReceiveTeamMessage); i {
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
		file_bff_chat_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*S2CCHATReceivePlayerMessage); i {
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
		file_bff_chat_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*C2SChatGetEmoji); i {
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
		file_bff_chat_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*S2CChatGetEmoji); i {
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
			RawDescriptor: file_bff_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bff_chat_proto_goTypes,
		DependencyIndexes: file_bff_chat_proto_depIdxs,
		EnumInfos:         file_bff_chat_proto_enumTypes,
		MessageInfos:      file_bff_chat_proto_msgTypes,
	}.Build()
	File_bff_chat_proto = out.File
	file_bff_chat_proto_rawDesc = nil
	file_bff_chat_proto_goTypes = nil
	file_bff_chat_proto_depIdxs = nil
}
