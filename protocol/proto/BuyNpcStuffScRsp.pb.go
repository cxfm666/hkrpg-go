// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BuyNpcStuffScRsp.proto

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

type BuyNpcStuffScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     uint32 `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
	GBKGFOJIEAB uint32 `protobuf:"varint,8,opt,name=GBKGFOJIEAB,proto3" json:"GBKGFOJIEAB,omitempty"`
}

func (x *BuyNpcStuffScRsp) Reset() {
	*x = BuyNpcStuffScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BuyNpcStuffScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyNpcStuffScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyNpcStuffScRsp) ProtoMessage() {}

func (x *BuyNpcStuffScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_BuyNpcStuffScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyNpcStuffScRsp.ProtoReflect.Descriptor instead.
func (*BuyNpcStuffScRsp) Descriptor() ([]byte, []int) {
	return file_BuyNpcStuffScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *BuyNpcStuffScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *BuyNpcStuffScRsp) GetGBKGFOJIEAB() uint32 {
	if x != nil {
		return x.GBKGFOJIEAB
	}
	return 0
}

var File_BuyNpcStuffScRsp_proto protoreflect.FileDescriptor

var file_BuyNpcStuffScRsp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x42, 0x75, 0x79, 0x4e, 0x70, 0x63, 0x53, 0x74, 0x75, 0x66, 0x66, 0x53, 0x63, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x10, 0x42, 0x75, 0x79, 0x4e,
	0x70, 0x63, 0x53, 0x74, 0x75, 0x66, 0x66, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x42, 0x4b, 0x47, 0x46, 0x4f,
	0x4a, 0x49, 0x45, 0x41, 0x42, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x42, 0x4b,
	0x47, 0x46, 0x4f, 0x4a, 0x49, 0x45, 0x41, 0x42, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BuyNpcStuffScRsp_proto_rawDescOnce sync.Once
	file_BuyNpcStuffScRsp_proto_rawDescData = file_BuyNpcStuffScRsp_proto_rawDesc
)

func file_BuyNpcStuffScRsp_proto_rawDescGZIP() []byte {
	file_BuyNpcStuffScRsp_proto_rawDescOnce.Do(func() {
		file_BuyNpcStuffScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_BuyNpcStuffScRsp_proto_rawDescData)
	})
	return file_BuyNpcStuffScRsp_proto_rawDescData
}

var file_BuyNpcStuffScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BuyNpcStuffScRsp_proto_goTypes = []interface{}{
	(*BuyNpcStuffScRsp)(nil), // 0: BuyNpcStuffScRsp
}
var file_BuyNpcStuffScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BuyNpcStuffScRsp_proto_init() }
func file_BuyNpcStuffScRsp_proto_init() {
	if File_BuyNpcStuffScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BuyNpcStuffScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyNpcStuffScRsp); i {
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
			RawDescriptor: file_BuyNpcStuffScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BuyNpcStuffScRsp_proto_goTypes,
		DependencyIndexes: file_BuyNpcStuffScRsp_proto_depIdxs,
		MessageInfos:      file_BuyNpcStuffScRsp_proto_msgTypes,
	}.Build()
	File_BuyNpcStuffScRsp_proto = out.File
	file_BuyNpcStuffScRsp_proto_rawDesc = nil
	file_BuyNpcStuffScRsp_proto_goTypes = nil
	file_BuyNpcStuffScRsp_proto_depIdxs = nil
}
