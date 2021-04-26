// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: protos/nodes/nodes.proto

package nodes

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

type InternalServerObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Entity *anypb.Any `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *InternalServerObject) Reset() {
	*x = InternalServerObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_nodes_nodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalServerObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalServerObject) ProtoMessage() {}

func (x *InternalServerObject) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nodes_nodes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalServerObject.ProtoReflect.Descriptor instead.
func (*InternalServerObject) Descriptor() ([]byte, []int) {
	return file_protos_nodes_nodes_proto_rawDescGZIP(), []int{0}
}

func (x *InternalServerObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InternalServerObject) GetEntity() *anypb.Any {
	if x != nil {
		return x.Entity
	}
	return nil
}

type Acknowledge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Acknowledge) Reset() {
	*x = Acknowledge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_nodes_nodes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acknowledge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acknowledge) ProtoMessage() {}

func (x *Acknowledge) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nodes_nodes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acknowledge.ProtoReflect.Descriptor instead.
func (*Acknowledge) Descriptor() ([]byte, []int) {
	return file_protos_nodes_nodes_proto_rawDescGZIP(), []int{1}
}

var File_protos_nodes_nodes_proto protoreflect.FileDescriptor

var file_protos_nodes_nodes_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x64, 0x70, 0x63, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x14, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2c, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x0d, 0x0a,
	0x0b, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x32, 0x5b, 0x0a, 0x04,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x53, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x70, 0x63, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x64, 0x70, 0x63, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x6b, 0x6e,
	0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x64, 0x70, 0x63, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_nodes_nodes_proto_rawDescOnce sync.Once
	file_protos_nodes_nodes_proto_rawDescData = file_protos_nodes_nodes_proto_rawDesc
)

func file_protos_nodes_nodes_proto_rawDescGZIP() []byte {
	file_protos_nodes_nodes_proto_rawDescOnce.Do(func() {
		file_protos_nodes_nodes_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_nodes_nodes_proto_rawDescData)
	})
	return file_protos_nodes_nodes_proto_rawDescData
}

var file_protos_nodes_nodes_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_nodes_nodes_proto_goTypes = []interface{}{
	(*InternalServerObject)(nil), // 0: protos.dpc.nodes.InternalServerObject
	(*Acknowledge)(nil),          // 1: protos.dpc.nodes.Acknowledge
	(*anypb.Any)(nil),            // 2: google.protobuf.Any
}
var file_protos_nodes_nodes_proto_depIdxs = []int32{
	2, // 0: protos.dpc.nodes.InternalServerObject.entity:type_name -> google.protobuf.Any
	0, // 1: protos.dpc.nodes.Node.Transfer:input_type -> protos.dpc.nodes.InternalServerObject
	1, // 2: protos.dpc.nodes.Node.Transfer:output_type -> protos.dpc.nodes.Acknowledge
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_nodes_nodes_proto_init() }
func file_protos_nodes_nodes_proto_init() {
	if File_protos_nodes_nodes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_nodes_nodes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalServerObject); i {
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
		file_protos_nodes_nodes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Acknowledge); i {
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
			RawDescriptor: file_protos_nodes_nodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_nodes_nodes_proto_goTypes,
		DependencyIndexes: file_protos_nodes_nodes_proto_depIdxs,
		MessageInfos:      file_protos_nodes_nodes_proto_msgTypes,
	}.Build()
	File_protos_nodes_nodes_proto = out.File
	file_protos_nodes_nodes_proto_rawDesc = nil
	file_protos_nodes_nodes_proto_goTypes = nil
	file_protos_nodes_nodes_proto_depIdxs = nil
}
