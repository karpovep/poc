// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: protos/cloud.proto

package cloud

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OperationStatus int32

const (
	OperationStatus_OK    OperationStatus = 0
	OperationStatus_ERROR OperationStatus = 1
)

// Enum value maps for OperationStatus.
var (
	OperationStatus_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	OperationStatus_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x OperationStatus) Enum() *OperationStatus {
	p := new(OperationStatus)
	*p = x
	return p
}

func (x OperationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_cloud_proto_enumTypes[0].Descriptor()
}

func (OperationStatus) Type() protoreflect.EnumType {
	return &file_protos_cloud_proto_enumTypes[0]
}

func (x OperationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationStatus.Descriptor instead.
func (OperationStatus) EnumDescriptor() ([]byte, []int) {
	return file_protos_cloud_proto_rawDescGZIP(), []int{0}
}

type CloudObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *anypb.Any `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *CloudObject) Reset() {
	*x = CloudObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cloud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudObject) ProtoMessage() {}

func (x *CloudObject) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cloud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudObject.ProtoReflect.Descriptor instead.
func (*CloudObject) Descriptor() ([]byte, []int) {
	return file_protos_cloud_proto_rawDescGZIP(), []int{0}
}

func (x *CloudObject) GetEntity() *anypb.Any {
	if x != nil {
		return x.Entity
	}
	return nil
}

type OperationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status OperationStatus `protobuf:"varint,1,opt,name=status,proto3,enum=protos.OperationStatus" json:"status,omitempty"`
}

func (x *OperationResult) Reset() {
	*x = OperationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cloud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationResult) ProtoMessage() {}

func (x *OperationResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cloud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationResult.ProtoReflect.Descriptor instead.
func (*OperationResult) Descriptor() ([]byte, []int) {
	return file_protos_cloud_proto_rawDescGZIP(), []int{1}
}

func (x *OperationResult) GetStatus() OperationStatus {
	if x != nil {
		return x.Status
	}
	return OperationStatus_OK
}

type TestEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TestEntity) Reset() {
	*x = TestEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cloud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestEntity) ProtoMessage() {}

func (x *TestEntity) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cloud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestEntity.ProtoReflect.Descriptor instead.
func (*TestEntity) Descriptor() ([]byte, []int) {
	return file_protos_cloud_proto_rawDescGZIP(), []int{2}
}

func (x *TestEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cloud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cloud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_protos_cloud_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_protos_cloud_proto protoreflect.FileDescriptor

var file_protos_cloud_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0b, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x42, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x20, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x2a, 0x24, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x32, 0x7e, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_cloud_proto_rawDescOnce sync.Once
	file_protos_cloud_proto_rawDescData = file_protos_cloud_proto_rawDesc
)

func file_protos_cloud_proto_rawDescGZIP() []byte {
	file_protos_cloud_proto_rawDescOnce.Do(func() {
		file_protos_cloud_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_cloud_proto_rawDescData)
	})
	return file_protos_cloud_proto_rawDescData
}

var file_protos_cloud_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_cloud_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_cloud_proto_goTypes = []interface{}{
	(OperationStatus)(0),     // 0: protos.OperationStatus
	(*CloudObject)(nil),      // 1: protos.CloudObject
	(*OperationResult)(nil),  // 2: protos.OperationResult
	(*TestEntity)(nil),       // 3: protos.TestEntity
	(*SubscribeRequest)(nil), // 4: protos.SubscribeRequest
	(*anypb.Any)(nil),        // 5: google.protobuf.Any
}
var file_protos_cloud_proto_depIdxs = []int32{
	5, // 0: protos.CloudObject.entity:type_name -> google.protobuf.Any
	0, // 1: protos.OperationResult.status:type_name -> protos.OperationStatus
	1, // 2: protos.Cloud.Commit:input_type -> protos.CloudObject
	1, // 3: protos.Cloud.Subscribe:input_type -> protos.CloudObject
	2, // 4: protos.Cloud.Commit:output_type -> protos.OperationResult
	1, // 5: protos.Cloud.Subscribe:output_type -> protos.CloudObject
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_cloud_proto_init() }
func file_protos_cloud_proto_init() {
	if File_protos_cloud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_cloud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudObject); i {
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
		file_protos_cloud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationResult); i {
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
		file_protos_cloud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestEntity); i {
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
		file_protos_cloud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
			RawDescriptor: file_protos_cloud_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_cloud_proto_goTypes,
		DependencyIndexes: file_protos_cloud_proto_depIdxs,
		EnumInfos:         file_protos_cloud_proto_enumTypes,
		MessageInfos:      file_protos_cloud_proto_msgTypes,
	}.Build()
	File_protos_cloud_proto = out.File
	file_protos_cloud_proto_rawDesc = nil
	file_protos_cloud_proto_goTypes = nil
	file_protos_cloud_proto_depIdxs = nil
}
