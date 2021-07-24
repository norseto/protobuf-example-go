// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: enum_example.proto

package enumpb

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

type DayOfTheWeek int32

const (
	DayOfTheWeek_UNKNOWN   DayOfTheWeek = 0
	DayOfTheWeek_MONDAY    DayOfTheWeek = 1
	DayOfTheWeek_TUESDAY   DayOfTheWeek = 2
	DayOfTheWeek_WEDNESDAY DayOfTheWeek = 3
	DayOfTheWeek_THURSDAY  DayOfTheWeek = 4
	DayOfTheWeek_FRIDAY    DayOfTheWeek = 5
	DayOfTheWeek_SATURDAY  DayOfTheWeek = 6
	DayOfTheWeek_SUNDAY    DayOfTheWeek = 7
)

// Enum value maps for DayOfTheWeek.
var (
	DayOfTheWeek_name = map[int32]string{
		0: "UNKNOWN",
		1: "MONDAY",
		2: "TUESDAY",
		3: "WEDNESDAY",
		4: "THURSDAY",
		5: "FRIDAY",
		6: "SATURDAY",
		7: "SUNDAY",
	}
	DayOfTheWeek_value = map[string]int32{
		"UNKNOWN":   0,
		"MONDAY":    1,
		"TUESDAY":   2,
		"WEDNESDAY": 3,
		"THURSDAY":  4,
		"FRIDAY":    5,
		"SATURDAY":  6,
		"SUNDAY":    7,
	}
)

func (x DayOfTheWeek) Enum() *DayOfTheWeek {
	p := new(DayOfTheWeek)
	*p = x
	return p
}

func (x DayOfTheWeek) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DayOfTheWeek) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_example_proto_enumTypes[0].Descriptor()
}

func (DayOfTheWeek) Type() protoreflect.EnumType {
	return &file_enum_example_proto_enumTypes[0]
}

func (x DayOfTheWeek) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DayOfTheWeek.Descriptor instead.
func (DayOfTheWeek) EnumDescriptor() ([]byte, []int) {
	return file_enum_example_proto_rawDescGZIP(), []int{0}
}

type EnumMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DayOfTheWeek DayOfTheWeek `protobuf:"varint,2,opt,name=day_of_the_week,json=dayOfTheWeek,proto3,enum=example.enumerations.DayOfTheWeek" json:"day_of_the_week,omitempty"`
}

func (x *EnumMessage) Reset() {
	*x = EnumMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enum_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumMessage) ProtoMessage() {}

func (x *EnumMessage) ProtoReflect() protoreflect.Message {
	mi := &file_enum_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumMessage.ProtoReflect.Descriptor instead.
func (*EnumMessage) Descriptor() ([]byte, []int) {
	return file_enum_example_proto_rawDescGZIP(), []int{0}
}

func (x *EnumMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EnumMessage) GetDayOfTheWeek() DayOfTheWeek {
	if x != nil {
		return x.DayOfTheWeek
	}
	return DayOfTheWeek_UNKNOWN
}

var File_enum_example_proto protoreflect.FileDescriptor

var file_enum_example_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x68, 0x0a, 0x0b, 0x45, 0x6e,
	0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x49, 0x0a, 0x0f, 0x64, 0x61, 0x79,
	0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x68, 0x65, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x22, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x54,
	0x68, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x0c, 0x64, 0x61, 0x79, 0x4f, 0x66, 0x54, 0x68, 0x65,
	0x57, 0x65, 0x65, 0x6b, 0x2a, 0x77, 0x0a, 0x0c, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x54, 0x68, 0x65,
	0x57, 0x65, 0x65, 0x6b, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x54, 0x55, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x45,
	0x44, 0x4e, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x48, 0x55,
	0x52, 0x53, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x49, 0x44, 0x41,
	0x59, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x41, 0x54, 0x55, 0x52, 0x44, 0x41, 0x59, 0x10,
	0x06, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x07, 0x42, 0x11, 0x5a,
	0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enum_example_proto_rawDescOnce sync.Once
	file_enum_example_proto_rawDescData = file_enum_example_proto_rawDesc
)

func file_enum_example_proto_rawDescGZIP() []byte {
	file_enum_example_proto_rawDescOnce.Do(func() {
		file_enum_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_enum_example_proto_rawDescData)
	})
	return file_enum_example_proto_rawDescData
}

var file_enum_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enum_example_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enum_example_proto_goTypes = []interface{}{
	(DayOfTheWeek)(0),   // 0: example.enumerations.DayOfTheWeek
	(*EnumMessage)(nil), // 1: example.enumerations.EnumMessage
}
var file_enum_example_proto_depIdxs = []int32{
	0, // 0: example.enumerations.EnumMessage.day_of_the_week:type_name -> example.enumerations.DayOfTheWeek
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_enum_example_proto_init() }
func file_enum_example_proto_init() {
	if File_enum_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enum_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumMessage); i {
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
			RawDescriptor: file_enum_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enum_example_proto_goTypes,
		DependencyIndexes: file_enum_example_proto_depIdxs,
		EnumInfos:         file_enum_example_proto_enumTypes,
		MessageInfos:      file_enum_example_proto_msgTypes,
	}.Build()
	File_enum_example_proto = out.File
	file_enum_example_proto_rawDesc = nil
	file_enum_example_proto_goTypes = nil
	file_enum_example_proto_depIdxs = nil
}
