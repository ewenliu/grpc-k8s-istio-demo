// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: gcd.proto

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

type GCDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A uint64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B uint64 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *GCDRequest) Reset() {
	*x = GCDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gcd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCDRequest) ProtoMessage() {}

func (x *GCDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gcd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCDRequest.ProtoReflect.Descriptor instead.
func (*GCDRequest) Descriptor() ([]byte, []int) {
	return file_gcd_proto_rawDescGZIP(), []int{0}
}

func (x *GCDRequest) GetA() uint64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *GCDRequest) GetB() uint64 {
	if x != nil {
		return x.B
	}
	return 0
}

type GCDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result uint64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GCDResponse) Reset() {
	*x = GCDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gcd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCDResponse) ProtoMessage() {}

func (x *GCDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gcd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCDResponse.ProtoReflect.Descriptor instead.
func (*GCDResponse) Descriptor() ([]byte, []int) {
	return file_gcd_proto_rawDescGZIP(), []int{1}
}

func (x *GCDResponse) GetResult() uint64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_gcd_proto protoreflect.FileDescriptor

var file_gcd_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x63, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x28, 0x0a, 0x0a, 0x47, 0x43, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a,
	0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01, 0x62, 0x22, 0x25, 0x0a, 0x0b, 0x47, 0x43, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0x3a, 0x0a, 0x0a, 0x47, 0x43, 0x44, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x43, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x43, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02,
	0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gcd_proto_rawDescOnce sync.Once
	file_gcd_proto_rawDescData = file_gcd_proto_rawDesc
)

func file_gcd_proto_rawDescGZIP() []byte {
	file_gcd_proto_rawDescOnce.Do(func() {
		file_gcd_proto_rawDescData = protoimpl.X.CompressGZIP(file_gcd_proto_rawDescData)
	})
	return file_gcd_proto_rawDescData
}

var file_gcd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gcd_proto_goTypes = []interface{}{
	(*GCDRequest)(nil),  // 0: pb.GCDRequest
	(*GCDResponse)(nil), // 1: pb.GCDResponse
}
var file_gcd_proto_depIdxs = []int32{
	0, // 0: pb.GCDService.Compute:input_type -> pb.GCDRequest
	1, // 1: pb.GCDService.Compute:output_type -> pb.GCDResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gcd_proto_init() }
func file_gcd_proto_init() {
	if File_gcd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gcd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCDRequest); i {
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
		file_gcd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCDResponse); i {
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
			RawDescriptor: file_gcd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gcd_proto_goTypes,
		DependencyIndexes: file_gcd_proto_depIdxs,
		MessageInfos:      file_gcd_proto_msgTypes,
	}.Build()
	File_gcd_proto = out.File
	file_gcd_proto_rawDesc = nil
	file_gcd_proto_goTypes = nil
	file_gcd_proto_depIdxs = nil
}
