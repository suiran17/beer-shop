// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: v1/courier_error.proto

package v1

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

type CourierJobErrorReason int32

const (
	CourierJobErrorReason_UNKNOWN_ERROR CourierJobErrorReason = 0
)

// Enum value maps for CourierJobErrorReason.
var (
	CourierJobErrorReason_name = map[int32]string{
		0: "UNKNOWN_ERROR",
	}
	CourierJobErrorReason_value = map[string]int32{
		"UNKNOWN_ERROR": 0,
	}
)

func (x CourierJobErrorReason) Enum() *CourierJobErrorReason {
	p := new(CourierJobErrorReason)
	*p = x
	return p
}

func (x CourierJobErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CourierJobErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_courier_error_proto_enumTypes[0].Descriptor()
}

func (CourierJobErrorReason) Type() protoreflect.EnumType {
	return &file_v1_courier_error_proto_enumTypes[0]
}

func (x CourierJobErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CourierJobErrorReason.Descriptor instead.
func (CourierJobErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_v1_courier_error_proto_rawDescGZIP(), []int{0}
}

var File_v1_courier_error_proto protoreflect.FileDescriptor

var file_v1_courier_error_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2a,
	0x2a, 0x0a, 0x15, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x42, 0x18, 0x50, 0x01, 0x5a,
	0x14, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_courier_error_proto_rawDescOnce sync.Once
	file_v1_courier_error_proto_rawDescData = file_v1_courier_error_proto_rawDesc
)

func file_v1_courier_error_proto_rawDescGZIP() []byte {
	file_v1_courier_error_proto_rawDescOnce.Do(func() {
		file_v1_courier_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_courier_error_proto_rawDescData)
	})
	return file_v1_courier_error_proto_rawDescData
}

var file_v1_courier_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_courier_error_proto_goTypes = []interface{}{
	(CourierJobErrorReason)(0), // 0: shop.interface.errors.CourierJobErrorReason
}
var file_v1_courier_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_courier_error_proto_init() }
func file_v1_courier_error_proto_init() {
	if File_v1_courier_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_courier_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_courier_error_proto_goTypes,
		DependencyIndexes: file_v1_courier_error_proto_depIdxs,
		EnumInfos:         file_v1_courier_error_proto_enumTypes,
	}.Build()
	File_v1_courier_error_proto = out.File
	file_v1_courier_error_proto_rawDesc = nil
	file_v1_courier_error_proto_goTypes = nil
	file_v1_courier_error_proto_depIdxs = nil
}
