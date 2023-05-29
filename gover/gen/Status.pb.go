// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Status.proto

package gen

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

type Status int32

const (
	Status_STATUS_INVALID      Status = 0
	Status_STATUS_UNFINISHED   Status = 1
	Status_STATUS_FINISHED     Status = 2
	Status_STATUS_REWARD_TAKEN Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_INVALID",
		1: "STATUS_UNFINISHED",
		2: "STATUS_FINISHED",
		3: "STATUS_REWARD_TAKEN",
	}
	Status_value = map[string]int32{
		"STATUS_INVALID":      0,
		"STATUS_UNFINISHED":   1,
		"STATUS_FINISHED":     2,
		"STATUS_REWARD_TAKEN": 3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_Status_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_Status_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_Status_proto_rawDescGZIP(), []int{0}
}

var File_Status_proto protoreflect.FileDescriptor

var file_Status_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x61,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49,
	0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x54, 0x41, 0x4b, 0x45, 0x4e, 0x10,
	0x03, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Status_proto_rawDescOnce sync.Once
	file_Status_proto_rawDescData = file_Status_proto_rawDesc
)

func file_Status_proto_rawDescGZIP() []byte {
	file_Status_proto_rawDescOnce.Do(func() {
		file_Status_proto_rawDescData = protoimpl.X.CompressGZIP(file_Status_proto_rawDescData)
	})
	return file_Status_proto_rawDescData
}

var file_Status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Status_proto_goTypes = []interface{}{
	(Status)(0), // 0: Status
}
var file_Status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Status_proto_init() }
func file_Status_proto_init() {
	if File_Status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Status_proto_goTypes,
		DependencyIndexes: file_Status_proto_depIdxs,
		EnumInfos:         file_Status_proto_enumTypes,
	}.Build()
	File_Status_proto = out.File
	file_Status_proto_rawDesc = nil
	file_Status_proto_goTypes = nil
	file_Status_proto_depIdxs = nil
}
