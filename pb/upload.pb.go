// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: upload.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UploadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FilePath      string                 `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Chunks        []byte                 `protobuf:"bytes,2,opt,name=chunks,proto3" json:"chunks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	mi := &file_upload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *UploadRequest) GetChunks() []byte {
	if x != nil {
		return x.Chunks
	}
	return nil
}

type UploadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileSize      int64                  `protobuf:"varint,1,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	mi := &file_upload_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResponse) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *UploadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_upload_proto protoreflect.FileDescriptor

const file_upload_proto_rawDesc = "" +
	"\n" +
	"\fupload.proto\"D\n" +
	"\rUploadRequest\x12\x1b\n" +
	"\tfile_path\x18\x01 \x01(\tR\bfilePath\x12\x16\n" +
	"\x06chunks\x18\x02 \x01(\fR\x06chunks\"G\n" +
	"\x0eUploadResponse\x12\x1b\n" +
	"\tfile_size\x18\x01 \x01(\x03R\bfileSize\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2;\n" +
	"\fStreamUpload\x12+\n" +
	"\x06Upload\x12\x0e.UploadRequest\x1a\x0f.UploadResponse(\x01B\x06Z\x04.;pbb\x06proto3"

var (
	file_upload_proto_rawDescOnce sync.Once
	file_upload_proto_rawDescData []byte
)

func file_upload_proto_rawDescGZIP() []byte {
	file_upload_proto_rawDescOnce.Do(func() {
		file_upload_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_upload_proto_rawDesc), len(file_upload_proto_rawDesc)))
	})
	return file_upload_proto_rawDescData
}

var file_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_upload_proto_goTypes = []any{
	(*UploadRequest)(nil),  // 0: UploadRequest
	(*UploadResponse)(nil), // 1: UploadResponse
}
var file_upload_proto_depIdxs = []int32{
	0, // 0: StreamUpload.Upload:input_type -> UploadRequest
	1, // 1: StreamUpload.Upload:output_type -> UploadResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_upload_proto_init() }
func file_upload_proto_init() {
	if File_upload_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_upload_proto_rawDesc), len(file_upload_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upload_proto_goTypes,
		DependencyIndexes: file_upload_proto_depIdxs,
		MessageInfos:      file_upload_proto_msgTypes,
	}.Build()
	File_upload_proto = out.File
	file_upload_proto_goTypes = nil
	file_upload_proto_depIdxs = nil
}
