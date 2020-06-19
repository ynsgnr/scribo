// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: device.proto

package device

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DeviceType int32

const (
	DeviceType_UNKNOWN_DEVICE DeviceType = 0
	DeviceType_KINDLE         DeviceType = 1
)

// Enum value maps for DeviceType.
var (
	DeviceType_name = map[int32]string{
		0: "UNKNOWN_DEVICE",
		1: "KINDLE",
	}
	DeviceType_value = map[string]int32{
		"UNKNOWN_DEVICE": 0,
		"KINDLE":         1,
	}
)

func (x DeviceType) Enum() *DeviceType {
	p := new(DeviceType)
	*p = x
	return p
}

func (x DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_device_proto_enumTypes[0].Descriptor()
}

func (DeviceType) Type() protoreflect.EnumType {
	return &file_device_proto_enumTypes[0]
}

func (x DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceType.Descriptor instead.
func (DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{0}
}

type AddDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceName string     `protobuf:"bytes,1,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	DeviceID   string     `protobuf:"bytes,2,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	DeviceType DeviceType `protobuf:"varint,3,opt,name=deviceType,proto3,enum=scribo.DeviceType" json:"deviceType,omitempty"`
	AddKindle  *AddKindle `protobuf:"bytes,4,opt,name=addKindle,proto3" json:"addKindle,omitempty"`
}

func (x *AddDevice) Reset() {
	*x = AddDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDevice) ProtoMessage() {}

func (x *AddDevice) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDevice.ProtoReflect.Descriptor instead.
func (*AddDevice) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{0}
}

func (x *AddDevice) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *AddDevice) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *AddDevice) GetDeviceType() DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return DeviceType_UNKNOWN_DEVICE
}

func (x *AddDevice) GetAddKindle() *AddKindle {
	if x != nil {
		return x.AddKindle
	}
	return nil
}

type Sync2Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SyncID       string `protobuf:"bytes,1,opt,name=syncID,proto3" json:"syncID,omitempty"`
	DeviceID     string `protobuf:"bytes,2,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	FileID       string `protobuf:"bytes,3,opt,name=fileID,proto3" json:"fileID,omitempty"`
	FileLocation string `protobuf:"bytes,4,opt,name=fileLocation,proto3" json:"fileLocation,omitempty"`
}

func (x *Sync2Device) Reset() {
	*x = Sync2Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sync2Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sync2Device) ProtoMessage() {}

func (x *Sync2Device) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sync2Device.ProtoReflect.Descriptor instead.
func (*Sync2Device) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{1}
}

func (x *Sync2Device) GetSyncID() string {
	if x != nil {
		return x.SyncID
	}
	return ""
}

func (x *Sync2Device) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *Sync2Device) GetFileID() string {
	if x != nil {
		return x.FileID
	}
	return ""
}

func (x *Sync2Device) GetFileLocation() string {
	if x != nil {
		return x.FileLocation
	}
	return ""
}

type AddKindle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KindleEmail string `protobuf:"bytes,1,opt,name=kindleEmail,proto3" json:"kindleEmail,omitempty"`
}

func (x *AddKindle) Reset() {
	*x = AddKindle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKindle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKindle) ProtoMessage() {}

func (x *AddKindle) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKindle.ProtoReflect.Descriptor instead.
func (*AddKindle) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{2}
}

func (x *AddKindle) GetKindleEmail() string {
	if x != nil {
		return x.KindleEmail
	}
	return ""
}

var File_device_proto protoreflect.FileDescriptor

var file_device_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44,
	0x12, 0x32, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x6f, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x4b, 0x69, 0x6e, 0x64, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x6f,
	0x2e, 0x41, 0x64, 0x64, 0x4b, 0x69, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x09, 0x61, 0x64, 0x64, 0x4b,
	0x69, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x7d, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x32, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4b, 0x69, 0x6e, 0x64, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6b, 0x69, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6b, 0x69, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x2a, 0x2c, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x44, 0x45, 0x56,
	0x49, 0x43, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4b, 0x49, 0x4e, 0x44, 0x4c, 0x45, 0x10,
	0x01, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_device_proto_rawDescOnce sync.Once
	file_device_proto_rawDescData = file_device_proto_rawDesc
)

func file_device_proto_rawDescGZIP() []byte {
	file_device_proto_rawDescOnce.Do(func() {
		file_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_proto_rawDescData)
	})
	return file_device_proto_rawDescData
}

var file_device_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_device_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_device_proto_goTypes = []interface{}{
	(DeviceType)(0),     // 0: scribo.DeviceType
	(*AddDevice)(nil),   // 1: scribo.AddDevice
	(*Sync2Device)(nil), // 2: scribo.Sync2Device
	(*AddKindle)(nil),   // 3: scribo.AddKindle
}
var file_device_proto_depIdxs = []int32{
	0, // 0: scribo.AddDevice.deviceType:type_name -> scribo.DeviceType
	3, // 1: scribo.AddDevice.addKindle:type_name -> scribo.AddKindle
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_device_proto_init() }
func file_device_proto_init() {
	if File_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDevice); i {
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
		file_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sync2Device); i {
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
		file_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKindle); i {
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
			RawDescriptor: file_device_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_device_proto_goTypes,
		DependencyIndexes: file_device_proto_depIdxs,
		EnumInfos:         file_device_proto_enumTypes,
		MessageInfos:      file_device_proto_msgTypes,
	}.Build()
	File_device_proto = out.File
	file_device_proto_rawDesc = nil
	file_device_proto_goTypes = nil
	file_device_proto_depIdxs = nil
}
