// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: StartWolfBroGameCsReq.proto

package proto

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

type StartWolfBroGameCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupStateInfo *GroupStateInfo `protobuf:"bytes,10,opt,name=group_state_info,json=groupStateInfo,proto3" json:"group_state_info,omitempty"`
	Motion         *MotionInfo     `protobuf:"bytes,4,opt,name=motion,proto3" json:"motion,omitempty"`
	PECAGANEKGI    bool            `protobuf:"varint,3,opt,name=PECAGANEKGI,proto3" json:"PECAGANEKGI,omitempty"`
	Id             uint32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartWolfBroGameCsReq) Reset() {
	*x = StartWolfBroGameCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StartWolfBroGameCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartWolfBroGameCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartWolfBroGameCsReq) ProtoMessage() {}

func (x *StartWolfBroGameCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_StartWolfBroGameCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartWolfBroGameCsReq.ProtoReflect.Descriptor instead.
func (*StartWolfBroGameCsReq) Descriptor() ([]byte, []int) {
	return file_StartWolfBroGameCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *StartWolfBroGameCsReq) GetGroupStateInfo() *GroupStateInfo {
	if x != nil {
		return x.GroupStateInfo
	}
	return nil
}

func (x *StartWolfBroGameCsReq) GetMotion() *MotionInfo {
	if x != nil {
		return x.Motion
	}
	return nil
}

func (x *StartWolfBroGameCsReq) GetPECAGANEKGI() bool {
	if x != nil {
		return x.PECAGANEKGI
	}
	return false
}

func (x *StartWolfBroGameCsReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_StartWolfBroGameCsReq_proto protoreflect.FileDescriptor

var file_StartWolfBroGameCsReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x57, 0x6f, 0x6c, 0x66, 0x42, 0x72, 0x6f, 0x47, 0x61,
	0x6d, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x72, 0x74, 0x57,
	0x6f, 0x6c, 0x66, 0x42, 0x72, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x39, 0x0a, 0x10, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x06, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x50, 0x45, 0x43, 0x41, 0x47, 0x41, 0x4e, 0x45, 0x4b, 0x47, 0x49, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x50, 0x45, 0x43, 0x41, 0x47, 0x41, 0x4e, 0x45, 0x4b, 0x47,
	0x49, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_StartWolfBroGameCsReq_proto_rawDescOnce sync.Once
	file_StartWolfBroGameCsReq_proto_rawDescData = file_StartWolfBroGameCsReq_proto_rawDesc
)

func file_StartWolfBroGameCsReq_proto_rawDescGZIP() []byte {
	file_StartWolfBroGameCsReq_proto_rawDescOnce.Do(func() {
		file_StartWolfBroGameCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_StartWolfBroGameCsReq_proto_rawDescData)
	})
	return file_StartWolfBroGameCsReq_proto_rawDescData
}

var file_StartWolfBroGameCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StartWolfBroGameCsReq_proto_goTypes = []interface{}{
	(*StartWolfBroGameCsReq)(nil), // 0: StartWolfBroGameCsReq
	(*GroupStateInfo)(nil),        // 1: GroupStateInfo
	(*MotionInfo)(nil),            // 2: MotionInfo
}
var file_StartWolfBroGameCsReq_proto_depIdxs = []int32{
	1, // 0: StartWolfBroGameCsReq.group_state_info:type_name -> GroupStateInfo
	2, // 1: StartWolfBroGameCsReq.motion:type_name -> MotionInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_StartWolfBroGameCsReq_proto_init() }
func file_StartWolfBroGameCsReq_proto_init() {
	if File_StartWolfBroGameCsReq_proto != nil {
		return
	}
	file_MotionInfo_proto_init()
	file_GroupStateInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_StartWolfBroGameCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartWolfBroGameCsReq); i {
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
			RawDescriptor: file_StartWolfBroGameCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StartWolfBroGameCsReq_proto_goTypes,
		DependencyIndexes: file_StartWolfBroGameCsReq_proto_depIdxs,
		MessageInfos:      file_StartWolfBroGameCsReq_proto_msgTypes,
	}.Build()
	File_StartWolfBroGameCsReq_proto = out.File
	file_StartWolfBroGameCsReq_proto_rawDesc = nil
	file_StartWolfBroGameCsReq_proto_goTypes = nil
	file_StartWolfBroGameCsReq_proto_depIdxs = nil
}
