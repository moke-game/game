// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: bff/matchmaking.proto

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

// 组队开始进入匹配 通知客户端组队进入匹配中
type NtfMatchingGroupStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayId int32 `protobuf:"varint,1,opt,name=play_id,json=playId,proto3" json:"play_id,omitempty"` //玩法id
}

func (x *NtfMatchingGroupStart) Reset() {
	*x = NtfMatchingGroupStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_matchmaking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NtfMatchingGroupStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NtfMatchingGroupStart) ProtoMessage() {}

func (x *NtfMatchingGroupStart) ProtoReflect() protoreflect.Message {
	mi := &file_bff_matchmaking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NtfMatchingGroupStart.ProtoReflect.Descriptor instead.
func (*NtfMatchingGroupStart) Descriptor() ([]byte, []int) {
	return file_bff_matchmaking_proto_rawDescGZIP(), []int{0}
}

func (x *NtfMatchingGroupStart) GetPlayId() int32 {
	if x != nil {
		return x.PlayId
	}
	return 0
}

// 匹配--个人
type C2SMatchingSingleStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayMap int32 `protobuf:"varint,1,opt,name=playMap,proto3" json:"playMap,omitempty"` //玩法id
}

func (x *C2SMatchingSingleStart) Reset() {
	*x = C2SMatchingSingleStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_matchmaking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SMatchingSingleStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SMatchingSingleStart) ProtoMessage() {}

func (x *C2SMatchingSingleStart) ProtoReflect() protoreflect.Message {
	mi := &file_bff_matchmaking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SMatchingSingleStart.ProtoReflect.Descriptor instead.
func (*C2SMatchingSingleStart) Descriptor() ([]byte, []int) {
	return file_bff_matchmaking_proto_rawDescGZIP(), []int{1}
}

func (x *C2SMatchingSingleStart) GetPlayMap() int32 {
	if x != nil {
		return x.PlayMap
	}
	return 0
}

type S2CMatchingSingleStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayMap int32 `protobuf:"varint,1,opt,name=playMap,proto3" json:"playMap,omitempty"` //玩法id
}

func (x *S2CMatchingSingleStart) Reset() {
	*x = S2CMatchingSingleStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_matchmaking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CMatchingSingleStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CMatchingSingleStart) ProtoMessage() {}

func (x *S2CMatchingSingleStart) ProtoReflect() protoreflect.Message {
	mi := &file_bff_matchmaking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CMatchingSingleStart.ProtoReflect.Descriptor instead.
func (*S2CMatchingSingleStart) Descriptor() ([]byte, []int) {
	return file_bff_matchmaking_proto_rawDescGZIP(), []int{2}
}

func (x *S2CMatchingSingleStart) GetPlayMap() int32 {
	if x != nil {
		return x.PlayMap
	}
	return 0
}

// 取消匹配
type C2SMatchingCancel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2SMatchingCancel) Reset() {
	*x = C2SMatchingCancel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_matchmaking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SMatchingCancel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SMatchingCancel) ProtoMessage() {}

func (x *C2SMatchingCancel) ProtoReflect() protoreflect.Message {
	mi := &file_bff_matchmaking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SMatchingCancel.ProtoReflect.Descriptor instead.
func (*C2SMatchingCancel) Descriptor() ([]byte, []int) {
	return file_bff_matchmaking_proto_rawDescGZIP(), []int{3}
}

type S2CMatchingCancel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *S2CMatchingCancel) Reset() {
	*x = S2CMatchingCancel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_matchmaking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CMatchingCancel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CMatchingCancel) ProtoMessage() {}

func (x *S2CMatchingCancel) ProtoReflect() protoreflect.Message {
	mi := &file_bff_matchmaking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CMatchingCancel.ProtoReflect.Descriptor instead.
func (*S2CMatchingCancel) Descriptor() ([]byte, []int) {
	return file_bff_matchmaking_proto_rawDescGZIP(), []int{4}
}

// 查询匹配状态
type C2SMatchingStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2SMatchingStatus) Reset() {
	*x = C2SMatchingStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_matchmaking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SMatchingStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SMatchingStatus) ProtoMessage() {}

func (x *C2SMatchingStatus) ProtoReflect() protoreflect.Message {
	mi := &file_bff_matchmaking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SMatchingStatus.ProtoReflect.Descriptor instead.
func (*C2SMatchingStatus) Descriptor() ([]byte, []int) {
	return file_bff_matchmaking_proto_rawDescGZIP(), []int{5}
}

type S2CMatchingStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayId int32 `protobuf:"varint,1,opt,name=play_id,json=playId,proto3" json:"play_id,omitempty"` //玩法id
	Status int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`               //状态 0=空闲 1=准备状态 2=匹配状态
}

func (x *S2CMatchingStatus) Reset() {
	*x = S2CMatchingStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bff_matchmaking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CMatchingStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CMatchingStatus) ProtoMessage() {}

func (x *S2CMatchingStatus) ProtoReflect() protoreflect.Message {
	mi := &file_bff_matchmaking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CMatchingStatus.ProtoReflect.Descriptor instead.
func (*S2CMatchingStatus) Descriptor() ([]byte, []int) {
	return file_bff_matchmaking_proto_rawDescGZIP(), []int{6}
}

func (x *S2CMatchingStatus) GetPlayId() int32 {
	if x != nil {
		return x.PlayId
	}
	return 0
}

func (x *S2CMatchingStatus) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_bff_matchmaking_proto protoreflect.FileDescriptor

var file_bff_matchmaking_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x66, 0x66, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x62, 0x66, 0x66, 0x22, 0x30, 0x0a, 0x15,
	0x4e, 0x74, 0x66, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x22, 0x32,
	0x0a, 0x16, 0x43, 0x32, 0x53, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79,
	0x4d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x4d,
	0x61, 0x70, 0x22, 0x32, 0x0a, 0x16, 0x53, 0x32, 0x43, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x6c, 0x61, 0x79, 0x4d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x79, 0x4d, 0x61, 0x70, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x32, 0x53, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x22, 0x13, 0x0a, 0x11, 0x53,
	0x32, 0x43, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x22, 0x13, 0x0a, 0x11, 0x43, 0x32, 0x53, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x44, 0x0a, 0x11, 0x53, 0x32, 0x43, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6c, 0x61,
	0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x62,
	0x66, 0x66, 0x2f, 0x3b, 0x62, 0x66, 0x66, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_bff_matchmaking_proto_rawDescOnce sync.Once
	file_bff_matchmaking_proto_rawDescData = file_bff_matchmaking_proto_rawDesc
)

func file_bff_matchmaking_proto_rawDescGZIP() []byte {
	file_bff_matchmaking_proto_rawDescOnce.Do(func() {
		file_bff_matchmaking_proto_rawDescData = protoimpl.X.CompressGZIP(file_bff_matchmaking_proto_rawDescData)
	})
	return file_bff_matchmaking_proto_rawDescData
}

var file_bff_matchmaking_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bff_matchmaking_proto_goTypes = []any{
	(*NtfMatchingGroupStart)(nil),  // 0: bff.NtfMatchingGroupStart
	(*C2SMatchingSingleStart)(nil), // 1: bff.C2SMatchingSingleStart
	(*S2CMatchingSingleStart)(nil), // 2: bff.S2CMatchingSingleStart
	(*C2SMatchingCancel)(nil),      // 3: bff.C2SMatchingCancel
	(*S2CMatchingCancel)(nil),      // 4: bff.S2CMatchingCancel
	(*C2SMatchingStatus)(nil),      // 5: bff.C2SMatchingStatus
	(*S2CMatchingStatus)(nil),      // 6: bff.S2CMatchingStatus
}
var file_bff_matchmaking_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bff_matchmaking_proto_init() }
func file_bff_matchmaking_proto_init() {
	if File_bff_matchmaking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bff_matchmaking_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NtfMatchingGroupStart); i {
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
		file_bff_matchmaking_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*C2SMatchingSingleStart); i {
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
		file_bff_matchmaking_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*S2CMatchingSingleStart); i {
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
		file_bff_matchmaking_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*C2SMatchingCancel); i {
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
		file_bff_matchmaking_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*S2CMatchingCancel); i {
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
		file_bff_matchmaking_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*C2SMatchingStatus); i {
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
		file_bff_matchmaking_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*S2CMatchingStatus); i {
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
			RawDescriptor: file_bff_matchmaking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bff_matchmaking_proto_goTypes,
		DependencyIndexes: file_bff_matchmaking_proto_depIdxs,
		MessageInfos:      file_bff_matchmaking_proto_msgTypes,
	}.Build()
	File_bff_matchmaking_proto = out.File
	file_bff_matchmaking_proto_rawDesc = nil
	file_bff_matchmaking_proto_goTypes = nil
	file_bff_matchmaking_proto_depIdxs = nil
}
