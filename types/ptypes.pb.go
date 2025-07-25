// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: types/ptypes.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetInvoiceRequets struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Obuid         int64                  `protobuf:"varint,1,opt,name=obuid,proto3" json:"obuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInvoiceRequets) Reset() {
	*x = GetInvoiceRequets{}
	mi := &file_types_ptypes_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInvoiceRequets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoiceRequets) ProtoMessage() {}

func (x *GetInvoiceRequets) ProtoReflect() protoreflect.Message {
	mi := &file_types_ptypes_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoiceRequets.ProtoReflect.Descriptor instead.
func (*GetInvoiceRequets) Descriptor() ([]byte, []int) {
	return file_types_ptypes_proto_rawDescGZIP(), []int{0}
}

func (x *GetInvoiceRequets) GetObuid() int64 {
	if x != nil {
		return x.Obuid
	}
	return 0
}

type None struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *None) Reset() {
	*x = None{}
	mi := &file_types_ptypes_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *None) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*None) ProtoMessage() {}

func (x *None) ProtoReflect() protoreflect.Message {
	mi := &file_types_ptypes_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use None.ProtoReflect.Descriptor instead.
func (*None) Descriptor() ([]byte, []int) {
	return file_types_ptypes_proto_rawDescGZIP(), []int{1}
}

type AggregatorDistance struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Obuid         int64                  `protobuf:"varint,1,opt,name=obuid,proto3" json:"obuid,omitempty"`
	UnixTimestamp int64                  `protobuf:"varint,2,opt,name=unix_timestamp,json=unixTimestamp,proto3" json:"unix_timestamp,omitempty"`
	Value         float64                `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AggregatorDistance) Reset() {
	*x = AggregatorDistance{}
	mi := &file_types_ptypes_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AggregatorDistance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregatorDistance) ProtoMessage() {}

func (x *AggregatorDistance) ProtoReflect() protoreflect.Message {
	mi := &file_types_ptypes_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregatorDistance.ProtoReflect.Descriptor instead.
func (*AggregatorDistance) Descriptor() ([]byte, []int) {
	return file_types_ptypes_proto_rawDescGZIP(), []int{2}
}

func (x *AggregatorDistance) GetObuid() int64 {
	if x != nil {
		return x.Obuid
	}
	return 0
}

func (x *AggregatorDistance) GetUnixTimestamp() int64 {
	if x != nil {
		return x.UnixTimestamp
	}
	return 0
}

func (x *AggregatorDistance) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_types_ptypes_proto protoreflect.FileDescriptor

const file_types_ptypes_proto_rawDesc = "" +
	"\n" +
	"\x12types/ptypes.proto\")\n" +
	"\x11GetInvoiceRequets\x12\x14\n" +
	"\x05obuid\x18\x01 \x01(\x03R\x05obuid\"\x06\n" +
	"\x04None\"g\n" +
	"\x12AggregatorDistance\x12\x14\n" +
	"\x05obuid\x18\x01 \x01(\x03R\x05obuid\x12%\n" +
	"\x0eunix_timestamp\x18\x02 \x01(\x03R\runixTimestamp\x12\x14\n" +
	"\x05value\x18\x03 \x01(\x01R\x05value25\n" +
	"\n" +
	"Aggregator\x12'\n" +
	"\tAggregate\x12\x13.AggregatorDistance\x1a\x05.NoneB)Z'github.com/pdrm26/toll-calculator/typesb\x06proto3"

var (
	file_types_ptypes_proto_rawDescOnce sync.Once
	file_types_ptypes_proto_rawDescData []byte
)

func file_types_ptypes_proto_rawDescGZIP() []byte {
	file_types_ptypes_proto_rawDescOnce.Do(func() {
		file_types_ptypes_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_types_ptypes_proto_rawDesc), len(file_types_ptypes_proto_rawDesc)))
	})
	return file_types_ptypes_proto_rawDescData
}

var file_types_ptypes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_types_ptypes_proto_goTypes = []any{
	(*GetInvoiceRequets)(nil),  // 0: GetInvoiceRequets
	(*None)(nil),               // 1: None
	(*AggregatorDistance)(nil), // 2: AggregatorDistance
}
var file_types_ptypes_proto_depIdxs = []int32{
	2, // 0: Aggregator.Aggregate:input_type -> AggregatorDistance
	1, // 1: Aggregator.Aggregate:output_type -> None
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_ptypes_proto_init() }
func file_types_ptypes_proto_init() {
	if File_types_ptypes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_types_ptypes_proto_rawDesc), len(file_types_ptypes_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_types_ptypes_proto_goTypes,
		DependencyIndexes: file_types_ptypes_proto_depIdxs,
		MessageInfos:      file_types_ptypes_proto_msgTypes,
	}.Build()
	File_types_ptypes_proto = out.File
	file_types_ptypes_proto_goTypes = nil
	file_types_ptypes_proto_depIdxs = nil
}
