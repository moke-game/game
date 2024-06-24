// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: gm/notice.proto

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

type SendCycleNoticeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Sign string `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *SendCycleNoticeRequest) Reset() {
	*x = SendCycleNoticeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gm_notice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCycleNoticeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCycleNoticeRequest) ProtoMessage() {}

func (x *SendCycleNoticeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gm_notice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCycleNoticeRequest.ProtoReflect.Descriptor instead.
func (*SendCycleNoticeRequest) Descriptor() ([]byte, []int) {
	return file_gm_notice_proto_rawDescGZIP(), []int{0}
}

func (x *SendCycleNoticeRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *SendCycleNoticeRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

type CycleNoticeReqMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelRange string `protobuf:"bytes,1,opt,name=channel_range,json=channelRange,proto3" json:"channel_range,omitempty"`
	ChannelId    string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	NoticeRange  string `protobuf:"bytes,3,opt,name=notice_range,json=noticeRange,proto3" json:"notice_range,omitempty"`
	PlatformId   string `protobuf:"bytes,4,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
	ServerId     string `protobuf:"bytes,5,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Content      string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	StartTime    string `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime      string `protobuf:"bytes,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	CycleTime    string `protobuf:"bytes,9,opt,name=cycle_time,json=cycleTime,proto3" json:"cycle_time,omitempty"` //循环时间 秒
	CycleNum     string `protobuf:"bytes,10,opt,name=cycle_num,json=cycleNum,proto3" json:"cycle_num,omitempty"`   //循环次数
}

func (x *CycleNoticeReqMsg) Reset() {
	*x = CycleNoticeReqMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gm_notice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CycleNoticeReqMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CycleNoticeReqMsg) ProtoMessage() {}

func (x *CycleNoticeReqMsg) ProtoReflect() protoreflect.Message {
	mi := &file_gm_notice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CycleNoticeReqMsg.ProtoReflect.Descriptor instead.
func (*CycleNoticeReqMsg) Descriptor() ([]byte, []int) {
	return file_gm_notice_proto_rawDescGZIP(), []int{1}
}

func (x *CycleNoticeReqMsg) GetChannelRange() string {
	if x != nil {
		return x.ChannelRange
	}
	return ""
}

func (x *CycleNoticeReqMsg) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *CycleNoticeReqMsg) GetNoticeRange() string {
	if x != nil {
		return x.NoticeRange
	}
	return ""
}

func (x *CycleNoticeReqMsg) GetPlatformId() string {
	if x != nil {
		return x.PlatformId
	}
	return ""
}

func (x *CycleNoticeReqMsg) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *CycleNoticeReqMsg) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CycleNoticeReqMsg) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *CycleNoticeReqMsg) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *CycleNoticeReqMsg) GetCycleTime() string {
	if x != nil {
		return x.CycleTime
	}
	return ""
}

func (x *CycleNoticeReqMsg) GetCycleNum() string {
	if x != nil {
		return x.CycleNum
	}
	return ""
}

type CycleNoticeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CycleTime int64             `protobuf:"varint,1,opt,name=cycle_time,json=cycleTime,proto3" json:"cycle_time,omitempty"`
	CycleNum  int32             `protobuf:"varint,2,opt,name=cycle_num,json=cycleNum,proto3" json:"cycle_num,omitempty"`
	Notice    map[string]string `protobuf:"bytes,3,rep,name=notice,proto3" json:"notice,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CycleNoticeInfo) Reset() {
	*x = CycleNoticeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gm_notice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CycleNoticeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CycleNoticeInfo) ProtoMessage() {}

func (x *CycleNoticeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gm_notice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CycleNoticeInfo.ProtoReflect.Descriptor instead.
func (*CycleNoticeInfo) Descriptor() ([]byte, []int) {
	return file_gm_notice_proto_rawDescGZIP(), []int{2}
}

func (x *CycleNoticeInfo) GetCycleTime() int64 {
	if x != nil {
		return x.CycleTime
	}
	return 0
}

func (x *CycleNoticeInfo) GetCycleNum() int32 {
	if x != nil {
		return x.CycleNum
	}
	return 0
}

func (x *CycleNoticeInfo) GetNotice() map[string]string {
	if x != nil {
		return x.Notice
	}
	return nil
}

type SendCycleNoticeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Code   string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Info   string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SendCycleNoticeResponse) Reset() {
	*x = SendCycleNoticeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gm_notice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCycleNoticeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCycleNoticeResponse) ProtoMessage() {}

func (x *SendCycleNoticeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gm_notice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCycleNoticeResponse.ProtoReflect.Descriptor instead.
func (*SendCycleNoticeResponse) Descriptor() ([]byte, []int) {
	return file_gm_notice_proto_rawDescGZIP(), []int{3}
}

func (x *SendCycleNoticeResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SendCycleNoticeResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SendCycleNoticeResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_gm_notice_proto protoreflect.FileDescriptor

var file_gm_notice_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x6d, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x67, 0x6d, 0x2e, 0x70, 0x62, 0x22, 0x40, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x79, 0x63, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22, 0xc8, 0x02, 0x0a, 0x11, 0x43,
	0x79, 0x63, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x4d, 0x73, 0x67,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0xc4, 0x01, 0x0a, 0x0f, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x3a, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x79,
	0x63, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x59, 0x0a, 0x17,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gm_notice_proto_rawDescOnce sync.Once
	file_gm_notice_proto_rawDescData = file_gm_notice_proto_rawDesc
)

func file_gm_notice_proto_rawDescGZIP() []byte {
	file_gm_notice_proto_rawDescOnce.Do(func() {
		file_gm_notice_proto_rawDescData = protoimpl.X.CompressGZIP(file_gm_notice_proto_rawDescData)
	})
	return file_gm_notice_proto_rawDescData
}

var file_gm_notice_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gm_notice_proto_goTypes = []interface{}{
	(*SendCycleNoticeRequest)(nil),  // 0: gm.pb.SendCycleNoticeRequest
	(*CycleNoticeReqMsg)(nil),       // 1: gm.pb.CycleNoticeReqMsg
	(*CycleNoticeInfo)(nil),         // 2: gm.pb.CycleNoticeInfo
	(*SendCycleNoticeResponse)(nil), // 3: gm.pb.SendCycleNoticeResponse
	nil,                             // 4: gm.pb.CycleNoticeInfo.NoticeEntry
}
var file_gm_notice_proto_depIdxs = []int32{
	4, // 0: gm.pb.CycleNoticeInfo.notice:type_name -> gm.pb.CycleNoticeInfo.NoticeEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gm_notice_proto_init() }
func file_gm_notice_proto_init() {
	if File_gm_notice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gm_notice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCycleNoticeRequest); i {
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
		file_gm_notice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CycleNoticeReqMsg); i {
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
		file_gm_notice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CycleNoticeInfo); i {
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
		file_gm_notice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCycleNoticeResponse); i {
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
			RawDescriptor: file_gm_notice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gm_notice_proto_goTypes,
		DependencyIndexes: file_gm_notice_proto_depIdxs,
		MessageInfos:      file_gm_notice_proto_msgTypes,
	}.Build()
	File_gm_notice_proto = out.File
	file_gm_notice_proto_rawDesc = nil
	file_gm_notice_proto_goTypes = nil
	file_gm_notice_proto_depIdxs = nil
}