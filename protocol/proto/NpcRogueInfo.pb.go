// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: NpcRogueInfo.proto

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

type NpcRogueInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HCGOEHLPCMD map[uint32]uint32 `protobuf:"bytes,15,rep,name=HCGOEHLPCMD,proto3" json:"HCGOEHLPCMD,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	AGBFDIIKPKF bool              `protobuf:"varint,13,opt,name=AGBFDIIKPKF,proto3" json:"AGBFDIIKPKF,omitempty"`
	BHPGGEOBJHM bool              `protobuf:"varint,5,opt,name=BHPGGEOBJHM,proto3" json:"BHPGGEOBJHM,omitempty"`
	RogueNpcId  uint32            `protobuf:"varint,6,opt,name=rogue_npc_id,json=rogueNpcId,proto3" json:"rogue_npc_id,omitempty"`
	GBMDBBBMBEJ uint32            `protobuf:"varint,14,opt,name=GBMDBBBMBEJ,proto3" json:"GBMDBBBMBEJ,omitempty"`
	INJPFALMDHJ uint32            `protobuf:"varint,12,opt,name=INJPFALMDHJ,proto3" json:"INJPFALMDHJ,omitempty"`
	// repeated KELFICDMDDG CEEFGPGODAA = 9;
	GENMIFOCMJA uint32 `protobuf:"varint,1,opt,name=GENMIFOCMJA,proto3" json:"GENMIFOCMJA,omitempty"`
	MNINDBMAJKL bool   `protobuf:"varint,7,opt,name=MNINDBMAJKL,proto3" json:"MNINDBMAJKL,omitempty"`
}

func (x *NpcRogueInfo) Reset() {
	*x = NpcRogueInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NpcRogueInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NpcRogueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NpcRogueInfo) ProtoMessage() {}

func (x *NpcRogueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_NpcRogueInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NpcRogueInfo.ProtoReflect.Descriptor instead.
func (*NpcRogueInfo) Descriptor() ([]byte, []int) {
	return file_NpcRogueInfo_proto_rawDescGZIP(), []int{0}
}

func (x *NpcRogueInfo) GetHCGOEHLPCMD() map[uint32]uint32 {
	if x != nil {
		return x.HCGOEHLPCMD
	}
	return nil
}

func (x *NpcRogueInfo) GetAGBFDIIKPKF() bool {
	if x != nil {
		return x.AGBFDIIKPKF
	}
	return false
}

func (x *NpcRogueInfo) GetBHPGGEOBJHM() bool {
	if x != nil {
		return x.BHPGGEOBJHM
	}
	return false
}

func (x *NpcRogueInfo) GetRogueNpcId() uint32 {
	if x != nil {
		return x.RogueNpcId
	}
	return 0
}

func (x *NpcRogueInfo) GetGBMDBBBMBEJ() uint32 {
	if x != nil {
		return x.GBMDBBBMBEJ
	}
	return 0
}

func (x *NpcRogueInfo) GetINJPFALMDHJ() uint32 {
	if x != nil {
		return x.INJPFALMDHJ
	}
	return 0
}

func (x *NpcRogueInfo) GetGENMIFOCMJA() uint32 {
	if x != nil {
		return x.GENMIFOCMJA
	}
	return 0
}

func (x *NpcRogueInfo) GetMNINDBMAJKL() bool {
	if x != nil {
		return x.MNINDBMAJKL
	}
	return false
}

var File_NpcRogueInfo_proto protoreflect.FileDescriptor

var file_NpcRogueInfo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x4e, 0x70, 0x63, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x0c, 0x4e, 0x70, 0x63, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x0b, 0x48, 0x43, 0x47, 0x4f, 0x45, 0x48, 0x4c,
	0x50, 0x43, 0x4d, 0x44, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x4e, 0x70, 0x63,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x48, 0x43, 0x47, 0x4f, 0x45, 0x48,
	0x4c, 0x50, 0x43, 0x4d, 0x44, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x48, 0x43, 0x47, 0x4f,
	0x45, 0x48, 0x4c, 0x50, 0x43, 0x4d, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x47, 0x42, 0x46, 0x44,
	0x49, 0x49, 0x4b, 0x50, 0x4b, 0x46, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x41, 0x47,
	0x42, 0x46, 0x44, 0x49, 0x49, 0x4b, 0x50, 0x4b, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x48, 0x50,
	0x47, 0x47, 0x45, 0x4f, 0x42, 0x4a, 0x48, 0x4d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x42, 0x48, 0x50, 0x47, 0x47, 0x45, 0x4f, 0x42, 0x4a, 0x48, 0x4d, 0x12, 0x20, 0x0a, 0x0c, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x6e, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x4e, 0x70, 0x63, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x47, 0x42, 0x4d, 0x44, 0x42, 0x42, 0x42, 0x4d, 0x42, 0x45, 0x4a, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x47, 0x42, 0x4d, 0x44, 0x42, 0x42, 0x42, 0x4d, 0x42, 0x45, 0x4a, 0x12,
	0x20, 0x0a, 0x0b, 0x49, 0x4e, 0x4a, 0x50, 0x46, 0x41, 0x4c, 0x4d, 0x44, 0x48, 0x4a, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x4e, 0x4a, 0x50, 0x46, 0x41, 0x4c, 0x4d, 0x44, 0x48,
	0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x45, 0x4e, 0x4d, 0x49, 0x46, 0x4f, 0x43, 0x4d, 0x4a, 0x41,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x45, 0x4e, 0x4d, 0x49, 0x46, 0x4f, 0x43,
	0x4d, 0x4a, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x4e, 0x49, 0x4e, 0x44, 0x42, 0x4d, 0x41, 0x4a,
	0x4b, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4d, 0x4e, 0x49, 0x4e, 0x44, 0x42,
	0x4d, 0x41, 0x4a, 0x4b, 0x4c, 0x1a, 0x3e, 0x0a, 0x10, 0x48, 0x43, 0x47, 0x4f, 0x45, 0x48, 0x4c,
	0x50, 0x43, 0x4d, 0x44, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NpcRogueInfo_proto_rawDescOnce sync.Once
	file_NpcRogueInfo_proto_rawDescData = file_NpcRogueInfo_proto_rawDesc
)

func file_NpcRogueInfo_proto_rawDescGZIP() []byte {
	file_NpcRogueInfo_proto_rawDescOnce.Do(func() {
		file_NpcRogueInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_NpcRogueInfo_proto_rawDescData)
	})
	return file_NpcRogueInfo_proto_rawDescData
}

var file_NpcRogueInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_NpcRogueInfo_proto_goTypes = []interface{}{
	(*NpcRogueInfo)(nil), // 0: NpcRogueInfo
	nil,                  // 1: NpcRogueInfo.HCGOEHLPCMDEntry
}
var file_NpcRogueInfo_proto_depIdxs = []int32{
	1, // 0: NpcRogueInfo.HCGOEHLPCMD:type_name -> NpcRogueInfo.HCGOEHLPCMDEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_NpcRogueInfo_proto_init() }
func file_NpcRogueInfo_proto_init() {
	if File_NpcRogueInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_NpcRogueInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NpcRogueInfo); i {
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
			RawDescriptor: file_NpcRogueInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NpcRogueInfo_proto_goTypes,
		DependencyIndexes: file_NpcRogueInfo_proto_depIdxs,
		MessageInfos:      file_NpcRogueInfo_proto_msgTypes,
	}.Build()
	File_NpcRogueInfo_proto = out.File
	file_NpcRogueInfo_proto_rawDesc = nil
	file_NpcRogueInfo_proto_goTypes = nil
	file_NpcRogueInfo_proto_depIdxs = nil
}
