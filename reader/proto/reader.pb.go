// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: reader.proto

package proto

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

type VolumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerAddr string `protobuf:"bytes,1,opt,name=serverAddr,proto3" json:"serverAddr,omitempty"`
	VolumeName string `protobuf:"bytes,2,opt,name=volumeName,proto3" json:"volumeName,omitempty"`
	BackupName string `protobuf:"bytes,3,opt,name=BackupName,proto3" json:"BackupName,omitempty"`
}

func (x *VolumeRequest) Reset() {
	*x = VolumeRequest{}
	mi := &file_reader_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VolumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeRequest) ProtoMessage() {}

func (x *VolumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reader_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeRequest.ProtoReflect.Descriptor instead.
func (*VolumeRequest) Descriptor() ([]byte, []int) {
	return file_reader_proto_rawDescGZIP(), []int{0}
}

func (x *VolumeRequest) GetServerAddr() string {
	if x != nil {
		return x.ServerAddr
	}
	return ""
}

func (x *VolumeRequest) GetVolumeName() string {
	if x != nil {
		return x.VolumeName
	}
	return ""
}

func (x *VolumeRequest) GetBackupName() string {
	if x != nil {
		return x.BackupName
	}
	return ""
}

type VolumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *VolumeResponse) Reset() {
	*x = VolumeResponse{}
	mi := &file_reader_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VolumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeResponse) ProtoMessage() {}

func (x *VolumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reader_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeResponse.ProtoReflect.Descriptor instead.
func (*VolumeResponse) Descriptor() ([]byte, []int) {
	return file_reader_proto_rawDescGZIP(), []int{1}
}

func (x *VolumeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_reader_proto protoreflect.FileDescriptor

var file_reader_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x6f, 0x0a, 0x0d, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x4f, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x62, 0x68, 0x69, 0x73, 0x68, 0x65, 0x6b, 0x2d, 0x6b, 0x61, 0x6d, 0x61,
	0x74, 0x2d, 0x6e, 0x75, 0x74, 0x61, 0x6e, 0x69, 0x78, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reader_proto_rawDescOnce sync.Once
	file_reader_proto_rawDescData = file_reader_proto_rawDesc
)

func file_reader_proto_rawDescGZIP() []byte {
	file_reader_proto_rawDescOnce.Do(func() {
		file_reader_proto_rawDescData = protoimpl.X.CompressGZIP(file_reader_proto_rawDescData)
	})
	return file_reader_proto_rawDescData
}

var file_reader_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_reader_proto_goTypes = []any{
	(*VolumeRequest)(nil),  // 0: reader.VolumeRequest
	(*VolumeResponse)(nil), // 1: reader.VolumeResponse
}
var file_reader_proto_depIdxs = []int32{
	0, // 0: reader.ReaderService.MigrateVolume:input_type -> reader.VolumeRequest
	1, // 1: reader.ReaderService.MigrateVolume:output_type -> reader.VolumeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reader_proto_init() }
func file_reader_proto_init() {
	if File_reader_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reader_proto_goTypes,
		DependencyIndexes: file_reader_proto_depIdxs,
		MessageInfos:      file_reader_proto_msgTypes,
	}.Build()
	File_reader_proto = out.File
	file_reader_proto_rawDesc = nil
	file_reader_proto_goTypes = nil
	file_reader_proto_depIdxs = nil
}