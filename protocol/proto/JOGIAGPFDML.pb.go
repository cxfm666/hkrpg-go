// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: JOGIAGPFDML.proto

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

type JOGIAGPFDML struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DMMNJBLPDJO []*PCFEIBPMOKL `protobuf:"bytes,8,rep,name=DMMNJBLPDJO,proto3" json:"DMMNJBLPDJO,omitempty"`
	MACFPPFFFEJ *PCFEIBPMOKL   `protobuf:"bytes,2,opt,name=MACFPPFFFEJ,proto3" json:"MACFPPFFFEJ,omitempty"`
	DABBJIMJPNF *NIOJAOAKEJP   `protobuf:"bytes,7,opt,name=DABBJIMJPNF,proto3" json:"DABBJIMJPNF,omitempty"`
}

func (x *JOGIAGPFDML) Reset() {
	*x = JOGIAGPFDML{}
	if protoimpl.UnsafeEnabled {
		mi := &file_JOGIAGPFDML_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JOGIAGPFDML) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JOGIAGPFDML) ProtoMessage() {}

func (x *JOGIAGPFDML) ProtoReflect() protoreflect.Message {
	mi := &file_JOGIAGPFDML_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JOGIAGPFDML.ProtoReflect.Descriptor instead.
func (*JOGIAGPFDML) Descriptor() ([]byte, []int) {
	return file_JOGIAGPFDML_proto_rawDescGZIP(), []int{0}
}

func (x *JOGIAGPFDML) GetDMMNJBLPDJO() []*PCFEIBPMOKL {
	if x != nil {
		return x.DMMNJBLPDJO
	}
	return nil
}

func (x *JOGIAGPFDML) GetMACFPPFFFEJ() *PCFEIBPMOKL {
	if x != nil {
		return x.MACFPPFFFEJ
	}
	return nil
}

func (x *JOGIAGPFDML) GetDABBJIMJPNF() *NIOJAOAKEJP {
	if x != nil {
		return x.DABBJIMJPNF
	}
	return nil
}

var File_JOGIAGPFDML_proto protoreflect.FileDescriptor

var file_JOGIAGPFDML_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4a, 0x4f, 0x47, 0x49, 0x41, 0x47, 0x50, 0x46, 0x44, 0x4d, 0x4c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x43, 0x46, 0x45, 0x49, 0x42, 0x50, 0x4d, 0x4f, 0x4b, 0x4c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x49, 0x4f, 0x4a, 0x41, 0x4f, 0x41, 0x4b,
	0x45, 0x4a, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x0b, 0x4a, 0x4f,
	0x47, 0x49, 0x41, 0x47, 0x50, 0x46, 0x44, 0x4d, 0x4c, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x4d, 0x4d,
	0x4e, 0x4a, 0x42, 0x4c, 0x50, 0x44, 0x4a, 0x4f, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x50, 0x43, 0x46, 0x45, 0x49, 0x42, 0x50, 0x4d, 0x4f, 0x4b, 0x4c, 0x52, 0x0b, 0x44, 0x4d,
	0x4d, 0x4e, 0x4a, 0x42, 0x4c, 0x50, 0x44, 0x4a, 0x4f, 0x12, 0x2e, 0x0a, 0x0b, 0x4d, 0x41, 0x43,
	0x46, 0x50, 0x50, 0x46, 0x46, 0x46, 0x45, 0x4a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x50, 0x43, 0x46, 0x45, 0x49, 0x42, 0x50, 0x4d, 0x4f, 0x4b, 0x4c, 0x52, 0x0b, 0x4d, 0x41,
	0x43, 0x46, 0x50, 0x50, 0x46, 0x46, 0x46, 0x45, 0x4a, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x41, 0x42,
	0x42, 0x4a, 0x49, 0x4d, 0x4a, 0x50, 0x4e, 0x46, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4e, 0x49, 0x4f, 0x4a, 0x41, 0x4f, 0x41, 0x4b, 0x45, 0x4a, 0x50, 0x52, 0x0b, 0x44, 0x41,
	0x42, 0x42, 0x4a, 0x49, 0x4d, 0x4a, 0x50, 0x4e, 0x46, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_JOGIAGPFDML_proto_rawDescOnce sync.Once
	file_JOGIAGPFDML_proto_rawDescData = file_JOGIAGPFDML_proto_rawDesc
)

func file_JOGIAGPFDML_proto_rawDescGZIP() []byte {
	file_JOGIAGPFDML_proto_rawDescOnce.Do(func() {
		file_JOGIAGPFDML_proto_rawDescData = protoimpl.X.CompressGZIP(file_JOGIAGPFDML_proto_rawDescData)
	})
	return file_JOGIAGPFDML_proto_rawDescData
}

var file_JOGIAGPFDML_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_JOGIAGPFDML_proto_goTypes = []interface{}{
	(*JOGIAGPFDML)(nil), // 0: JOGIAGPFDML
	(*PCFEIBPMOKL)(nil), // 1: PCFEIBPMOKL
	(*NIOJAOAKEJP)(nil), // 2: NIOJAOAKEJP
}
var file_JOGIAGPFDML_proto_depIdxs = []int32{
	1, // 0: JOGIAGPFDML.DMMNJBLPDJO:type_name -> PCFEIBPMOKL
	1, // 1: JOGIAGPFDML.MACFPPFFFEJ:type_name -> PCFEIBPMOKL
	2, // 2: JOGIAGPFDML.DABBJIMJPNF:type_name -> NIOJAOAKEJP
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_JOGIAGPFDML_proto_init() }
func file_JOGIAGPFDML_proto_init() {
	if File_JOGIAGPFDML_proto != nil {
		return
	}
	file_PCFEIBPMOKL_proto_init()
	file_NIOJAOAKEJP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_JOGIAGPFDML_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JOGIAGPFDML); i {
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
			RawDescriptor: file_JOGIAGPFDML_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JOGIAGPFDML_proto_goTypes,
		DependencyIndexes: file_JOGIAGPFDML_proto_depIdxs,
		MessageInfos:      file_JOGIAGPFDML_proto_msgTypes,
	}.Build()
	File_JOGIAGPFDML_proto = out.File
	file_JOGIAGPFDML_proto_rawDesc = nil
	file_JOGIAGPFDML_proto_goTypes = nil
	file_JOGIAGPFDML_proto_depIdxs = nil
}
