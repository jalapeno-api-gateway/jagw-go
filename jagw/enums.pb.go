// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: core/enums.proto

package jagw

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

type StringOperator int32

const (
	StringOperator_EQUAL     StringOperator = 1
	StringOperator_NOT_EQUAL StringOperator = 2
)

// Enum value maps for StringOperator.
var (
	StringOperator_name = map[int32]string{
		1: "EQUAL",
		2: "NOT_EQUAL",
	}
	StringOperator_value = map[string]int32{
		"EQUAL":     1,
		"NOT_EQUAL": 2,
	}
)

func (x StringOperator) Enum() *StringOperator {
	p := new(StringOperator)
	*p = x
	return p
}

func (x StringOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StringOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_core_enums_proto_enumTypes[0].Descriptor()
}

func (StringOperator) Type() protoreflect.EnumType {
	return &file_core_enums_proto_enumTypes[0]
}

func (x StringOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *StringOperator) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = StringOperator(num)
	return nil
}

// Deprecated: Use StringOperator.Descriptor instead.
func (StringOperator) EnumDescriptor() ([]byte, []int) {
	return file_core_enums_proto_rawDescGZIP(), []int{0}
}

var File_core_enums_proto protoreflect.FileDescriptor

var file_core_enums_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6a, 0x61, 0x67, 0x77, 0x2a, 0x2a, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x51,
	0x55, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55,
	0x41, 0x4c, 0x10, 0x02, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6c, 0x61, 0x70, 0x65, 0x6e, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6a, 0x61, 0x67, 0x77, 0x2d, 0x67, 0x6f, 0x3b,
	0x6a, 0x61, 0x67, 0x77,
}

var (
	file_core_enums_proto_rawDescOnce sync.Once
	file_core_enums_proto_rawDescData = file_core_enums_proto_rawDesc
)

func file_core_enums_proto_rawDescGZIP() []byte {
	file_core_enums_proto_rawDescOnce.Do(func() {
		file_core_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_enums_proto_rawDescData)
	})
	return file_core_enums_proto_rawDescData
}

var file_core_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_enums_proto_goTypes = []interface{}{
	(StringOperator)(0), // 0: jagw.StringOperator
}
var file_core_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_enums_proto_init() }
func file_core_enums_proto_init() {
	if File_core_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_enums_proto_goTypes,
		DependencyIndexes: file_core_enums_proto_depIdxs,
		EnumInfos:         file_core_enums_proto_enumTypes,
	}.Build()
	File_core_enums_proto = out.File
	file_core_enums_proto_rawDesc = nil
	file_core_enums_proto_goTypes = nil
	file_core_enums_proto_depIdxs = nil
}
