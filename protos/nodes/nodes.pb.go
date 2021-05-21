// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: protos/nodes/nodes.proto

package nodes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	cloud "poc/protos/cloud"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IsoMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitialNodeId string `protobuf:"bytes,1,opt,name=initialNodeId,proto3" json:"initialNodeId,omitempty"`
	RetryIn       int32  `protobuf:"varint,2,opt,name=retryIn,proto3" json:"retryIn,omitempty"`
}

func (x *IsoMeta) Reset() {
	*x = IsoMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_nodes_nodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsoMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsoMeta) ProtoMessage() {}

func (x *IsoMeta) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IsoMeta.ProtoReflect.Descriptor instead.
func (*IsoMeta) Descriptor() ([]byte, []int) {
	return file_protos_nodes_nodes_proto_rawDescGZIP(), []int{0}
}

func (x *IsoMeta) GetInitialNodeId() string {
	if x != nil {
		return x.InitialNodeId
	}
	return ""
}

func (x *IsoMeta) GetRetryIn() int32 {
	if x != nil {
		return x.RetryIn
	}
	return 0
}

// ISO stands for Internal Server Object
type ISO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudObj *cloud.CloudObject `protobuf:"bytes,1,opt,name=cloudObj,proto3" json:"cloudObj,omitempty"`
	Metadata *IsoMeta           `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	//nodeId to timestamp map when the object was transferred from the node
	TransferredByNodes map[string]int64 `protobuf:"bytes,3,rep,name=transferredByNodes,proto3" json:"transferredByNodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	SenderNodeId       string           `protobuf:"bytes,4,opt,name=senderNodeId,proto3" json:"senderNodeId,omitempty"`
}

func (x *ISO) Reset() {
	*x = ISO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_nodes_nodes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ISO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ISO) ProtoMessage() {}

func (x *ISO) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ISO.ProtoReflect.Descriptor instead.
func (*ISO) Descriptor() ([]byte, []int) {
	return file_protos_nodes_nodes_proto_rawDescGZIP(), []int{1}
}

func (x *ISO) GetCloudObj() *cloud.CloudObject {
	if x != nil {
		return x.CloudObj
	}
	return nil
}

func (x *ISO) GetMetadata() *IsoMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ISO) GetTransferredByNodes() map[string]int64 {
	if x != nil {
		return x.TransferredByNodes
	}
	return nil
}

func (x *ISO) GetSenderNodeId() string {
	if x != nil {
		return x.SenderNodeId
	}
	return ""
}

type Acknowledge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Acknowledge) Reset() {
	*x = Acknowledge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_nodes_nodes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acknowledge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acknowledge) ProtoMessage() {}

func (x *Acknowledge) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nodes_nodes_proto_msgTypes[2]
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
	return file_protos_nodes_nodes_proto_rawDescGZIP(), []int{2}
}

type NodeInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NodeInfoRequest) Reset() {
	*x = NodeInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_nodes_nodes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfoRequest) ProtoMessage() {}

func (x *NodeInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nodes_nodes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfoRequest.ProtoReflect.Descriptor instead.
func (*NodeInfoRequest) Descriptor() ([]byte, []int) {
	return file_protos_nodes_nodes_proto_rawDescGZIP(), []int{3}
}

type NodeInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NodeInfoResponse) Reset() {
	*x = NodeInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_nodes_nodes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfoResponse) ProtoMessage() {}

func (x *NodeInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nodes_nodes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfoResponse.ProtoReflect.Descriptor instead.
func (*NodeInfoResponse) Descriptor() ([]byte, []int) {
	return file_protos_nodes_nodes_proto_rawDescGZIP(), []int{4}
}

func (x *NodeInfoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_protos_nodes_nodes_proto protoreflect.FileDescriptor

var file_protos_nodes_nodes_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x6f, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x1a, 0x18, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x07, 0x49, 0x73, 0x6f, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x49, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x79, 0x49,
	0x6e, 0x22, 0xc1, 0x02, 0x0a, 0x03, 0x49, 0x53, 0x4f, 0x12, 0x39, 0x0a, 0x08, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x4f, 0x62, 0x6a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x6f,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x4f, 0x62, 0x6a, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x49, 0x73, 0x6f, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5d, 0x0a, 0x12, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x42, 0x79, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x6f, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x49, 0x53, 0x4f, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x42, 0x79, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x64, 0x42, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x1a, 0x45,
	0x0a, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x42, 0x79, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x9e, 0x01, 0x0a, 0x04,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x52, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x21, 0x2e, 0x70, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x49, 0x53, 0x4f, 0x1a, 0x1d, 0x2e, 0x70, 0x6f,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x41,
	0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10,
	0x70, 0x6f, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_protos_nodes_nodes_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_nodes_nodes_proto_goTypes = []interface{}{
	(*IsoMeta)(nil),           // 0: poc.protos.nodes.IsoMeta
	(*ISO)(nil),               // 1: poc.protos.nodes.ISO
	(*Acknowledge)(nil),       // 2: poc.protos.nodes.Acknowledge
	(*NodeInfoRequest)(nil),   // 3: poc.protos.nodes.NodeInfoRequest
	(*NodeInfoResponse)(nil),  // 4: poc.protos.nodes.NodeInfoResponse
	nil,                       // 5: poc.protos.nodes.ISO.TransferredByNodesEntry
	(*cloud.CloudObject)(nil), // 6: poc.protos.cloud.CloudObject
}
var file_protos_nodes_nodes_proto_depIdxs = []int32{
	6, // 0: poc.protos.nodes.ISO.cloudObj:type_name -> poc.protos.cloud.CloudObject
	0, // 1: poc.protos.nodes.ISO.metadata:type_name -> poc.protos.nodes.IsoMeta
	5, // 2: poc.protos.nodes.ISO.transferredByNodes:type_name -> poc.protos.nodes.ISO.TransferredByNodesEntry
	3, // 3: poc.protos.nodes.Node.GetInfo:input_type -> poc.protos.nodes.NodeInfoRequest
	1, // 4: poc.protos.nodes.Node.Transfer:input_type -> poc.protos.nodes.ISO
	4, // 5: poc.protos.nodes.Node.GetInfo:output_type -> poc.protos.nodes.NodeInfoResponse
	2, // 6: poc.protos.nodes.Node.Transfer:output_type -> poc.protos.nodes.Acknowledge
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_nodes_nodes_proto_init() }
func file_protos_nodes_nodes_proto_init() {
	if File_protos_nodes_nodes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_nodes_nodes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsoMeta); i {
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
			switch v := v.(*ISO); i {
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
		file_protos_nodes_nodes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_nodes_nodes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfoRequest); i {
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
		file_protos_nodes_nodes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfoResponse); i {
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
			NumMessages:   6,
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
