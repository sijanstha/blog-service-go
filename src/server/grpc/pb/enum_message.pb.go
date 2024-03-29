// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: enum_message.proto

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

type Sort int32

const (
	Sort_ASC  Sort = 0
	Sort_DESC Sort = 1
)

// Enum value maps for Sort.
var (
	Sort_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	Sort_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x Sort) Enum() *Sort {
	p := new(Sort)
	*p = x
	return p
}

func (x Sort) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sort) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_message_proto_enumTypes[0].Descriptor()
}

func (Sort) Type() protoreflect.EnumType {
	return &file_enum_message_proto_enumTypes[0]
}

func (x Sort) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sort.Descriptor instead.
func (Sort) EnumDescriptor() ([]byte, []int) {
	return file_enum_message_proto_rawDescGZIP(), []int{0}
}

type SortBy int32

const (
	SortBy_title       SortBy = 0
	SortBy_created_at  SortBy = 1
	SortBy_updated_at  SortBy = 2
	SortBy_description SortBy = 3
	SortBy_id          SortBy = 4
)

// Enum value maps for SortBy.
var (
	SortBy_name = map[int32]string{
		0: "title",
		1: "created_at",
		2: "updated_at",
		3: "description",
		4: "id",
	}
	SortBy_value = map[string]int32{
		"title":       0,
		"created_at":  1,
		"updated_at":  2,
		"description": 3,
		"id":          4,
	}
)

func (x SortBy) Enum() *SortBy {
	p := new(SortBy)
	*p = x
	return p
}

func (x SortBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortBy) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_message_proto_enumTypes[1].Descriptor()
}

func (SortBy) Type() protoreflect.EnumType {
	return &file_enum_message_proto_enumTypes[1]
}

func (x SortBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortBy.Descriptor instead.
func (SortBy) EnumDescriptor() ([]byte, []int) {
	return file_enum_message_proto_rawDescGZIP(), []int{1}
}

var File_enum_message_proto protoreflect.FileDescriptor

var file_enum_message_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x2a, 0x19, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53,
	0x43, 0x10, 0x01, 0x2a, 0x4c, 0x0a, 0x06, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x09, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10,
	0x04, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_enum_message_proto_rawDescOnce sync.Once
	file_enum_message_proto_rawDescData = file_enum_message_proto_rawDesc
)

func file_enum_message_proto_rawDescGZIP() []byte {
	file_enum_message_proto_rawDescOnce.Do(func() {
		file_enum_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_enum_message_proto_rawDescData)
	})
	return file_enum_message_proto_rawDescData
}

var file_enum_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_enum_message_proto_goTypes = []interface{}{
	(Sort)(0),   // 0: pb.Sort
	(SortBy)(0), // 1: pb.SortBy
}
var file_enum_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enum_message_proto_init() }
func file_enum_message_proto_init() {
	if File_enum_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enum_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enum_message_proto_goTypes,
		DependencyIndexes: file_enum_message_proto_depIdxs,
		EnumInfos:         file_enum_message_proto_enumTypes,
	}.Build()
	File_enum_message_proto = out.File
	file_enum_message_proto_rawDesc = nil
	file_enum_message_proto_goTypes = nil
	file_enum_message_proto_depIdxs = nil
}
